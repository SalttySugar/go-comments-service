package service

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Connect(conf *DatabaseConfiguration) *gorm.DB {
	if connection != nil {return connection
	}

	var connectors = map[DatabaseDriverType]func(*DatabaseConfiguration) *gorm.DB{
		MARIADB:  MySqlConnection,
		MYSQL:    MySqlConnection,
		POSTGRES: PostgresqlConnection,
		SQLITE:   SqliteConnection,
	}

	if val, ok := connectors[conf.Driver]; ok {
		connection = val(conf)
	} else {
		panic("could not find connector for: " + conf.Driver)
	}

	connection.AutoMigrate(&Comment{})
	return  connection
}

func MySqlConnection(conf *DatabaseConfiguration) *gorm.DB {
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	conn, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		panic("failed to connect to mysql database")
	}

	return conn
}

func PostgresqlConnection(conf *DatabaseConfiguration) *gorm.DB {
	source := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Name,
		conf.Port,
	)

	conn, err := gorm.Open(postgres.Open(source), &gorm.Config{})

	if err != nil {
		panic("failed to connect to mysql database")
	}

	return conn
}

func SqliteConnection(conf *DatabaseConfiguration) *gorm.DB {
	conn, err := gorm.Open(sqlite.Open(conf.Path))
	if err != nil {
		panic("failed to connect to mysql database")
	}
	return conn
}
