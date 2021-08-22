package mysql

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/solution_case_2/internal/application/movie-service/dto"
	"github.com/solution_case_2/internal/domain/model"
	"gorm.io/gorm"
)

// MovieRepo ...
type MovieRepo struct {
	db *gorm.DB
}

// NewMovieRepo ...
func NewMovieRepo(db *gorm.DB) *MovieRepo {
	return &MovieRepo{
		db: db,
	}
}

// Find
func (repo *MovieRepo) Search(ctx context.Context, movieName string, page int, model model.Log) (res dto.SearchResponse, err error) {
	model.CreatedAt = time.Now()
	result := repo.db.Save(&model)
	err = result.Error
	if err != nil {
		log.Println(err.Error())
	}
	return res, err
}

func (repo *MovieRepo) Detail(ctx context.Context, movieId string, model model.Log) (res dto.MovieDetail, err error) {
	model.CreatedAt = time.Now()
	result := repo.db.Save(&model)
	err = result.Error
	if err != nil {
		log.Println(err.Error())
	}
	return res, err
}
