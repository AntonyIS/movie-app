package app

type CharacterRepository interface {
	GetCharacter(id string) (*Character, error)
	GetCharacters() (*[]Character, error)
	CreateCharacter(c *Character) (*Character, error)
	UpdateCharacter(id string) (*Character, error)
	DeleteCharacter(id string) error
}

type MovieRepository interface {
	GetMovie(id string) (*Movie, error)
	GetMovies() (*[]Movie, error)
	CreateMovie(c *Movie) (*Movie, error)
	UpdateMovie(id string) (*Movie, error)
	DeleteMovie(id string) error
}

type CommentRepository interface {
	GetComment(id string) (*Comment, error)
	GetComments() (*[]Comment, error)
	CreateComment(c *Comment) (*Comment, error)
	UpdateComment(id string) (*Comment, error)
	DeleteComment(id string) error
}
