package database

import (
	"football-forum/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:61611616@tcp(127.0.0.1:3306)/football?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı: ", err)
	}

	log.Println("Veritabanı bağlantısı başarılı")

	err = DB.AutoMigrate(&models.User{}, &models.Topic{}, &models.Comment{}, &models.Category{})
	if err != nil {
		log.Fatal("Veritabanı migrasyonu başarısız: ", err)
	}

	log.Println("Veritabanı migrasyonu başarılı")
}

// GetDB veritabanı bağlantısını döndüren fonksiyon
func GetDB() *gorm.DB {
	return DB
}
