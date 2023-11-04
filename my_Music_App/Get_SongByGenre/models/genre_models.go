package models

// SongInfo represents information about a song.
type SongInfo struct {
	ID          int    `json:"id"`          // Unique identifier for the song
	Name        string `json:"name"`        // Name of the song
	Artists     string `json:"artists"`     // Names of the artists associated with the song
	Genre       string `json:"genre"`       // Genre of the song
	PublishYear int    `json:"publishyear"` // Year the song was published
	Language    string `json:"language"`    // Language of the song
}
