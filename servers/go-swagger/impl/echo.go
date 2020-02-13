package impl

import (
	"net/http"

	"github.com/ericatcorsha/rpc-compare/servers/go-swagger/models"
	"github.com/ericatcorsha/rpc-compare/servers/go-swagger/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func errorReply(code int, msg string) middleware.Responder {
	codeText := http.StatusText(code)
	return operations.NewEchoDefault(code).WithPayload(
		&models.ErrorReply{
			Code:    &codeText,
			Message: &msg,
		},
	)
}

func NewEchoHandler() operations.EchoHandler {
	// Ew, no reply type in the handler function signature? How do we know we've replied with a valid response?
	return operations.EchoHandlerFunc(func(params operations.EchoParams) middleware.Responder {
		// nit: EchoRequest is a pointer despite saying it is required in the swagger... That requirement is something that the type system isn't enfocing
		if params.EchoRequest == nil {
			return errorReply(http.StatusInternalServerError, "EchoRequest is unexpectedly nil")
		}

		if *params.EchoRequest.Message == "user-error" {
			return errorReply(http.StatusBadRequest, "The User Made an Error")
		}

		return operations.NewEchoOK().WithPayload(&models.EchoReply{
			Message: params.EchoRequest.Message,
		})
	})
}
