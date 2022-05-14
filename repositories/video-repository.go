package repositories

import (
	"golang-gin-poc/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VideoRepository interface {
	Create(video entities.Video)
	Update(video entities.Video)
	Delete(video entities.Video)
	FindAll() []entities.Video
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	dsn := "host=localhost user=root password=secret dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entities.Video{}, &entities.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) Create(video entities.Video) {
	db.connection.Create(&video)
}
func (db *database) Update(video entities.Video) {
	db.connection.Save(&video)
}
func (db *database) Delete(video entities.Video) {
	db.connection.Delete(&video)
}
func (db *database) FindAll() []entities.Video {
	var videos []entities.Video
	db.connection.Preload(clause.Associations).Find(&videos)
	return videos
}
