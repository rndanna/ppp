package domain

type AlbumRepository interface {
	GetAlbumByTitle(title string) (*Album, error)
	CreateAlbum(dto CreateAlbumDTO) (int, error)
}

type AlbumUseCase interface {
	GetAlbumByTitle(title string) (*Album, error)
	CreateAlbum(track Track) (*Track, error)
}

//TODO context
