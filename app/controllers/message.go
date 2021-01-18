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

	s.Channel = r.Header.Get("channel")

	ctx, msg, st, err := u.uc.InsertMessage(ctx, &s)
	if err != nil {
		ctx = logger.Logf(ctx, "insert error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

func (u *ctrl) HandlerGetAllMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	Channel := r.Header.Get("channel")

	ctx, res, msg, st, err := u.uc.GetAllMessage(ctx, Channel)
	if err != nil {
		ctx = logger.Logf(ctx, "get error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}
