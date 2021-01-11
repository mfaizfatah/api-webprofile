package usecases

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/mfaizfatah/api-webprofile/app/models"
)

func (r *uc) InsertMessage(ctx context.Context, req *models.Message) (context.Context, string, int, error) {
	var (
		msg  string
		code int
		err  error

		i = new(models.Message)
	)

	if req.Email == "" || !isEmailValid(req.Email) {
		msg = "Invalid Request"
		code = http.StatusBadRequest
		return ctx, msg, code, errors.New("badRequest")
	}

	i = req
	go r.query.Insert(TableMessage, i)

	msg = "Success insert message"
	code = http.StatusCreated
	return ctx, msg, code, err
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
