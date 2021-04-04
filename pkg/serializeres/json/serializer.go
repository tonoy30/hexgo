package json

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/tonoy30/hexgo/pkg/models"
)

type Redirect struct{}

func (r *Redirect) Decode(data []byte) (*models.Short, error) {
	short := new(models.Short)
	if err := json.Unmarshal(data, short); err != nil {
		return nil, errors.Wrap(err, "json.serializer.decode")
	}
	return short, nil
}
func (r *Redirect) Encode(data models.Short) ([]byte, error) {
	msg, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "json.serializer.encode")
	}
	return msg, nil
}
