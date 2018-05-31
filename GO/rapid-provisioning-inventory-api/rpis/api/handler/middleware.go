package handler

import (
	"net/http"
	"time"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

// RequestLogger logs all Request
func RequestLogger(f httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		t := time.Now()
		f(w, r, p)
		log.WithFields(log.Fields{
			"method":   r.Method,
			"resource": r.URL.Path,
			"took":     fmt.Sprintf("%d%s", time.Since(t).Nanoseconds()/1000000, "ms"),
		}).Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	})
}
