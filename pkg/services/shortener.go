package services

import "github.com/tonoy30/hexgo/pkg/models"

type ShortenerService interface {
	Find(code string) (*models.Short, error)
	Store(short *models.Short) error
}
