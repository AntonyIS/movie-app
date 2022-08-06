package app

type CharacterRepository interface {
	GetCharacter(id string) (*Character, error)
	GetCharacters() (*[]Character, error)
	CreateCharacter(c *Character) (*Character, error)
	UpdateCharacter(c *Character) (*Character, error)
	DeleteCharacter(id string) error
}

type MovieRepository interface {
	GetMovie(id string) (*Movie, error)
	GetMovies() (*[]Movie, error)
	CreateMovie(c *Movie) (*Movie, error)
	UpdateMovie(c *Movie) (*Movie, error)
	DeleteMovie(id string) error
	GetMovieComments(id string) (*[]Comment, error)
	GetMovieCharacters(id string) (*[]Character, error)
}

type CommentRepository interface {
	GetComment(id string) (*Comment, error)
	GetComments() (*[]Comment, error)
	CreateComment(c *Comment) (*Comment, error)
	UpdateComment(c *Comment) (*Comment, error)
	DeleteComment(id string) error
}
