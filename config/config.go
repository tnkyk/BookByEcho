package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	DriveName string
	Database  string
	Password  string
	User      string
	Host      string
	LogFile   string
}

var DB *sql.DB
var Config ConfigList

func init() {
	cfg, err := ini.Load("configs.ini")
	if err != nil {
		log.Printf("can not load config.ini: %v\n", err)
		os.Exit(1)
	}
	Config = ConfigList{
		Port:      cfg.Section("Web").Key("port").String(),
		DriveName: cfg.Section("DB").Key("DriveName").String(),
		Database:  cfg.Section("DB").Key("Database").String(),
		Password:  cfg.Section("DB").Key("Password").String(),
		User:      cfg.Section("DB").Key("User").String(),
		Host:      cfg.Section("DB").Key("Host").String(),
		LogFile:   cfg.Section("Log").Key("log_file").String(),
	}
	/* ===== データベースへ接続する. ===== */
	// ユーザ
	//user := os.Getenv("MYSQL_USER")
	user := Config.User
	// パスワード
	password := Config.Password
	// 接続先ホスト
	//host := os.Getenv("MYSQL_HOST")
	host := Config.Host
	// 接続先ポート
	//port := os.Getenv("MYSQL_PORT")
	port := Config.Port
	// 接続先データベース
	//database := os.Getenv("MYSQL_DATABASE")
	database := Config.Database
	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database

	DB, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
}
