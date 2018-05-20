package cunter

import (
	"encoding/json"
	"io/ioutil"

	"github.com/isaac-weisberg/cundef"
)

func read(path string) (*cundef.Project, error) {
	var error error
	var data []byte
	data, error = ioutil.ReadFile(path)
	if error != nil {
		return nil, error
	}
	var project = cundef.Project{}
	error = json.Unmarshal(data, project)
	if error != nil {
		return nil, error
	}
	return &project, nil
}

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
