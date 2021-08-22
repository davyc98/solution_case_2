package interactor

import (
	"context"

	"github.com/solution_case_2/internal/application/movie-service/dto"
	"github.com/solution_case_2/internal/domain/model"
)

type IMovieService interface {
	Search(context.Context, string, int, model.Log) (dto.SearchResponse, error)
	Detail(context.Context, string, model.Log) (dto.MovieDetail, error)
}
