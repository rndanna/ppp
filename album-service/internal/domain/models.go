package domain

type CreateAlbumDTO struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	ArtistID int    `json:"artist_id"`
}

type Album struct {
	ID     *int    `json:"id,omitempty"`
	Title  string  `json:"title"`
	URL    *string `json:"url,omitempty"`
	Artist *Artist `json:"artist"`
}

type Artist struct {
	ID   *int    `json:"id,omitempty"`
	Name string  `json:"name"`
	URL  *string `json:"url,omitempty"`
}

type Tag struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Track struct {
	ID        *int    `json:"id,omitempty"`
	Name      string  `json:"name"`
	URL       *string `json:"url,omitempty"`
	Listeners *int    `json:"listeners,omitempty"`
	Playcount *int    `json:"playcount,omitempty"`
	Artist    *Artist `json:"artist"`
	Album     *Album  `json:"album"`
	Tags      []*Tag  `json:"tags,omitempty"`
}
