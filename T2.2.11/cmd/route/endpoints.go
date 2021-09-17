package route

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/listnt/tasks2/T2.2.11/internal/models"
)

// MiddleWareLogging ...
func (s *server) MiddleWareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.lg.Println(r.Method, r.URL.Path, r.URL.Query())
		// Logs
		next.ServeHTTP(w, r)
	})
}

func (s *server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &event)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = s.validator.Validate(event); err != nil {
		s.lg.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.service.CreateEvent(event)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

}

func (s *server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event := models.UpdateEventRequest{}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &event)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.service.UpdateEvent(event.UserId, event.Date, event.Event, event.NewValue)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
func (s *server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	event := models.DeleteEventRequest{}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &event)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.service.DeleteEvent(event.UserId, event.Date, event.Event)
	if err != nil {
		s.lg.Error(err)
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}

func (s *server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	var err error
	event := models.EventsForDate{}
	args := r.URL.Query()
	event.UserId, err = strconv.Atoi(args.Get("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	event.Date = args.Get("date")
	if event.UserId < 1 {
		_, err = w.Write([]byte("{\"error\":" + "wrong UserId" + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var resp models.EventsForDateResponse
	resp.Event, err = s.service.EventForDay(event.UserId, event.Date)
	if err != nil {
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	resp.Date = event.Date
	resp.UserId = event.UserId
	if len(resp.Event) == 0 {
		_, err = w.Write([]byte("{\"error\":" + "NotFound" + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		s.lg.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(respBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	var err error
	event := models.EventsForDate{}
	args := r.URL.Query()
	event.UserId, err = strconv.Atoi(args.Get("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	event.Date = args.Get("date")
	if event.UserId < 1 {
		_, err = w.Write([]byte("{\"error\":" + "wrong UserId" + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var resp models.EventsForDateResponse
	resp.Event, err = s.service.EventsForWeek(event.UserId, event.Date)
	if err != nil {
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	resp.Date = event.Date
	resp.UserId = event.UserId
	if len(resp.Event) == 0 {
		_, err = w.Write([]byte("{\"error\":" + "NotFound" + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(resp)
	if err != nil {
		s.lg.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	var err error
	event := models.EventsForDate{}
	args := r.URL.Query()
	event.UserId, err = strconv.Atoi(args.Get("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	event.Date = args.Get("date")
	if event.UserId < 1 {
		_, err = w.Write([]byte("{\"error\":" + "wrong UserId" + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var resp models.EventsForDateResponse
	resp.Event, err = s.service.EventForMonth(event.UserId, event.Date)
	if err != nil {
		_, err = w.Write([]byte("{\"error\":" + err.Error() + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	resp.Date = event.Date
	resp.UserId = event.UserId
	if len(resp.Event) == 0 {
		_, err = w.Write([]byte("{\"error\":" + "NotFound" + "}"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		s.lg.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(respBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
