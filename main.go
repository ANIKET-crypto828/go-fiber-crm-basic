package main

import (
	"github.com/ANIKET-crypto828/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/ANIKET-crypto828/go-fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase(){
		var err error
		database.DBConn, err = gorm.Open("sqlite3", "leads.db")
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Println("Connection opened to database")
		database.DBConn.AutoMigrate(&lead.Lead{})
		fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}