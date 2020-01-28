package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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

var config = GetTmdbConfiguration()

// tmdbAPICall returns a get for the tmdb endpoint
func tmdbAPICall(endpoint, queryParams string) (string, error) {
	resp, err := http.Get("https://api.themoviedb.org/3" + endpoint + "?api_key=" + os.Getenv("TMDBApiKey") + "&" + queryParams)

	if err != nil {
		fmt.Println(err)
	}
	// defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "tmdbApi error", err
	}

	return string(body), err

}

// GetTmdbConfiguration return configuration
func GetTmdbConfiguration() Configuration {

	conf, err := tmdbAPICall("/configuration", "")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(conf)

	var config Configuration
	json.Unmarshal([]byte(conf), &config)
	fmt.Println(config)
	return config
}

// GetMovieData Get Movie data from tmdb
// func GetMovieData(ID string) {
// 	resp, err := tmdbAPICall("/movie/"+ID, "")

// }

// func mapMovie(dataSet map[string]interface{}) (models.Movie, error) {

// }

type searchResults struct {
	Page         uint
	Results      []Movie
	TotalPages   uint
	TotalResults uint
}

// SearchMovie search movie
func SearchMovie(searchTerm string) ([]Movie, error) {

	resp, err := tmdbAPICall("/search/movie", "query="+searchTerm)

	result := searchResults{}
	json.Unmarshal([]byte(resp), &result)

	var movies []Movie
	movies = result.Results

	for i := range movies {
		movies[i] = updatePaths(movies[i])
	}
	// for index := 0; index < len(movies); index++ {
	// 	var movie = &TmdbMovie{}
	// 	movie = movies[index]
	// 	movie = updatePaths(movie)
	// 	// movies[index] = updatePaths(movies[index])
	// }
	if err != nil {
		fmt.Println(err)
	}
	return movies, err
}

// PosterSize postersize string
type PosterSize int

// Fill postersize enum
const (
	W92      PosterSize = 0
	w154     PosterSize = 1
	w185     PosterSize = 2
	w342     PosterSize = 3
	w500     PosterSize = 4
	w780     PosterSize = 5
	original PosterSize = 6
)

func updatePaths(movie Movie) Movie {

	movie.PosterPath = config.Images.BaseURL + config.Images.PosterSizes[w500] + movie.PosterPath
	movie.BackdropPath = config.Images.BaseURL + config.Images.BackdropSizes[2] + movie.BackdropPath
	fmt.Println(config)
	return movie
}
