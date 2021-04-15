package render

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type StringMap struct {
	Data map[string]string
}

func writeStringMap(w http.ResponseWriter, sm StringMap) (err error) {
	var jsonBytes []byte
	writeContentType(w, jsonContentType)
	if jsonBytes, err = json.Marshal(sm.Data); err != nil {
		err = errors.WithStack(err)
		return
	}
	if _, err = w.Write(jsonBytes); err != nil {
		err = errors.WithStack(err)
	}
	return
}

func (sm StringMap)Render(w http.ResponseWriter) error  {
	return writeStringMap(w,sm)
}

// WriteContentType write json ContentType.
func (sm StringMap) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
