package app

type CharacterService interface {
	GetCharacter(id string) (*Character, error)
	GetCharacters() (*[]Character, error)
	CreateCharacter(c *Character) (*Character, error)
	UpdateCharacter(c *Character) (*Character, error)
	DeleteCharacter(id string) error
}

type MovieService interface {
	GetMovie(id string) (*Movie, error)
	GetMovies() (*[]Movie, error)
	CreateMovie(m *Movie) (*Movie, error)
	UpdateMovie(m *Movie) (*Movie, error)
	DeleteMovie(id string) error
}

type CommentService interface {
	GetComment(id string) (*Comment, error)
	GetComments() (*[]Comment, error)
	CreateComment(c *Comment) (*Comment, error)
	UpdateComment(c *Comment) (*Comment, error)
	DeleteComment(id string) error
}
