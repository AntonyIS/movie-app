package app

type CharacterSerializer interface {
	Decode(input []byte) (*Character, error)
	Encode(input *Character) ([]byte, error)
}

type MovieSerializer interface {
	Decode(input []byte) (*Movie, error)
	Encode(input *Movie) ([]byte, error)
}

type CommentSerializer interface {
	Decode(input []byte) (*Comment, error)
	Encode(input *Comment) ([]byte, error)
}
