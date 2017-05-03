package timeserver

import (
	"net/http"
	"time"

	"io"
)

type timeProvider interface {
	getCurrentTime() string
}
type timerProv struct{}

func (tp timerProv) getCurrentTime() string {
	return time.Now().String()
}

type timer struct {
	tp timeProvider
}

func (t timer) RealTime(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, t.tp.getCurrentTime())
}

func main() {
	TimeHandler := timer{timerProv{}}

	http.HandleFunc("/", TimeHandler.RealTime)
	http.ListenAndServe(":8080", nil)
}
