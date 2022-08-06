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

func PostCharacters() interface{} {

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
	repo, err := redis.NewRedisRepository("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected: ", err)
	}
	test := d.([]interface{})

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

		repo.CreateCharacter(&c)

	}

	return results
}

func marshal(d []interface{}) []string {
	s := []string{}
	for _, k := range d {
		s = append(s, k.(string))
	}
	return s

}
