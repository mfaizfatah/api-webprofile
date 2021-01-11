package usecases

import (
	"context"
	"net/http"

	"github.com/mfaizfatah/api-webprofile/app/models"
)

func (r *uc) InsertMessage(ctx context.Context, req *models.Message) (context.Context, string, int, error) {
	var (
		msg  string
		code int
		err  error

		i = new(models.Message)
	)

	i = req
	go r.query.Insert(TableMessage, i)

	msg = "Success insert message"
	code = http.StatusCreated
	return ctx, msg, code, err
}
