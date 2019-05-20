package impl

import "errors"

type SayHello struct{}

func (s SayHello) SayHello(name string) (string, error) {

	if name == "wrong" {
		return "", errors.New("sorry wrong name")
	}

	return "Hello " + name, nil

}
