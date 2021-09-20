package route

import (
	"net/http"

	"github.com/gorilla/mux"
	Config "github.com/listnt/tasks2/T2.2.11/config"
	Service "github.com/listnt/tasks2/T2.2.11/internal"
	"github.com/listnt/tasks2/T2.2.11/utils"
	"github.com/sirupsen/logrus"
)

//Server ...
type Server interface {
	Launch() chan error
	Close()
}

type server struct {
	service    Service.Service
	Handler    *mux.Router
	httpserver *http.Server
	validator  utils.Validator
	lg         *logrus.Logger
}

//NewServer ...
func NewServer(cfgpath string) Server {
	s := new(server)
	cfg := Config.ParseConfig(cfgpath)
	s.lg = logrus.New()
	s.Handler = mux.NewRouter()
	s.validator = utils.ValidatorNew()
	go func() {
		r := http.NewServeMux()
		r.HandleFunc("/status", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		err := http.ListenAndServe(":8081", r)
		if err != nil {
			logrus.Fatal(err)
		}
	}()
	s.Handler.Handle("/create_event/", s.MiddleWareLogging(http.HandlerFunc(s.CreateEvent))).Methods("POST")
	s.Handler.Handle("/update_event/", s.MiddleWareLogging(http.HandlerFunc(s.UpdateEvent))).Methods("POST")
	s.Handler.Handle("/delete_event/", s.MiddleWareLogging(http.HandlerFunc(s.DeleteEvent))).Methods("POST")
	s.Handler.Handle("/events_for_day/", s.MiddleWareLogging(http.HandlerFunc(s.EventsForDay))).Methods("GET")
	s.Handler.Handle("/events_for_week/", s.MiddleWareLogging(http.HandlerFunc(s.EventsForWeek))).Methods("GET")
	s.Handler.Handle("/events_for_month/", s.MiddleWareLogging(http.HandlerFunc(s.EventsForMonth))).Methods("GET")
	s.httpserver = &http.Server{Addr: ":8080", Handler: s.MiddleWareLogging(s.Handler)}
	s.service = Service.NewService(s.lg, &cfg)
	return s
}

func (s *server) Launch() chan error {
	return s.serve()
}

func (s *server) serve() chan error {
	listenErr := make(chan error, 1)
	go func() {
		listenErr <- s.httpserver.ListenAndServe()
	}()
	return listenErr
}

func (s *server) Close() {
	s.service.Close()
}
