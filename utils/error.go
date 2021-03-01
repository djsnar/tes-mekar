package utils

import "errors"

func ValidationHandler(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("Data input not allowed empty!")
		case 0:
			return errors.New("Data input not allowed = 0")
		case nil:
			return errors.New("Data input cannot be nil!")
		}
	}
	return nil
}
