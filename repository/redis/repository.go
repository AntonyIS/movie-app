package redis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

func NewCharacterRedisRepository(redisURL string) (app.CharacterRepository, error) {
	repo := &redisRepository{}
	client, err := NewRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "character.NewRedisReposiory")
	}
	repo.Client = client
	return repo, nil
}

func NewMovieRedisRepository(redisURL string) (app.MovieRepository, error) {
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

func (r redisRepository) CreateMovie(m *app.Movie) (*app.Movie, error) {
	fmt.Println(m)
	fmt.Println(m.ID)
	json, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Error adding new Movie")
	}
	r.Client.HSet("movies", m.ID, json)
	defer r.Client.Close()
	return m, nil
}

func (r redisRepository) GetMovie(id string) (*app.Movie, error) {
	data, err := r.Client.HGet("movies", id).Result()

	if data == " " {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	var movie app.Movie
	err = json.Unmarshal([]byte(data), &movie)

	if err != nil {
		return nil, err
	}

	return &movie, nil

}

func (r redisRepository) GetMovies() (*[]app.Movie, error) {
	data, err := r.Client.HGetAll("movies").Result()

	if err != nil {
		return nil, err
	}

	movies := []app.Movie{}
	for _, value := range data {
		movie := app.Movie{}
		err := json.Unmarshal([]byte(value), &movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return &movies, nil
}

func (r redisRepository) UpdateMovie(m *app.Movie) (*app.Movie, error) {

	data, err := r.Client.HGet("movies", m.ID).Result()

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	var Movie app.Movie
	err = json.Unmarshal([]byte(data), &Movie)

	if err != nil {
		return nil, err
	}

	Movie.Title = m.Title
	Movie.OpeningCrawl = m.OpeningCrawl
	Movie.Created = m.Created
	Movie.Edited = m.Edited

	json, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Error adding new Movie")
	}
	r.Client.HSet("movies", m.ID, json)
	return &Movie, nil

}

func (r redisRepository) DeleteMovie(id string) error {
	err := r.Client.HDel("Movies", id)

	if fmt.Sprintf("%T", err) == "IntCMD" {
		return nil
	}

	if err != nil {
		return nil
	}

	return nil

}

func (r redisRepository) GetMovieComments(id string) (*[]app.Comment, error) {
	data, err := r.GetMovie(id)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	comments := []app.Comment{}
	URL := data.Comments

	for _, link := range URL {
		response, err := http.Get(link)

		comment := app.Comment{}
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		res, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(res, &comment)
		if err != nil {
			log.Fatal(err)
		}

		comments = append(comments, comment)

	}

	return &comments, nil
}

func (r redisRepository) GetMovieCharacters(id string) (*[]app.Character, error) {
	movie, err := r.GetMovie(id)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	characters := []app.Character{}
	URL := movie.Characters

	for _, link := range URL {
		response, err := http.Get(link)

		character := app.Character{}
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		res, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(res, &character)
		if err != nil {
			log.Fatal(err)
		}

		characters = append(characters, character)

	}

	return &characters, nil
}
