package data

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/Arjun-P-Jayakrishnan/LCVS/data/sqlc"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

type FileDB struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

const (
	dbPath       = "data/lcvs.db"
	migrationDir = "migrations"
)

func NewFilesDB() *FileDB {

	filesDB:=&FileDB{}

	//Establishes a link to local database
	if connErr := connectDB(filesDB); connErr != nil {
		log.Fatal("Error while establishing a connection : ",connErr)
	}

	//Performs Migrations if not done
	if migrateErr := runMigrations(filesDB, migrationDir); migrateErr != nil {
		log.Fatal("Error while migrations :",migrateErr)
	}

	return filesDB
}

// ConnectDB establishes a link to local database
//
//	`fileDB` refernce to db state
func connectDB(fileDB *FileDB) error {

	//Ensure directory exists
	if osErr := os.MkdirAll(filepath.Dir(dbPath), 0755); osErr != nil {
		log.Fatal("\n Error finding file :", osErr)
		return osErr
	}

	//Open database
	conn, openErr := sql.Open("sqlite3", dbPath)

	if openErr != nil {
		log.Fatal("\n Error connecting to database :", openErr)
		return openErr
	}

	//Check if the database is connected
	if pingErr := conn.Ping(); pingErr != nil {
		log.Fatal("\n Failed to Ping the database :", pingErr)
		return pingErr

	}

	queries := sqlc.New(conn)

	fileDB.DB = conn
	fileDB.Queries = queries

	return nil
}

// RunMigrations for inital migrations
func runMigrations(fileDB *FileDB, migrationDir string) error {

	//Select which db type for migration
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal("\n Failed to set goose dialect :", err)
		return err
	}

	//Track migrations using goose up
	if err := goose.Up(fileDB.DB, migrationDir); err != nil {
		log.Fatal("\n Failed to run Up migrations :", err)
		return err
	}

	return nil
}

// Close Connection closes active connection
func CloseConnection(fileDB *FileDB) {
	defer fileDB.DB.Close()
}
