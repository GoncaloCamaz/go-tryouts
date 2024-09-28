package main

func (app *Config) sendEmail(msg Message) {
	// add to the wait group
	app.WaitGroup.Add(1)
	// send the message
	app.Mailer.MailerChan <- msg
}
