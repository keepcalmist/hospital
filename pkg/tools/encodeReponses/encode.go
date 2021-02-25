package encodeReponses

import (
	"encoding/json"
)

func Json(data interface{}) ([]byte, error) {
	dataR, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataR, err
}
