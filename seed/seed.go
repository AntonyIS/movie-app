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
)

func PostCharacters() {
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

	repo, err := redis.NewRedisRepository("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected: ", err)
	}

	switch x := d.(type) {
	case []interface{}:

		for _, e := range x {
			d := e.(map[string]interface{})
			var character app.Character

			character.ID = d["id"].(string)
			character.Name = d["name"].(string)
			character.Height = d["height"].(string)
			character.Mass = d["mass"].(string)
			character.HairColor = d["hair_color"].(string)
			character.SkinColor = d["skin_color"].(string)
			character.EyeColor = d["eye_color"].(string)
			character.BirthYear = d["birth_year"].(string)
			character.Gender = d["gender"].(string)
			character.Homeworld = d["homeworld"].(string)
			character.FilmURLs = d["films"].([]string)
			character.SpeciesURLs = d["species"].([]string)
			character.VehicleURLs = d["vehicles"].([]string)
			character.StarshipURLs = d["starships"].([]string)
			character.Created = d["created"].(string)
			character.Edited = d["edited"].(string)
			character.URL = d["url"].(string)
			repo.CreateCharacter(&character)

		}
		fmt.Println("Done seeding")
	default:
		fmt.Printf("I don't know how to handle %T\n", d)
	}

}
