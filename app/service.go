package app

type CharacterService interface {
	GetCharacter(id string) (*Character, error)
	GetCharacters() (*[]Character, error)
	CreateCharacter(c *Character) (*Character, error)
	UpdateCharacter(id string) (*Character, error)
	DeleteCharacter(id string) error
}

type MovieService interface {
	GetMovie(id string) (*Movie, error)
	GetMovies() (*[]Movie, error)
	CreateMovie(c *Movie) (*Movie, error)
	UpdateMovie(id string) (*Movie, error)
	DeleteMovie(id string) error
}

type CommentService interface {
	GetComment(id string) (*Comment, error)
	GetComments() (*[]Comment, error)
	CreateComment(c *Comment) (*Comment, error)
	UpdateComment(c *Comment) (*Comment, error)
	DeleteComment(id string) error
}
