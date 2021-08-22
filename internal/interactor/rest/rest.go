package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/solution_case_2/internal/domain/model"
	"github.com/solution_case_2/internal/interactor"
)

// Error collection
var (
	ErrPayloadEmpty = errors.New("Payload is empty")
)

//adapter adapter
type adapter struct {
	movieService interactor.IMovieService
}

// API API entrypoint
type API struct {
	ctx     context.Context
	adapter adapter
}

// NewRestAPI instance
func NewRestAPI(ctx context.Context, service interactor.IMovieService) *API {
	return &API{
		ctx: ctx,
		adapter: adapter{
			movieService: service,
		},
	}
}

// WithRoutes ...
func (c *API) WithRoutes(router *httprouter.Router) {
	router.GET("/healthcheck", healthcheck)
	router.GET("/search/:movieName", c.Search)
	router.GET("/detail", c.Detail)
}

func healthcheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	// TODO: return CPU & Mem Usage? Running go routines?
}

//Pusdafil Send
func (c *API) Search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		page  int
		err   error
		model model.Log
	)
	queryValues := r.URL.Query()
	resQuery := queryValues.Get("page")
	if resQuery != "" {
		page, err = strconv.Atoi(resQuery)
		if err != nil {
			log.Println(err.Error())
		}
	}
	movieName := p.ByName("movieName")

	res, err := c.adapter.movieService.Search(r.Context(), movieName, page, model)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	marshalRes, _ := json.Marshal(res)
	fmt.Fprint(w, string(marshalRes))

}

//Pusdafil Find
func (c *API) Detail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		err   error
		model model.Log
	)

	queryValues := r.URL.Query()
	movieId := queryValues.Get("movieId")
	res, _ := c.adapter.movieService.Detail(r.Context(), movieId, model)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	marshalRes, _ := json.Marshal(res)
	fmt.Fprint(w, string(marshalRes))

}
