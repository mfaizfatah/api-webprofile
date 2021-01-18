package repository

import (
	"log"

	"github.com/mfaizfatah/api-webprofile/app/models"
)

func (r *repo) FindAllMessageWithChannel(table string, where interface{}, whereValue ...interface{}) ([]models.Message, error) {
	var data []models.Message
	err := r.db.Table(table).Where(where, whereValue...).Scan(&data)
	log.Printf("msg: %v", data)
	if err != nil {
		return data, nil
	}

	return nil, nil
}
