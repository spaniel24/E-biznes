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
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Payment{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//db.Create(&models.Product{Name: "Sluchawki 1", Description: "Opis sluchawek 1", Category: "Headphones", Price: 300, ImageUrl: "https://media.istockphoto.com/photos/blue-headphones-isolated-on-a-white-background-picture-id860853774?s=612x612"})
	//db.Create(&models.Product{Name: "Sluchawki 2", Description: "Opis sluchawek 2", Category: "Headphones", Price: 305, ImageUrl: "https://image.shutterstock.com/image-photo/highquality-headphones-on-white-background-600w-1574611990.jpg"})
	//db.Create(&models.Product{Name: "Sluchawki 3", Description: "Opis sluchawek 3", Category: "Headphones", Price: 500, ImageUrl: "https://image.shutterstock.com/image-photo/wireless-headphones-isolated-on-white-600w-1402966736.jpg"})
	//db.Create(&models.Product{Name: "Sluchawki 4", Description: "Opis sluchawek 4", Category: "Headphones", Price: 360, ImageUrl: "https://image.shutterstock.com/image-photo/black-headphones-isolated-on-white-600w-763855648.jpg"})
	//db.Create(&models.Product{Name: "Sluchawki 5", Description: "Opis sluchawek 5", Category: "Headphones", Price: 200, ImageUrl: "https://image.shutterstock.com/image-illustration/3d-rendering-headphones-isolated-on-600w-786005032.jpg"})
	//
	//db.Create(&models.Product{Name: "Klawiatura 1", Description: "Opis klawiatury 1", Category: "Keyboard", Price: 200, ImageUrl: "https://image.shutterstock.com/image-vector/black-computer-qwerty-keyboard-simple-600w-1687409758.jpg"})
	//db.Create(&models.Product{Name: "Klawiatura 2", Description: "Opis klawiatury 2", Category: "Keyboard", Price: 200, ImageUrl: "https://image.shutterstock.com/image-photo/computer-keyboard-isolated-on-white-600w-41303413.jpg"})
	//
	//db.Create(&models.Product{Name: "Myszka 1", Description: "Opis myszki 1", Category: "Mouse", Price: 100, ImageUrl: "https://image.shutterstock.com/image-photo/gray-modern-computer-wireless-mouse-600w-1917358601.jpg"})
	//db.Create(&models.Product{Name: "Myszka 2", Description: "Opis myszki 2", Category: "Mouse", Price: 300, ImageUrl: "https://image.shutterstock.com/z/stock-photo-wireless-computer-mouse-on-white-background-1867466266.jpg"})
	//
	//db.Create(&models.Category{Name: "Headphones"})
	//db.Create(&models.Category{Name: "Keyboard"})
	//db.Create(&models.Category{Name: "Mouse"})
}

func GetDatabase() *gorm.DB {
	return db
}
