package json

import (
	"io/ioutil"

	"github.com/goccy/go-json"

	"github.com/rs/zerolog/log"
)

type Json[T interface{}] struct {
	path  string
	value T
}

func (j *Json[T]) Read() (result T, err error) {
	data, e := ioutil.ReadFile(j.path)
	if e != nil {
		log.Error().Err(e).Msg("json read fail")
		err = e
		return
	}
	if e := json.Unmarshal(data, j.value); e != nil {
		log.Error().Err(e).Msg("json read fail")
		err = e
		return
	}
	result = j.value
	return
}

func (j *Json[T]) Write(data T) error {
	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Error().Err(err).Msg("json read fail")
		return err
	}
	err = ioutil.WriteFile(j.path, res, 0777)
	if err != nil {
		log.Error().Err(err).Msg("json read fail")
		return err
	}
	return nil
}

func Init[T interface{}](path string, value T) Json[T] {
	return Json[T]{
		path:  path,
		value: value,
	}
}
