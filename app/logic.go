package app

import (
	"errors"
	"strconv"

	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

var (
	ErrorItemNotFound = errors.New("item not found")
	ErrorInvalidItem  = errors.New("item invalid")
)

type characterService struct {
	characterRepo CharacterRepository
}
type movieService struct {
	movieRepo MovieRepository
}

type commentService struct {
	commentRepo CommentRepository
}

func NewCharacterRepository(characterRepo CharacterRepository) CharacterService {
	return &characterService{
		characterRepo,
	}
}

func NewMovieRepository(movieRepo MovieRepository) MovieService {
	return &movieService{
		movieRepo,
	}
}

func NewcommentRepository(commentRepo CommentRepository) CommentService {
	return &commentService{
		commentRepo,
	}
}

func (cSvc *characterService) CreateCharacter(c *Character) (*Character, error) {
	c.ID = shortid.MustGenerate()
	return cSvc.characterRepo.CreateCharacter(c)
}
func (cSvc *characterService) GetCharacters() (*[]Character, error) {
	return cSvc.characterRepo.GetCharacters()
}

func (cSvc *characterService) GetCharacter(id string) (*Character, error) {
	return cSvc.characterRepo.GetCharacter(id)
}

func (cSvc *characterService) UpdateCharacter(id string) (*Character, error) {
	return cSvc.characterRepo.UpdateCharacter(id)
}

func (cSvc *characterService) DeleteCharacter(id string) error {
	return cSvc.characterRepo.DeleteCharacter(id)
}

func (mSvc *movieService) CreateMovie(m *Movie) (*Movie, error) {
	id, err := strconv.Atoi(shortid.MustGenerate())
	if err != nil {
		return nil, errors.Wrap(app.ErrorInvalidItem, "repository.Todo.Update")
	}
	m.EpisodeID = id
	return mSvc.movieRepo.CreateMovie(m)
}
func (mSvc *movieService) GetMovies() (*[]Movie, error) {
	return mSvc.movieRepo.GetMovies()
}

func (mSvc *movieService) GetMovie(id string) (*Movie, error) {
	return mSvc.movieRepo.GetMovie(id)
}

func (mSvc *movieService) UpdateMovie(id string) (*Movie, error) {
	return mSvc.movieRepo.UpdateMovie(id)
}

func (mSvc *movieService) DeleteMovie(id string) error {
	return mSvc.movieRepo.DeleteMovie(id)
}

func (cSvc *commentService) CreateComment(m *Comment) (*Comment, error) {
	m.CommentID = shortid.MustGenerate()
	return cSvc.commentRepo.CreateComment(m)
}
func (cSvc *commentService) GetComments() (*[]Comment, error) {
	return cSvc.GetComments()
}

func (cSvc *commentService) GetComment(id string) (*Comment, error) {
	return cSvc.commentRepo.GetComment(id)
}

func (cSvc *commentService) UpdateComment(c *Comment) (*Comment, error) {
	return cSvc.commentRepo.UpdateComment(c)
}

func (cSvc *commentService) DeleteComment(id string) error {
	return cSvc.commentRepo.DeleteComment(id)
}
