package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open("mysql", DbURL(BuildDBConfig()))
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	// db.LogMode(true)
	DB = db
	return DB
}
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}
// Using this function to get a connection, you can create your connection pool here.
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "raul",
		Password: "hELLO@fRIEND2",
		DBName:   "first_go",
	}
	// fmt.Println(&dbConfig)
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func GetDB() *gorm.DB {
	return DB
}