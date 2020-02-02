package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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

	var config Configuration
	json.Unmarshal([]byte(conf), &config)

	return config
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

	if err != nil {
		fmt.Println(err)
	}
	return movies, err
}

// GetMovieDetailsById get the movie by id
func GetMovieDetailsById(id string) {
	movie := &Movie{}
	resp, err := tmdbAPICall("/movie/"+id, "")

	if err != nil {
		fmt.Println("Could not get movie.")
		return
	}

	json.Unmarshal([]byte(resp), &movie)

	return movie

}

func updatePaths(movie Movie) Movie {
	if len(movie.PosterPath) != 0 {
		movie.PosterPath = config.Images.BaseURL + config.Images.PosterSizes[4] + movie.PosterPath
	}

	if len(movie.BackdropPath) != 0 {
		movie.BackdropPath = config.Images.BaseURL + config.Images.BackdropSizes[2] + movie.BackdropPath
	}

	return movie
}
