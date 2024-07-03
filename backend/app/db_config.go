package app

import (
	"backend/clients/reminders"
	"backend/clients/users"
	"backend/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := os.Getenv("DB_NAME") //Nombre de la base de datos local de ustedes
	DBUser := os.Getenv("DB_USER") //usuario de la base de datos, habitualmente root
	DBPass := os.Getenv("DB_PASS") //password del root en la instalacion
	DBHost := os.Getenv("DB_HOST") //host de la base de datos. hbitualmente 127.0.0.1
	// ------------------------
	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	users.Db = db
	reminders.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Reminder{})

	log.Info("Finishing Migration Database Tables")
}
