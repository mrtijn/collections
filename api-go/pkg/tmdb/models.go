package tmdb

// Movie struct
type Movie struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	ReleaseDate  string `json:"release_date"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

// Configuration struct
type Configuration struct {
	Images struct {
		BaseURL       string    `json:"base_url"`
		PosterSizes   [7]string `json:"poster_sizes"`
		BackdropSizes [4]string `json:"backdrop_sizes"`
	}
}

type searchResults struct {
	Page         uint
	Results      []Movie
	TotalPages   uint
	TotalResults uint
}
