package msgpack

import (
	"github.com/pkg/errors"
	"github.com/tonoy30/hexgo/pkg/models"
	"github.com/vmihailenco/msgpack"
)

type Redirect struct{}

func (r *Redirect) Decode(data []byte) (*models.Short, error) {
	short := new(models.Short)
	if err := msgpack.Unmarshal(data, short); err != nil {
		return nil, errors.Wrap(err, "msgpack.serializer.decode")
	}
	return short, nil
}
func (r *Redirect) Encode(data *models.Short) ([]byte, error) {
	msg, err := msgpack.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "msgpack.serializer.encode")
	}
	return msg, nil
}
