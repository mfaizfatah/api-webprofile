package controllers

import (
	"net/http"

	"github.com/mfaizfatah/api-webprofile/app/usecases"
)

// ctrl struct with value interface Usecases
type ctrl struct {
	uc usecases.Usecases
}

// Controllers represent the Controllers contract
type Controllers interface {
	HandlerInsertMessage(w http.ResponseWriter, r *http.Request)
}

/*NewCtrl will create an object that represent the Controllers interface (Controllers)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Controllers
 *
 * @return
 * uc struct with value interface Usecases
 */
func NewCtrl(u usecases.Usecases) Controllers {
	return &ctrl{uc: u}
}
