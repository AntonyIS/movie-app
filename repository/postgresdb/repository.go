package postgresdb

import (
	"fmt"
	"log"
	"os"

	"github.com/AntonyIS/movie-app/app"
	redis "github.com/AntonyIS/movie-app/repository/redis"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type postgresRepository struct {
	DB        *gorm.DB
	TableName string
}

func postgresDB() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := "5432"
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password))

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate(app.Comment{})

	return db

}

func NewRepostory() app.CommentRepository {
	return postgresRepository{
		DB: postgresDB(),
	}
}

func (p postgresRepository) CreateComment(c *app.Comment) (*app.Comment, error) {
	redis, err := redis.NewMovieRedisRepository("redis://localhost:6379")
	if err != nil {
		log.Fatal("Error ::", err)
	}

	movie, err := redis.GetMovie(c.MovieID)
	if err != nil {
		log.Fatal("Error ::", err)
	}
	movie.Comments = append(movie.Comments, c.URL)

	_, err = redis.UpdateMovie(movie)
	if err != nil {
		log.Fatal("Error ::", err)
	}

	if err := p.DB.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (p postgresRepository) GetComment(id string) (*app.Comment, error) {
	var comment app.Comment
	p.DB.First(&comment, id)
	return &comment, nil
}

func (p postgresRepository) GetComments() (*[]app.Comment, error) {
	comments := []app.Comment{}
	p.DB.Find(&comments)
	return &comments, nil
}

func (p postgresRepository) UpdateComment(c *app.Comment) (*app.Comment, error) {
	if result := p.DB.Save(c); result.Error != nil {
		return nil, errors.Wrap(result.Error, "repository.Comment.UpdateComment")
	}
	return c, nil
}

func (p postgresRepository) DeleteComment(id string) error {
	comment := app.Comment{}
	if result := p.DB.Find(&comment); result.Error != nil {
		return errors.Wrap(result.Error, "repository.Comment.DeleteComment")
	}
	if err := p.DB.Where("id = ? ", id).Delete(&comment).Error; err != nil {
		return errors.Wrap(err, "repository.Comment.DeleteComment")
	}
	return nil
}
