package gorm

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lelinu/api_utils/log/llogrus"
	migrate "github.com/rubenv/sql-migrate"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qor/validations"
)

const DatabaseLogProcess = "database"
const (
	DefaultMigrationsFolder       = "migrations/mysql"
	DefaultMaxIdleConnections     = 10
	DefaultMaxOpenConnections     = 50
	DefaultConnMaxLifetimeSeconds = 60
	DefaultRetries                = 10
)

type IDbManager interface{
	SetMigrationsFolder(migrationsFolder string)
	SetConfigurations(maxIdleConnections, maxOpenConnections, connMaxLifetimeSeconds, dbRetries int)
	Init(logMode, autoMigrate bool) error
	GetSession() *gorm.DB
}

//DbManager is a container for all data and methods related to the database connection.
type DbManager struct {
	Logger                 llogrus.IService
	DatabaseURL            string
	DatabaseName           string
	MigrationsFolder       string
	MaxIdleConnections     int
	MaxOpenConnections     int
	ConnMaxLifetimeSeconds int
	DBRetries              int

	Session *gorm.DB
}

// NewDbManager constructor to setup the basic requirements of a DbManager
func NewDbManager(logger llogrus.IService, databaseURL, databaseName string) IDbManager {
	return &DbManager{
		Logger:                 logger,
		DatabaseURL:            databaseURL,
		DatabaseName:           databaseName,
	}
}

// SetMigrationsFolder sets the migrations folder, must be executed before Init
func (dbm *DbManager) SetMigrationsFolder(migrationsFolder string) {
	dbm.MigrationsFolder = migrationsFolder
}

// SetConfigurations sets the configurations of the connection, must be executed before Init
func (dbm *DbManager) SetConfigurations(maxIdleConnections, maxOpenConnections, connMaxLifetimeSeconds, dbRetries int) {
	dbm.MaxIdleConnections = maxIdleConnections
	dbm.MaxOpenConnections = maxOpenConnections
	dbm.ConnMaxLifetimeSeconds = connMaxLifetimeSeconds
	dbm.DBRetries = dbRetries
}

//Init sets up the data needed for this object.
func (dbm *DbManager) Init(logMode, autoMigrate bool) error {
	if err := dbm.createDatabase(); err != nil {
		return err
	}

	if err := dbm.connect(); err != nil {
		return err
	}

	dbm.Logger.Info("setting log mode", fmt.Sprintf("process:%v,logMode:%v", DatabaseLogProcess, logMode))
	dbm.Session.LogMode(logMode)

	dbm.setConfigurations()
	validations.RegisterCallbacks(dbm.Session)

	if autoMigrate {
		if err := dbm.migrate(); err != nil {
			return err
		}
	}

	return nil
}

//GetSession gets the current session
func (dbm *DbManager) GetSession() *gorm.DB {
	return dbm.Session
}

// createDatabase creates the database if it doesn't exist
func (dbm *DbManager) createDatabase() error {

	fmt.Printf("database name %v", dbm.DatabaseName)
	fmt.Printf("database url %v", dbm.DatabaseURL)

	dbm.Logger.Info("creating database if it doesn't exist", fmt.Sprintf("process:%v,database Url:%v, database Name:%v", DatabaseLogProcess, dbm.DatabaseURL, dbm.DatabaseName))

	db, err := sql.Open("mysql", fmt.Sprintf("%s/", dbm.DatabaseURL))
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbm.DatabaseName)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';", dbm.DatabaseName))
		if err != nil {
			return err
		}
	}

	return nil
}

//connect establishes a connection to the database, using a naive retry method.DBManager
func (dbm *DbManager) connect() error {

	dbm.Logger.Info("establishing connection", fmt.Sprintf("process:%v,database Url:%v, database Name:%v", DatabaseLogProcess, dbm.DatabaseURL, dbm.DatabaseName))

	if dbm.DBRetries == 0 {
		dbm.DBRetries = DefaultRetries
	}

	var err error
	var db *gorm.DB

	for t := 1; t < dbm.DBRetries; t++ {
		if db, err = gorm.Open("mysql", fmt.Sprintf("%s/%s?charset=utf8&parseTime=true", dbm.DatabaseURL, dbm.DatabaseName)); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return err
	}

	dbm.Session = db
	return nil
}

// setConfigurations sets the configuration of the manager either by default values or DBManager values
func (dbm *DbManager) setConfigurations() {
	// Set the default values for the connection configuration variables
	if dbm.MaxIdleConnections == 0 {
		dbm.MaxIdleConnections = DefaultMaxIdleConnections
	}

	if dbm.MaxOpenConnections == 0 {
		dbm.MaxOpenConnections = DefaultMaxOpenConnections
	}

	if dbm.ConnMaxLifetimeSeconds == 0 {
		dbm.ConnMaxLifetimeSeconds = DefaultConnMaxLifetimeSeconds
	}

	dbm.Logger.Info("setting connection configurations", fmt.Sprintf("process:%v,max idle connections:%v,max open connections:%v, connection max lifetime seconds:%v", DatabaseLogProcess,
		dbm.MaxIdleConnections, dbm.MaxOpenConnections,  dbm.ConnMaxLifetimeSeconds))

	// Set connection configuration variables
	dbm.Session.DB().SetMaxIdleConns(dbm.MaxIdleConnections)
	dbm.Session.DB().SetMaxOpenConns(dbm.MaxOpenConnections)
	dbm.Session.DB().SetConnMaxLifetime(time.Second * time.Duration(dbm.ConnMaxLifetimeSeconds))
}

// migrate will apply migrations to the database if needed.
func (dbm *DbManager) migrate() error {
	// Set the migrations folder
	migrationsFolder := DefaultMigrationsFolder
	if dbm.MigrationsFolder != "" {
		migrationsFolder = dbm.MigrationsFolder
	}

	dbm.Logger.Info("executing migrations", fmt.Sprintf("process:%v,migration folder:%v", DatabaseLogProcess, dbm.MigrationsFolder))

	// Fetch the migrations folder
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsFolder,
	}
	migrate.SetTable("migrations")

	// Migrate the database if needed
	n, err := migrate.Exec(dbm.Session.DB(), "mysql", migrations, migrate.Up)
	if err != nil {
		return nil
	}

	dbm.Logger.Info("Migrating the database", fmt.Sprintf("process: %v, message:applied migrations,number:%v", DatabaseLogProcess, n))

	return nil
}
