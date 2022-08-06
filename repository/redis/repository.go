package redis

import (
	"encoding/json"
	"fmt"

	"github.com/AntonyIS/movie-app/app"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
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
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "character.NewRedisReposiory")
	}
	repo.client = client
	return repo, nil
}

func (r redisRepository) CreateCharacter(c *app.Character) (*app.Character, error) {
	data := map[string]interface{}{
		"id":         c.ID,
		"name":       c.Name,
		"height":     c.Height,
		"mass":       c.Mass,
		"hair_color": c.HairColor,
		"skin_color": c.SkinColor,
		"eye_color":  c.EyeColor,
		"birth_year": c.BirthYear,
		"gender":     c.Gender,
		"homeworld":  c.Homeworld,
		"films":      c.FilmURLs,
		"species":    c.SpeciesURLs,
		"vehicles":   c.VehicleURLs,
		"starships":  c.StarshipURLs,
		"created":    c.Created,
		"edited":     c.Edited,
		"url":        c.URL,
	}
	defer r.client.Close()

	charSlice := make(map[string]interface{})
	val, err := r.client.Get("characters").Result()
	if err != nil {
		err = r.client.Set("characters", charSlice, 0).Err()
		if err != nil {
			fmt.Println(err, "EEEEEErSSZDSEFVCSDEFDCSDFc")
		}
	}

	fmt.Println(val, "EEEEEEEEEEEEEEEEEE")

	// err := rdb.Set(ctx, "characters", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	key := r.generateKey(c.ID)
	// fmt.Println(key)
	_, err = r.client.HMSet(key, data).Result()
	if err != nil {
		return nil, err
	}

	// v, err := r.client.HGetAll(key).Result()
	// if err != nil {

	// 	return nil, err
	// }
	// fmt.Println(v)
	return c, nil
}

func (r redisRepository) GetCharacter(id string) (*app.Character, error) {

	// key := r.generateKey(id)
	data, err := r.client.HGetAll("characters:wXveCwz4g").Result()

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, err
	}
	c := app.Character{}
	m2 := make(map[string]interface{}, len(data))
	for k, v := range data {
		m2[k] = v
	}

	c.ID = m2["id"].(string)
	c.Name = m2["name"].(string)
	c.Height = m2["height"].(string)
	c.Mass = m2["mass"].(string)
	c.HairColor = m2["hair_color"].(string)
	c.SkinColor = m2["skin_color"].(string)
	c.EyeColor = m2["eye_color"].(string)
	c.BirthYear = m2["birth_year"].(string)
	c.Gender = m2["gender"].(string)
	c.Homeworld = m2["homeworld"].(string)
	c.FilmURLs = m2["films"].([]string)
	c.SpeciesURLs = m2["species"].([]string)
	c.VehicleURLs = m2["vehicles"].([]string)
	c.StarshipURLs = m2["starships"].([]string)
	c.Created = m2["created"].(string)
	c.Edited = m2["edited"].(string)
	c.URL = m2["url"].(string)

	return &c, nil

}

func (r redisRepository) GetCharacters() (*[]app.Character, error) {
	data, err := r.client.HGetAll("characters:").Result()
	if err != nil {

		return nil, err
	}
	v, err := r.client.HGetAll("key").Result()
	if err != nil {

		return nil, err
	}
	fmt.Println(v)
	characters := []app.Character{}
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return &characters, nil
	}

	for _, Character := range data {
		res := app.Character{}
		err := json.Unmarshal([]byte(Character), &res)
		if err != nil {
			return nil, errors.Wrap(app.ErrorInvalidItem, "repository.Character.GetCharacters")
		}
		characters = append(characters, res)
	}

	return &characters, nil
}

func (r redisRepository) UpdateCharacter(c *app.Character) (*app.Character, error) {
	key := r.generateKey(c.ID)
	data, err := r.client.HGetAll(key).Result()

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, err
	}

	m2 := make(map[string]interface{}, len(data))
	for k, v := range data {
		m2[k] = v
	}
	m2["name"] = c.Name
	m2["height"] = c.Height
	m2["mass"] = c.Mass
	m2["hair_color"] = c.HairColor
	m2["skin_color"] = c.SkinColor
	m2["eye_color"] = c.EyeColor
	m2["birth_year"] = c.BirthYear
	m2["gender"] = c.Gender
	m2["films"] = c.FilmURLs
	m2["species"] = c.SpeciesURLs
	m2["vehicles"] = c.VehicleURLs
	m2["starships"] = c.StarshipURLs
	m2["created"] = c.Created
	m2["edited"] = c.Edited
	m2["url"] = c.URL

	_, err = r.client.HMSet(key, m2).Result()
	if err != nil {
		return nil, err
	}

	return c, nil

}

func (r redisRepository) DeleteCharacter(id string) error {
	key := r.generateKey(id)
	_, err := r.client.HDel(key).Result()
	if err != nil {
		return err
	}
	return nil
}
func (r *redisRepository) generateKey(id string) string {
	return fmt.Sprintf("characters:%s", id)
}
