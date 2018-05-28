package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"log"
	"strings"
)

func InitTables(tables string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connectString := fmt.Sprintf("%s:%s@(%s)/%s", dbUser, dbPassword, dbHost, dbName)
	db, err := sql.Open("mysql", connectString)
	handleError("OPEN", err)
	defer db.Close()

	err = db.Ping()
	handleError("CONNECT", err)

	if tables == "" {
		createServerTable(db)
		createJobTable(db)
		createLogTable(db)
	} else {
		tablesSlice := strings.Split(tables, ",")
		for _, table := range tablesSlice {
			switch table {
			case "servers", "server":
				createServerTable(db)
			case "jobs", "job":
				createJobTable(db)
			case "logs", "log":
				createLogTable(db)
			default:
				log.Fatalf("DB: table %s is not needed.", table)
			}
		}
	}
}

func createServerTable(db *sql.DB) {
	createSql := `CREATE TABLE IF NOT EXISTS servers (
	id int(11) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(500) DEFAULT NULL,
	user varchar(100) DEFAULT NULL,
	host varchar(100) DEFAULT NULL,
	auth_type tinyint(11) DEFAULT NULL COMMENT 'auth type，0:password，1:private key，2: proxy',
	pass varchar(100) DEFAULT NULL,
	private_key_path varchar(1000) DEFAULT NULL,
	proxy_server_id int(11) DEFAULT NULL,
	PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	createTable(db, "servers", createSql)
}

func createJobTable(db *sql.DB) {
	createSql := `CREATE TABLE IF NOT EXISTS jobs (
  	id int(11) unsigned NOT NULL AUTO_INCREMENT,
  	name varchar(500) DEFAULT NULL,
  	server_id int(11) DEFAULT NULL,
  	command text,
  	status tinyint(11) DEFAULT NULL COMMENT 'job status, 0: disabled，1: enabled',
  	PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	createTable(db, "jobs", createSql)
}

func createLogTable(db *sql.DB) {
	createSql := `CREATE TABLE IF NOT EXISTS logs (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    job_id int(11) DEFAULT NULL,
    start_time datetime DEFAULT NULL,
    end_time datetime DEFAULT NULL,
    status tinyint(11) DEFAULT NULL COMMENT 'job run status，0: running，1: success，2: failed，3: waiting',
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	createTable(db, "logs", createSql)
}

func createTable(db *sql.DB, table string, createSql string) {
	deleteSql := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
	_, err := db.Exec(deleteSql)
	handleError("DROP", err)
	log.Printf("DB: table %s is deleted\n", table)
	_, err = db.Exec(createSql)
	handleError("CREATE", err)
	log.Printf("DB: table %s is created\n", table)
}

func handleError(msg string, err error) {
	if err != nil {
		msg := fmt.Sprintf("DB %s ERROR: %v", msg, err)
		fmt.Println(msg)
		log.Fatalln(msg)
		os.Exit(1)
	}
}
