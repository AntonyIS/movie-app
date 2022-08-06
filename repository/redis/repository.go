package redis

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AntonyIS/movie-app/app"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type redisRepository struct {
	Client *redis.Client
}

func NewRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(redisURL string) (app.CharacterRepository, error) {
	repo := &redisRepository{}
	client, err := NewRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "character.NewRedisReposiory")
	}
	repo.Client = client
	return repo, nil
}

func (r redisRepository) CreateCharacter(c *app.Character) (*app.Character, error) {

	json, err := json.Marshal(c)
	if err != nil {
		log.Fatal("Error adding new character")
	}
	r.Client.HSet("characters", c.ID, json)
	defer r.Client.Close()
	return c, nil
}

func (r redisRepository) GetCharacter(id string) (*app.Character, error) {
	data, err := r.Client.HGet("characters", id).Result()

	if data == " " {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	var character app.Character
	err = json.Unmarshal([]byte(data), &character)

	if err != nil {
		return nil, err
	}

	return &character, nil

}

func (r redisRepository) GetCharacters() (*[]app.Character, error) {
	data, err := r.Client.HGetAll("characters").Result()
	if err != nil {
		return nil, err
	}

	characters := []app.Character{}
	for _, value := range data {
		character := app.Character{}
		err := json.Unmarshal([]byte(value), &character)
		if err != nil {
			log.Fatal(err)
		}
		characters = append(characters, character)
	}

	return &characters, nil
}

func (r redisRepository) UpdateCharacter(c *app.Character) (*app.Character, error) {

	data, err := r.Client.HGet("characters", c.ID).Result()
	
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	var character app.Character
	err = json.Unmarshal([]byte(data), &character)

	if err != nil {
		return nil, err
	}

	character.Name = c.Name
	character.Height = c.Height
	character.Mass = c.Mass
	character.HairColor = c.HairColor
	character.SkinColor = c.SkinColor
	character.EyeColor = c.EyeColor
	character.BirthYear = c.BirthYear
	character.Gender = c.Gender
	character.Homeworld = c.Homeworld
	character.FilmURLs = c.FilmURLs
	character.SpeciesURLs = c.SpeciesURLs
	character.VehicleURLs = c.VehicleURLs
	character.StarshipURLs = c.StarshipURLs
	character.Created = c.Created
	character.Edited = c.Edited
	character.URL = c.URL

	json, err := json.Marshal(c)
	if err != nil {
		log.Fatal("Error adding new character")
	}
	r.Client.HSet("characters", c.ID, json)
	return &character, nil

}

func (r redisRepository) DeleteCharacter(id string) error {
	err := r.Client.HDel("characters", id)

	if fmt.Sprintf("%T", err) == "IntCMD" {
		return nil
	}

	if err != nil {
		return nil
	}

	return nil

}
