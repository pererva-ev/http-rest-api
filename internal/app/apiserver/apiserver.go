package apiserver

import (
	"io"
	"net/http"
	"sync/atomic"
	// "fmt"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)
var requests int64 = 0

//APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) configureRouter()  {
	s.router.HandleFunc("/hello", s.hahdleHello())
}

func (s *APIserver) hahdleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		incRequests()
		str := strconv.FormatInt(getRequests(), 10)
		// fmt.Println(str)
	

		io.WriteString(w, str)
		
		
	}
}

// increments the number of requests and returns the new value
func incRequests() int64 {
    return atomic.AddInt64(&requests, 1)
}

// returns the current value
func getRequests() int64 {
    return atomic.LoadInt64(&requests)
}
