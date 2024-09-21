package main

import (
	"database/sql"
	"encoding/gob"
	"final-project/data"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	// we added _ to the import because we are not using the package directly
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func main() {
	// connect to the database
	db := initDB()

	// create sessions
	session := initSession()

	// create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create channels

	// create waitgroup
	wg := sync.WaitGroup{}

	// setup the application config
	app := Config{
		Session:   session,
		DB:        db,
		InfoLog:   infoLog,
		ErrorLog:  errorLog,
		WaitGroup: &wg,
		Models:    data.New(db),
	}
	// set up mail
	app.Mailer = app.createMail()
	go app.listenForMail()

	// listen for signals
	go app.listenForShutdown()

	// listen for web connections
	app.serve()
}

func (app *Config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	app.InfoLog.Println("Starting web server...")
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func initDB() *sql.DB {
	conn := connectToDB()

	if conn == nil {
		log.Panic("Can't connect to the database")
	}

	return conn
}

// lets try to access some database
func connectToDB() *sql.DB {
	counts := 0

	// connection string
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Println("Connected to the database")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Println("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++

		continue
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initSession() *scs.SessionManager {
	// this allows us to store the user in the session
	gob.Register(data.User{})

	// set up session
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}

	return redisPool
}

func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // block until we receive the signal

	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// perform any cleanup tasks
	app.InfoLog.Println("Would run cleanup tasks...")

	// block until wg is empty
	app.WaitGroup.Wait()

	app.Mailer.DoneChan <- true
	app.InfoLog.Println("Closing channels and shutting down application")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
}

func (app *Config) createMail() Mail {
	mail := Mail{
		Domain:      "localhsot",
		Host:        "localhsot",
		Port:        1025,
		Encryption:  "none",
		FromName:    "info",
		FromAddress: "info@mycompany.com",
		Wait:        app.WaitGroup,
		MailerChan:  make(chan Message, 100), //allow 100 messages to be queued
		ErrorChan:   make(chan error),
		DoneChan:    make(chan bool),
	}

	return mail
}
