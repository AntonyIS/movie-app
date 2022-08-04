package postgresdb

import (
	"fmt"
	"log"
	"os"

	"github.com/AntonyIS/movie-app/app"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type postgresRepository struct {
	DB        *gorm.DB
	TableName string
}

func PostgresDB() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password))

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate(app.Character{})

	return db

}

func NewRepostory() app.CharacterRepository {
	return postgresRepository{
		DB: PostgresDB(),
	}
}

func (p postgresRepository) CreateCharacter(c *app.Character) (*app.Character, error) {
	if err := p.DB.Create(&c).Error; err != nil {
		return nil, errors.Wrap(err, "repository.Character.CreateCharacter")
	}
	return c, nil
}

func (p postgresRepository) GetCharacter(id string) (*app.Character, error) {
	var character app.Character
	response := p.DB.First(&character, "id= ?", id)
	if response.RowsAffected == 0 {
		return nil, errors.Wrap(errors.New("character not found"), "repository.Character.CreateCharacter")
	}

	return &character, nil
}

func (p postgresRepository) GetCharacters() (*[]app.Character, error) {
	characters := []app.Character{}
	if result := p.DB.Find(&characters); result.Error != nil {
		return nil, errors.Wrap(result.Error, "repository.Character.GetCharacters")
	}
	return &characters, nil
}

func (p postgresRepository) UpdateCharacter(c *app.Character) (*app.Character, error) {
	if result := p.DB.Save(c); result.Error != nil {
		return nil, errors.Wrap(result.Error, "repository.Character.UpdateCharacter")
	}
	return c, nil
}

func (p postgresRepository) DeleteCharacter(id string) error {
	character := app.Character{}
	if result := p.DB.Find(&character); result.Error != nil {
		return errors.Wrap(result.Error, "repository.Character.Deletecharacter")
	}
	if err := p.DB.Where("id = ? ", id).Delete(&character).Error; err != nil {
		return errors.Wrap(err, "repository.Character.Deletecharacter")
	}
	return nil
}
