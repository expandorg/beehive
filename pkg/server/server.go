package server

import (
	"net/http"

	"github.com/gemsorg/beehive/pkg/api/jobsolutionsmaker"
	"github.com/gemsorg/beehive/pkg/api/solutionmaker"

	"github.com/gemsorg/beehive/pkg/authentication"

	"github.com/jmoiron/sqlx"

	"github.com/gemsorg/beehive/pkg/api/healthchecker"
	"github.com/gemsorg/beehive/pkg/service"
	"github.com/gorilla/mux"
)

func New(
	db *sqlx.DB,
	s service.BeehiveService,
) http.Handler {
	r := mux.NewRouter()

	r.Handle("/_health", healthchecker.MakeHandler(s)).Methods("GET")
	r.Handle("/honeypots/{task_id}/solutions", solutionmaker.MakeHandler(s)).Methods("POST")
	r.Handle("/jobs/{job_id}/solutions", jobsolutionsmaker.MakeHandler(s)).Methods("POST")
	r.Use(authentication.AuthMiddleware)
	return withHandlers(r)
}
