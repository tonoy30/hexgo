package repo

import "github.com/tonoy30/hexgo/pkg/models"

type ShortenerRepository interface {
	Find(code string) (*models.Short, error)
	Store(short *models.Short) error
}
