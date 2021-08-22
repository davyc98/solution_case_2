package pusdafil

import (
	"context"

	"github.com/solution_case_2/internal/application/movie-service/dto"
	"github.com/solution_case_2/internal/domain/model"
)

// Adapter to lender repository
type iMovie interface {
	Search(ctx context.Context, movieName string, page int, model model.Log) (res dto.SearchResponse, err error)
	Detail(ctx context.Context, movieId string, model model.Log) (res dto.MovieDetail, err error)
}
