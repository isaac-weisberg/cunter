package cunter

import (
	"encoding/json"
	"io/ioutil"
)

func export(path string, object interface{}) error {
	var error error
	var data []byte
	data, error = json.Marshal(object)
	if error != nil {
		return error
	}
	error = ioutil.WriteFile(path, data, 644)
	return error
}
