package business

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"github.com/tonoy30/hexgo/pkg/models"
	"github.com/tonoy30/hexgo/pkg/repo"
	"github.com/tonoy30/hexgo/pkg/services"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrRedirectNotFound = errors.New("redirect not found")
	ErrRedirectInvalid  = errors.New("redirect invalid")
)

type shortenerService struct {
	repo repo.ShortenerRepository
}

func NewShortenerService(repo repo.ShortenerRepository) services.ShortenerService {
	return &shortenerService{
		repo,
	}
}
func (s *shortenerService) Find(code string) (*models.Short, error) {
	return s.repo.Find(code)
}
func (s *shortenerService) Store(short *models.Short) error {
	if err := validate.Validate(short); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "business.ShortenerService.Store")
	}
	short.Code = shortid.MustGenerate()
	short.CreatedAt = time.Now().UTC().Unix()
	return s.repo.Store(short)
}
