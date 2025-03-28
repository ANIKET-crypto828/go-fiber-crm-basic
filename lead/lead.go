package lead

import (
	"github.com/ANIKET-crypto828/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct{
	gorm.Model
	Name string  `json:"name"`
	Company string  `json:"company"`
	Email string   `json:"email"`
	Phone int      `json:"phone"`
}

func GetLeads(c *fiber.Ctx){
		db := database.DBConn
		var leads []Lead
		db.Find(&leads)
		c.JSON(leads)
}

func GetLead(c *fiber.Ctx){
		id := c.Params("id")
		db := database.DBConn
		var lead Lead
		db.First(&lead, id)
		c.JSON(lead)
}

func NewLead(c *fiber.Ctx){
		db := database.DBConn
		lead := new(Lead)
		if err := c.BodyParser(lead); err != nil{
			//c.Status(fiber.StatusBadRequest).SendString(err.Error())
			c.Status(503).Send(err)
			return
		}
		db.Create(&lead)
		c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
  db.First(&lead, id)
	if lead.Name == ""{
		c.Status(503).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead Successfully Deleted")
}