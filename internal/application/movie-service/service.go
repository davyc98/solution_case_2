package pusdafil

import "context"

type MovieRepo struct {
	movie iMovie
}

// Service ...
type MovieService struct {
	ctx       context.Context
	movieRepo MovieRepo
}

// NewService instantiation
func NewService(ctx context.Context, pusdafil iMovie) *MovieService {

	return &MovieService{
		ctx: ctx,
		movieRepo: MovieRepo{
			movie: pusdafil,
		},
	}
}
