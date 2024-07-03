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
	// Leer variables de entorno
	DBName := os.Getenv("DB_NAME") // Nombre de la base de datos
	DBUser := os.Getenv("DB_USER") // Usuario de la base de datos
	DBPass := os.Getenv("DB_PASS") // Contraseña del usuario de la base de datos
	DBHost := os.Getenv("DB_HOST") // Debería ser "localhost" cuando se usa Cloud SQL Proxy
	DBPort := os.Getenv("DB_PORT") // Debería ser el puerto que el proxy expone

	log.Info("Started connecting database...")

	// Construir la cadena de conexión
	dsn := DBUser + ":" + DBPass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8&parseTime=True"

	// Abrir conexión a la base de datos
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// Asignar la conexión de base de datos a los clientes
	users.Db = db
	reminders.Db = db
}

func StartDbEngine() {
	// Migrar las tablas de la base de datos
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Reminder{})

	log.Info("Finishing Migration Database Tables")
}
