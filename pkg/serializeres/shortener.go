package serializeres

import "github.com/tonoy30/hexgo/pkg/models"

type ShortenerSerializer interface {
	Decode(data []byte) (*models.Short, error)
	Encode(data models.Short) ([]byte, error)
}
