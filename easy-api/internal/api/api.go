package api

import (
	"fmt"
	"net/http"

	"github.com/ikiwq/ackme/easy-api/internal/domain"
	"github.com/ikiwq/ackme/easy-api/internal/repository"
	connection "github.com/ikiwq/ackme/mysql"
	"github.com/jmoiron/sqlx"
)

type api struct {
	apiAddress string
	apiPort    string

	httpClient *http.Client
	dbConnection *sqlx.DB

	easyUserRepository domain.EasyUserRepository
}

func NewApi(apiAddress string, apiPort string, connectionString string) *api {
	client := &http.Client{}

	conn := connection.InitDB(connectionString)
	easyUserRepository := repository.NewMySqlEasyUserRepository(conn)

	return &api{
		apiAddress: apiAddress,
		apiPort:    apiPort,

		httpClient:   client,
		dbConnection: conn,

		easyUserRepository: easyUserRepository,
	}
}

func (a *api) buildRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /api/v1/auth/login", a.login)
	
	return r
}


func (a *api) Start() {
	r := a.buildRoutes()

	listenAddress := fmt.Sprintf("%s:%s", a.apiAddress, a.apiPort)
	http.ListenAndServe(listenAddress, corsMiddleware(r))
}

func (a *api) Exit() {
	a.httpClient.CloseIdleConnections()
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}