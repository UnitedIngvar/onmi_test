package sendRequest

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/UnitedIngvar/onmi_test/internal/lib/slogext"
	"github.com/go-playground/validator"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Clienter interface {
	SendRequest(context.Context, uint64) error
}

// SendRequest sends request to the external server
// SendRequest             godoc
// @Summary      Send request to the external server
// @Description  Takes
// @Accept				json
// @Produce      json
// @Param        request body sendRequest.Request false "request"
// @Success      200  {object} sendRequest.Response "response"
// @Router       /send-request [post]
func NewHandler(log *slog.Logger, client Clienter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const loc = "onmi_test.internal.trapnsport.handler.NewHandler"

		log = log.With(
			slog.String("loc", loc),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")

			render.JSON(w, r, Error("empty request"))

			return
		}

		if err != nil {
			log.Error("failed to decode request body", err)

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", slogext.Error(err))

			render.JSON(w, r, Error(validateErr.Error()))

			return
		}

		log.Info("request body decoded", slog.Any("req", req))

		if err := client.SendRequest(r.Context(), req.ItemCount); err != nil {
			log.Error("business logic error occured",
				slogext.Error(err))

			render.JSON(w, r, Error(err.Error()))

			return
		}

		render.JSON(w, r, Ok())
	}
}
