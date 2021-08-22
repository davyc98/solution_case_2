package main

import (
	"context"
	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	pusdafil "github.com/solution_case_2/internal/application/movie-service"
	"github.com/solution_case_2/internal/interactor/rest"
	mysql "github.com/solution_case_2/internal/repository/implementor/mysql"
)

type envConfig struct {
	DBHost      string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBName      string `envconfig:"DB_NAME" default:"test"`
	DBUsername  string `envconfig:"DB_USER" default:"root"`
	DBPassword  string `envconfig:"DB_PWD" default:""`
	DBPort      int    `envconfig:"DB_PORT" default:"3306"`
	ServicePort string `envconfig:"SVC_PORT" default:"8585"`
	DebugMode   string `envconfig:"DEBUG_MODE" default:"true"`
	MovieUrl    string `envconfig:"URL_MOVIE" default:"http://www.omdbapi.com"`
	OMDBKey     string `envconfig:"OMDB_KEY" default:"faf7e5bb"`
}

var env envConfig

func init() {
	env = envConfig{}
	envconfig.MustProcess("", &env)
}

func main() {
	// Connect to database
	db, err := mysql.Connect(env.DBHost, env.DBName, env.DBPort, env.DBUsername, env.DBPassword)
	if err == nil {
		if err != nil {

			log.Println(err.Error())
		}
		ctx := context.WithValue(context.Background(), "URL_MOVIE", env.MovieUrl)
		ctx = context.WithValue(ctx, "OMDB_KEY", env.OMDBKey)

		application := pusdafil.NewService(ctx, mysql.NewMovieRepo(db))

		// REST API Routes
		router := httprouter.New()
		restapi := rest.NewRestAPI(ctx, application)
		restapi.WithRoutes(router)

		log.Fatal(http.ListenAndServe(":8089", router))
	} else {
		log.Println(err.Error())
	}
}
