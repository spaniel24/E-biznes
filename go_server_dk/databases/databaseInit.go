package databases

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go_server_dk/models"
)

var db *gorm.DB
var err error

func InitDatabase() {
	db, err = gorm.Open("sqlite3", "shop.db")

	if err != nil {
		panic("Error while connecting to DB")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Category{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&models.Product{Name: "Sluchawki 1", Description: "Opis sluchawek 1", Category: "Headphones", Price: 300, ImageUrl: "https://media.istockphoto.com/photos/blue-headphones-isolated-on-a-white-background-picture-id860853774?s=612x612"})
	db.Create(&models.Product{Name: "Sluchawki 2", Description: "Opis sluchawek 2", Category: "Headphones", Price: 305, ImageUrl: "https://media.istockphoto.com/photos/blue-headphones-isolated-on-a-white-background-picture-id860853774?s=612x612"})
	db.Create(&models.Product{Name: "Sluchawki 3", Description: "Opis sluchawek 3", Category: "Headphones", Price: 500, ImageUrl: "https://media.istockphoto.com/photos/blue-headphones-isolated-on-a-white-background-picture-id860853774?s=612x612"})
	db.Create(&models.Product{Name: "Sluchawki 4", Description: "Opis sluchawek 4", Category: "Headphones", Price: 360, ImageUrl: "https://media.istockphoto.com/photos/blue-headphones-isolated-on-a-white-background-picture-id860853774?s=612x612"})
	db.Create(&models.Product{Name: "Sluchawki 5", Description: "Opis sluchawek 5", Category: "Headphones", Price: 200, ImageUrl: "https://media.istockphoto.com/photos/blue-headphones-isolated-on-a-white-background-picture-id860853774?s=612x612"})

	db.Create(&models.Product{Name: "Klawiatura 1", Description: "Opis klawiatury 1", Category: "Keyboard", Price: 200, ImageUrl: "https://media.istockphoto.com/photos/blue-headphones-isolated-on-a-white-background-picture-id860853774?s=612x612"})
	db.Create(&models.Product{Name: "Klawiatura 2", Description: "Opis klawiatury 2", Category: "Keyboard", Price: 200, ImageUrl: "https://media.istockphoto.com/photos/blue-headphones-isolated-on-a-white-background-picture-id860853774?s=612x612"})

	db.Create(&models.Category{Name: "Headphones"})
	db.Create(&models.Category{Name: "Keyboard"})
}

func GetDatabase() *gorm.DB {
	return db
}
