package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sijanstha/electronic-voting-system/internal/adapters/handler"
	"github.com/sijanstha/electronic-voting-system/internal/adapters/repository"
	"github.com/sijanstha/electronic-voting-system/internal/core/domain"
	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"github.com/sijanstha/electronic-voting-system/internal/core/ports"
	"github.com/sijanstha/electronic-voting-system/internal/core/services"
	"github.com/sijanstha/electronic-voting-system/internal/core/utils"
)

var pollService ports.PollService
var userService ports.UserService
var authService ports.AuthenticationService
var jwtService ports.TokenService

type ApiServer struct {
	listenAddr string
	db         *sql.DB
}

func NewApiServer(listenAddr string, db *sql.DB) *ApiServer {
	return &ApiServer{listenAddr: listenAddr, db: db}
}

func (s *ApiServer) Run() {
	pollRepo := repository.NewPollMysqlRepository(s.db)
	userRepo := repository.NewUserRepository(s.db)
	pollOrganizerRepo := repository.NewPollOrganizerRepository(s.db)

	jwtService = services.NewJwtService()
	pollService = services.NewPollService(pollRepo, pollOrganizerRepo)
	userService = services.NewUserService(userRepo)
	authService = services.NewAuthenticationService(userRepo, jwtService)

	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()
	s.registerPublicRoutes(router)
	s.registerPollRoutes(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization", "Referer"},
	})

	handler := c.Handler(router)

	log.Println("electronic-voting-system API running at port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, handler)
}

func (s *ApiServer) registerPollRoutes(router *mux.Router) {
	pollHandler := handler.NewPollHandler(pollService)
	router.HandleFunc("/poll", authMiddleware(makeHTTPHandleFunc(pollHandler.HandlePoll))).Methods("POST", "PUT", "GET")
	router.HandleFunc("/poll/{id}", authMiddleware(makeHTTPHandleFunc(pollHandler.HandleGetPollById))).Methods("GET")
}

func (s *ApiServer) registerPublicRoutes(router *mux.Router) {
	authHandler := handler.NewAuthenticationHandler(userService, authService)
	router.HandleFunc("/login", makeHTTPHandleFunc(authHandler.HandleUserAuthentication)).Methods("POST")
	router.HandleFunc("/register", makeHTTPHandleFunc(authHandler.HandleRegisterUser)).Methods("POST")
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			var badRequestErr *commonError.ErrBadRequest
			var uniqueConstraintErr *commonError.ErrUniqueConstraintViolation
			var unauthorizedErr *commonError.ErrUnauthorized
			var notFoundErr *commonError.ErrNotFound
			switch {
			case errors.As(err, &badRequestErr):
				utils.WriteJSON(w, http.StatusBadRequest, domain.NewApiError(err.Error()))
			case errors.As(err, &uniqueConstraintErr):
				utils.WriteJSON(w, http.StatusConflict, domain.NewApiError(err.Error()))
			case errors.As(err, &unauthorizedErr):
				utils.WriteJSON(w, http.StatusUnauthorized, domain.NewApiError(err.Error()))
			case errors.As(err, &notFoundErr):
				utils.WriteJSON(w, http.StatusNotFound, domain.NewApiError(err.Error()))
			default:
				utils.WriteJSON(w, http.StatusInternalServerError, domain.NewApiError(err.Error()))
			}
		}
	}
}

func authMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("inside authentication middleware")
		token := r.Header.Get("Authorization")
		log.Println("token:", token)
		claims, err := jwtService.Validate(token)

		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, domain.NewApiError(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), "principal", claims)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}
