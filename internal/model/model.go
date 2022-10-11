package model

import "errors"

var (
	ErrPersonNotBeNil     = errors.New("person not be nil")
	ErrIDPersonDoesExists = errors.New("id person does exists")
)
