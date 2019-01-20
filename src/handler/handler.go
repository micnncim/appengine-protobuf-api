package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

func Bind(r *http.Request, dst proto.Unmarshaler) (err error) {
	contentType := r.Header.Get("Content-Type")
	if idx := strings.IndexByte(contentType, ';'); idx >= 0 {
		contentType = contentType[:idx]
	}
	log.Printf("content-type=%s\n", contentType)

	switch contentType {
	case "application/protobuf", "application/x-protobuf":
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if err = dst.Unmarshal(buf); err != nil {
			return errors.Wrap(err, "proto unmarshal failed")
		}
	case "application/json":
		if err = json.NewDecoder(r.Body).Decode(&dst); err != nil {
			return errors.Wrap(err, "json unmarshal failed")
		}
	default:
		return errors.New("unsupported content-type")
	}

	return
}
