package app

import (
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

func NewCharacterService(characterRepo CharacterRepository) CharacterService {
	return &characterService{
		characterRepo,
	}
}

func NewMovieService(movieRepo MovieRepository) MovieService {
	return &movieService{
		movieRepo,
	}
}

func NewcommentService(commentRepo CommentRepository) CommentService {
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

func (cSvc *characterService) UpdateCharacter(c *Character) (*Character, error) {
	return cSvc.characterRepo.UpdateCharacter(c)
}

func (cSvc *characterService) DeleteCharacter(id string) error {
	return cSvc.characterRepo.DeleteCharacter(id)
}

func (mSvc *movieService) CreateMovie(m *Movie) (*Movie, error) {

	m.ID = shortid.MustGenerate()
	return mSvc.movieRepo.CreateMovie(m)
}
func (mSvc *movieService) GetMovies() (*[]Movie, error) {
	return mSvc.movieRepo.GetMovies()
}

func (mSvc *movieService) GetMovie(id string) (*Movie, error) {
	return mSvc.movieRepo.GetMovie(id)
}

func (mSvc *movieService) UpdateMovie(m *Movie) (*Movie, error) {
	return mSvc.movieRepo.UpdateMovie(m)
}

func (mSvc *movieService) DeleteMovie(id string) error {
	return mSvc.movieRepo.DeleteMovie(id)
}

func (cSvc *commentService) CreateComment(m *Comment) (*Comment, error) {
	m.CommentID = shortid.MustGenerate()
	return cSvc.commentRepo.CreateComment(m)
}
func (cSvc *commentService) GetComments() (*[]Comment, error) {
	return cSvc.commentRepo.GetComments()
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
