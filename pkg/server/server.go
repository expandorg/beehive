package server

import (
	"net/http"

	"github.com/expandorg/beehive/pkg/api/jobsolutionsdestroyer"
	"github.com/expandorg/beehive/pkg/api/jobsolutionsmaker"
	"github.com/expandorg/beehive/pkg/api/solutionmaker"

	"github.com/expandorg/beehive/pkg/authentication"

	"github.com/jmoiron/sqlx"

	"github.com/expandorg/beehive/pkg/api/healthchecker"
	"github.com/expandorg/beehive/pkg/service"
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
	r.Handle("/jobs/{job_id}/solutions", jobsolutionsdestroyer.MakeHandler(s)).Methods("DELETE")
	r.Use(authentication.AuthMiddleware)
	return withHandlers(r)
}
