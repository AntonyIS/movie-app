package seed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/AntonyIS/movie-app/app"
	"github.com/AntonyIS/movie-app/repository/redis"
	"github.com/teris-io/shortid"
)

func SeedCharacters() interface{} {

	url := "https://swapi.dev/api/people"

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := string(res)

	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonMap)
	if err != nil {
		log.Fatal(err)
	}

	d := jsonMap["results"]
	results := d
	repo, err := redis.NewRedisClient("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected: ", err)
	}
	test := d.([]interface{})

	defer repo.Close()

	for _, v := range test {
		rep := v.(map[string]interface{})

		var c app.Character
		c.ID = shortid.MustGenerate()
		c.Name = rep["name"].(string)
		c.Height = rep["height"].(string)
		c.Mass = rep["mass"].(string)
		c.HairColor = rep["hair_color"].(string)
		c.SkinColor = rep["skin_color"].(string)
		c.EyeColor = rep["eye_color"].(string)
		c.BirthYear = rep["birth_year"].(string)
		c.Gender = rep["gender"].(string)
		c.Homeworld = rep["homeworld"].(string)
		c.FilmURLs = marshal(rep["films"].([]interface{}))
		c.SpeciesURLs = marshal(rep["species"].([]interface{}))
		c.VehicleURLs = marshal(rep["vehicles"].([]interface{}))
		c.StarshipURLs = marshal(rep["starships"].([]interface{}))
		c.Created = rep["created"].(string)
		c.Edited = rep["edited"].(string)
		c.URL = rep["url"].(string)

		json, err := json.Marshal(c)
		if err != nil {
			log.Fatal("Error adding new character")
		}

		repo.HSet("characters", c.ID, json)

	}

	return results
}

func SeedMovies() interface{} {

	url := "https://swapi.dev/api/films"

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := string(res)

	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonMap)
	if err != nil {
		log.Fatal(err)
	}

	d := jsonMap["results"]

	repo, err := redis.NewRedisClient("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected: ", err)
	}
	movieData := d.([]interface{})

	movies := []app.Movie{}
	for _, v := range movieData {
		rep := v.(map[string]interface{})

		var m app.Movie
		m.ID = shortid.MustGenerate()
		m.Title = rep["title"].(string)
		m.OpeningCrawl = rep["opening_crawl"].(string)
		m.Comments = []string{}
		m.Created = rep["created"].(string)
		m.Characters = marshal(rep["characters"].([]interface{}))
		m.Edited = rep["edited"].(string)

		json, err := json.Marshal(m)
		if err != nil {
			log.Fatal("Error adding new character")
		}

		repo.HSet("movies", m.ID, json)

		movies = append(movies, m)

	}

	return movies
}

func marshal(d []interface{}) []string {
	s := []string{}
	for _, k := range d {
		s = append(s, k.(string))
	}
	return s

}
