// helper functions for http reply
package httpreply

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type replier struct {
	opts Options
}

func NewHTTPReplier(opts Options) *replier {
	return &replier{opts: opts}
}

type Options struct {
	ResponseWriter http.ResponseWriter
	ContentType    string
	Accept         []string
}

func (re *replier) Reply(v interface{}, code int, w http.ResponseWriter, logError bool) {
	re.setHeader()

	b, err := json.Marshal(v)
	if err != nil {
		re.Error(err.Error(), http.StatusInternalServerError, w, logError)
		return
	}

	re.opts.ResponseWriter.WriteHeader(code)
	re.opts.ResponseWriter.Write(b)
}

func (re *replier) Error(err string, code int, w http.ResponseWriter, logError bool) {
	re.setHeader()
	re.opts.ResponseWriter.WriteHeader(code)

	if logError {
		log.Println(err)
		return
	}

	e := &httpError{Error: err, Code: code}
	b, _ := json.Marshal(e)
	re.opts.ResponseWriter.Write(b)
}

func (re *replier) setHeader() {
	re.opts.ResponseWriter.Header().Set("Content-Type", re.opts.ContentType)
	re.opts.ResponseWriter.Header().Set("Access", strings.Join(re.opts.Accept, ","))
}

type httpError struct {
	Error string
	Code  int
}
