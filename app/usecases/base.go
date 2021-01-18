package usecases

import (
	"context"

	"github.com/mfaizfatah/api-webprofile/app/models"
	"github.com/mfaizfatah/api-webprofile/app/repository"
)

// all variable const
const (
	// table
	TableMessage = "message"
)

// uc struct with value interface Repository
type uc struct {
	query repository.Repo
}

// Usecases represent the Usecases contract
type Usecases interface {
	InsertMessage(ctx context.Context, req *models.Message) (context.Context, string, int, error)
	GetAllMessage(ctx context.Context, channel string) (context.Context, interface{}, string, int, error)
}

/*NewUC will create an object that represent the Usecases interface (Usecases)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Usecases
 *
 * @return
 * uc struct with value interface Repository
 */
func NewUC(r repository.Repo) Usecases {
	return &uc{query: r}
}
