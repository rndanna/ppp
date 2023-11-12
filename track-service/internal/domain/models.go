package domain

type CreateTrackDTO struct {
	Name      string  `json:"name"`
	URL       *string `json:"url,omitempty"`
	Listeners *int    `json:"listeners,omitempty"`
	Playcount *int    `json:"playcount,omitempty"`
	AlbumID   int     `json:"album_id"`
}

type CreateTagDTO struct {
	Name string  `json:"name"`
	URL  *string `json:"url,omitempty"`
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

type Tag struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
