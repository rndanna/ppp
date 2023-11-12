package domain

type ArtistRepository interface {
	GetArtist(id int) (*Artist, error)
	GetArtistByName(name string) (*Artist, error)
	CreateArtist(dto CreateArtistDTO) (int, error)
}

type ArtistUseCase interface {
	GetArtist(id int) (*Artist, error)
	GetArtistByName(name string) (*Artist, error)
	CreateArtist(track Track) (*Track, error)
}

//TODO context
