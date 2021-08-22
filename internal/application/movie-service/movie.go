package pusdafil

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/solution_case_2/internal/application/movie-service/dto"
	"github.com/solution_case_2/internal/domain/model"
)

// Search OMDB
func (svc *MovieService) Search(ctx context.Context, movieName string, page int, model model.Log) (res dto.SearchResponse, err error) {
	url := svc.ctx.Value("URL_MOVIE").(string)
	key := svc.ctx.Value("OMDB_KEY").(string)
	url = fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", url, key, movieName, page)
	err, exec := searchMovie(url, &res)
	res.ExecutionTime = exec

	marshalRes, _ := json.Marshal(res)
	model.Response = marshalRes
	model.Url = url

	_, _ = svc.movieRepo.movie.Search(ctx, movieName, page, model)
	return res, err
}

// Get Detail
func (svc *MovieService) Detail(ctx context.Context, movieId string, model model.Log) (response dto.MovieDetail, err error) {
	url := svc.ctx.Value("URL_MOVIE").(string)
	key := svc.ctx.Value("OMDB_KEY").(string)

	url = fmt.Sprintf("%s/?apikey=%s&i=%s", url, key, movieId)
	_, exec := searchMovie(url, &response)
	response.ExecutionTime = exec

	marshalRes, _ := json.Marshal(response)
	model.Response = marshalRes
	model.Url = url

	_, _ = svc.movieRepo.movie.Detail(ctx, movieId, model)
	return response, nil
}
