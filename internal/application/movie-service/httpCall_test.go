package movie

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/solution_case_2/internal/application/movie-service/dto"
)

type input struct {
	url  string
	resp interface{}
}

func TestSearchMovie(t *testing.T) {
	var search dto.SearchResponse
	// var movie []dto.Movie{}
	tests := []struct {
		input
		result interface{}
	}{
		{
			input: input{url: "http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2", resp: search},
			result: dto.SearchResponse{
				Movie: []dto.Movie{
					{
						ImdbID: "tt4853102",
						Title:  "Batman: The Killing Joke",
						Year:   "2016",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt2166834",
						Title:  "Batman: The Dark Knight Returns, Part 2",
						Year:   "2013",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BYTEzMmE0ZDYtYWNmYi00ZWM4LWJjOTUtYTE0ZmQyYWM3ZjA0XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt0106364",
						Title:  "Batman: Mask of the Phantasm",
						Year:   "1993",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BYTRiMWM3MGItNjAxZC00M2E3LThhODgtM2QwOGNmZGU4OWZhXkEyXkFqcGdeQXVyNjExODE1MDc@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt1672723",
						Title:  "Batman: Year One",
						Year:   "2011",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BNTJjMmVkZjctNjNjMS00ZmI2LTlmYWEtOWNiYmQxYjY0YWVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt3139086",
						Title:  "Batman: Assault on Arkham",
						Year:   "2014",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BZDU1ZGRiY2YtYmZjMi00ZDQwLWJjMWMtNzUwNDMwYjQ4ZTVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt0060153",
						Title:  "Batman: The Movie",
						Year:   "1966",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BMmM1OGIzM2UtNThhZS00ZGNlLWI4NzEtZjlhOTNhNmYxZGQ0XkEyXkFqcGdeQXVyNTkxMzEwMzU@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt1117563",
						Title:  "Batman: Gotham Knight",
						Year:   "2008",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BM2I0YTFjOTUtMWYzNC00ZTgyLTk2NWEtMmE3N2VlYjEwN2JlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt1568322",
						Title:  "Batman: Arkham City",
						Year:   "2011",
						Type:   "game",
						Poster: "https://m.media-amazon.com/images/M/MV5BZDE2ZDFhMDAtMDAzZC00ZmY3LThlMTItMGFjMzRlYzExOGE1XkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt0147746",
						Title:  "Batman Beyond",
						Year:   "1999–2001",
						Type:   "series",
						Poster: "https://m.media-amazon.com/images/M/MV5BYTBiZjFlZDQtZjc1MS00YzllLWE5ZTQtMmM5OTkyNjZjMWI3XkEyXkFqcGdeQXVyMTA1OTEwNjE@._V1_SX300.jpg",
					},
					{
						ImdbID: "tt3139072",
						Title:  "Son of Batman",
						Year:   "2014",
						Type:   "movie",
						Poster: "https://m.media-amazon.com/images/M/MV5BYjdkZWFhNzctYmNhNy00NGM5LTg0Y2YtZWM4NmU2MWQ3ODVkXkEyXkFqcGdeQXVyNTA0OTU0OTQ@._V1_SX300.jpg",
					},
				},
				TotalResult:   "445",
				ExecutionTime: 0,
			},
		},
	}

	for _, test := range tests {

		_, _ = SearchMovie(test.input.url, &search)

		assert.Equal(t, search, test.result)
	}

}

func TestSearchDetailMovie(t *testing.T) {
	var detail dto.MovieDetail
	// var movie []dto.Movie{}
	tests := []struct {
		input
		result interface{}
	}{
		{
			input: input{url: "http://www.omdbapi.com/?apikey=faf7e5bb&i=tt4853102", resp: detail},
			result: dto.MovieDetail{
				ImdbID:    "tt4853102",
				Title:     "Batman: The Killing Joke",
				Year:      "2016",
				Type:      "movie",
				Poster:    "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
				Rated:     "R",
				Released:  "25 Jul 2016",
				Runtime:   "76 min",
				Genre:     "Animation, Action, Crime",
				Director:  "Sam Liu",
				Writer:    "Brian Azzarello, Brian Bolland, Bob Kane",
				Actors:    "Kevin Conroy, Mark Hamill, Tara Strong",
				Plot:      "As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.",
				Country:   "United States",
				Language:  "English",
				Awards:    "1 win & 2 nominations",
				ImdbVotes: "51,677",
			},
		},
	}

	for _, test := range tests {

		_, _ = SearchMovie(test.input.url, &detail)

		assert.Equal(t, detail, test.result)
	}

}
