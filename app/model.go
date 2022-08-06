package app

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Character struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Height       string   `json:"height"`
	Mass         string   `json:"mass"`
	HairColor    string   `json:"hair_color"`
	SkinColor    string   `json:"skin_color"`
	EyeColor     string   `json:"eye_color"`
	BirthYear    string   `json:"birth_year"`
	Gender       string   `json:"gender"`
	Homeworld    string   `json:"homeworld"`
	FilmURLs     []string `json:"films"`
	SpeciesURLs  []string `json:"species"`
	VehicleURLs  []string `json:"vehicles"`
	StarshipURLs []string `json:"starships"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	URL          string   `json:"url"`
}

type Movie struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	OpeningCrawl string    `json:"opening_crawl"`
	Comments     []Comment `json:"comments"`
	Created      string    `json:"created"`
	Edited       string    `json:"edited"`
}

type Comment struct {
	CommentID   string `json:"comment_id" gorm:"type:text"`
	MovieID     string `json:"movie_id" gorm:"type:text"`
	Message     string `json:"message" gorm:"type:text"`
	CommentorIP string `json:"commentor_ip" gorm:"type:text"`
	Created     string `json:"created" gorm:"type:text"`
	Edited      string `json:"edited" gorm:"type:text"`
}
