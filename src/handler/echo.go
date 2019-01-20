package handler

import (
	"net/http"

	pb "github.com/micnncim/appengine-protobuf-api/src/proto"
	"github.com/micnncim/appengine-protobuf-api/src/response"
)

// Echo handler for health check
func Echo(w http.ResponseWriter, r *http.Request) {
	req := pb.EchoRequest{}
	if err := Bind(r, &req); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.Succeed(w, r, &pb.EchoResponse{
		Body: req.Body,
	})
}
