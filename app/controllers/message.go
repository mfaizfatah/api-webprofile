package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mfaizfatah/api-webprofile/app/helpers/logger"
	"github.com/mfaizfatah/api-webprofile/app/models"
	"github.com/mfaizfatah/api-webprofile/app/utils"
)

func (u *ctrl) HandlerInsertMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var s models.Message

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	ctx, msg, st, err := u.uc.InsertMessage(ctx, &s)
	if err != nil {
		ctx = logger.Logf(ctx, "insert error() => %v", err)
		utils.Response(ctx, w, false, http.StatusInternalServerError, "error while insert message")
		return
	}

	utils.Response(ctx, w, true, st, msg)
}
