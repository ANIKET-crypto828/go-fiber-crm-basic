package lead

import (
	"github.com/ANIKET-crypto828/go-fiber-crm-basic/database"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct{
	gorm.Model
	Name string
	Company string
	Email string
	Phone int
}

func GetLeads(){

}

func GetLead(){
	
}

func NewLead(){
	
}

func DeleteLead(){
	
}