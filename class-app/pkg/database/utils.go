package database

import "github.com/go-pg/pg/v10"

// Check if a table exists in the database
func TableExists(db *pg.DB, tableName string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (
        SELECT 1 
        FROM information_schema.tables 
        WHERE table_name = ?
    )`
	_, err := db.QueryOne(pg.Scan(&exists), query, tableName)
	if err != nil {
		return false, err
	}
	return exists, nil
}
