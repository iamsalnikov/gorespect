package main

import "errors"

var ErrCanNotSayThanks = errors.New("can not say thanks")

type Thanker interface {
	CanProcess(p string) bool
	SayThankYou(p string) error
}


