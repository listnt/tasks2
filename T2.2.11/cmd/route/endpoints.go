package route

import (
	"encoding/json"
	"github.com/listnt/tasks2/T2.2.11/internal/models"
	"io"
	"net/http"
	"strconv"
)

// MiddleWareLogging ...
func (s *server) MiddleWareLogging (next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.lg.Println(r.Method, r.URL.Path, r.URL.Query())
		// Аутентификация прошла успешно, направляем запрос следующему обработчику
		next.ServeHTTP(w, r)
	})
}

func (s *server) CreateEvent (w http.ResponseWriter, r *http.Request) {
	e:=models.Event{}
	b,err:=io.ReadAll( r.Body)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err=json.Unmarshal(b ,&e)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err:=s.validator.Validate(e);err!=nil{
		s.lg.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err=s.service.CreateEvent(e)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	}

func (s *server) UpdateEvent (w http.ResponseWriter, r *http.Request) {
	e:=models.UpdateEventRequest{}
	b,err:=io.ReadAll( r.Body)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err=json.Unmarshal(b ,&e)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err=s.service.UpdateEvent(e.UserId,e.Date,e.Event,e.NewValue)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
func (s *server) DeleteEvent (w http.ResponseWriter, r *http.Request) {
	e:=models.DeleteEventRequest{}
	b,err:=io.ReadAll( r.Body)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err=json.Unmarshal(b ,&e)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err= s.service.DeleteEvent(e.UserId,e.Date,e.Event)
	if err!=nil{
		s.lg.Error(err)
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	}

func (s *server) EventsForDay (w http.ResponseWriter, r *http.Request) {
	e:=models.EventsForDate{}
	args:=r.URL.Query()
	e.UserId,_=strconv.Atoi( args.Get("user_id"))
	e.Date=args.Get("date")
	if e.UserId<1{
		_,err:=w.Write([]byte("{\"error\":"+"wrong UserId"+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var err error
	var resp models.EventsForDateResponse
	resp.Event, err = s.service.EventForDay(e.UserId,e.Date)
	if err!=nil{
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	resp.Date=e.Date
	resp.UserId=e.UserId
	if len(resp.Event)==0{
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b,err:=json.Marshal(resp)
	if err!=nil{
		s.lg.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_,err=w.Write(b)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) EventsForWeek (w http.ResponseWriter, r *http.Request) {
	e:=models.EventsForDate{}
	args:=r.URL.Query()
	e.UserId,_=strconv.Atoi( args.Get("user_id"))
	e.Date=args.Get("date")
	if e.UserId<1{
		_,err:=w.Write([]byte("{\"error\":"+"wrong UserId"+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var err error
	var resp models.EventsForDateResponse
	resp.Event, err = s.service.EventsForWeek(e.UserId,e.Date)
	if err!=nil{
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	resp.Date=e.Date
	resp.UserId=e.UserId
	if len(resp.Event)==0{
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b,err:=json.Marshal(resp)
	if err!=nil{
		s.lg.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_,err=w.Write(b)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) EventsForMonth (w http.ResponseWriter, r *http.Request) {
	e:=models.EventsForDate{}
	args:=r.URL.Query()
	e.UserId,_=strconv.Atoi( args.Get("user_id"))
	e.Date=args.Get("date")
	if e.UserId<1{
		_,err:=w.Write([]byte("{\"error\":"+"wrong UserId"+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var err error
	var resp models.EventsForDateResponse
	resp.Event, err = s.service.EventForMonth(e.UserId,e.Date)
	if err!=nil{
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	resp.Date=e.Date
	resp.UserId=e.UserId
	if len(resp.Event)==0{
		_,err:=w.Write([]byte("{\"error\":"+err.Error()+"}"))
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b,err:=json.Marshal(resp)
	if err!=nil{
		s.lg.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_,err=w.Write(b)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}