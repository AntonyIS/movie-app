package redis

import (
	"encoding/json"

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

func NewRedisRepository(redisURL string) (app.CommentRepository, error) {
	repo := &redisRepository{}

	client, err := newRedisClient(redisURL)

	if err != nil {
		return nil, errors.Wrap(err, "comment.NewRedisReposiory")
	}

	repo.client = client
	return repo, nil
}

func (r redisRepository) CreateComment(c *app.Comment) (*app.Comment, error) {

	data := map[string]interface{}{
		"comment_id":   c.CommentID,
		"message":      c.Message,
		"commentor_ip": c.CommentorIP,
	}
	rawmsg, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Comment.Create")
	}

	_, err = r.client.HSet("comments", c.CommentID, rawmsg).Result()

	if err != nil {
		return nil, errors.Wrap(err, "repository.Comment.Create")
	}

	return c, nil
}

func (r redisRepository) GetComment(id string) (*app.Comment, error) {
	comments, err := r.client.HGetAll("comments").Result()

	if err != nil {
		return nil, errors.Wrap(app.ErrorInvalidItem, "repository.Comment.GetComment")
	}
	comment := &app.Comment{}
	err = json.Unmarshal([]byte(comments[id]), comment)
	if err != nil {
		return nil, errors.Wrap(app.ErrorInvalidItem, "repository.Comment.GetComment")
	}

	return comment, nil
}

func (r redisRepository) GetComments() (*[]app.Comment, error) {
	data, err := r.client.HGetAll("Comments").Result()
	Comments := []app.Comment{}
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return &Comments, nil
	}

	for _, Comment := range data {
		res := app.Comment{}
		err := json.Unmarshal([]byte(Comment), &res)
		if err != nil {
			return nil, errors.Wrap(app.ErrorInvalidItem, "repository.Comment.GetComments")
		}
		Comments = append(Comments, res)
	}

	return &Comments, nil
}

func (r redisRepository) UpdateComment(c *app.Comment) (*app.Comment, error) {
	comments, err := r.client.HGetAll("comments").Result()

	if err != nil {
		return nil, errors.Wrap(app.ErrorInvalidItem, "repository.UpdateComment.Update")
	}
	res := &app.Comment{}
	err = json.Unmarshal([]byte(comments[c.CommentID]), res)

	if err != nil {
		return nil, errors.Wrap(app.ErrorInvalidItem, "repository.UpdateComment.Update")
	}

	res.Message = c.Message

	rawmsg, err := json.Marshal(res)
	if err != nil {
		return nil, errors.Wrap(app.ErrorInvalidItem, "repository.UpdateComment.Update")
	}

	found, err := r.client.HSet("comments", c.CommentID, rawmsg).Result()

	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Store")
	}
	if found {
		return res, nil
	}
	return nil, nil

}

func (r redisRepository) DeleteComment(id string) error {
	_, err := r.client.HDel("comments", id).Result()
	if err != nil {
		return errors.Wrap(err, "repository.Comment.DeleteComment")
	}
	return nil
}
