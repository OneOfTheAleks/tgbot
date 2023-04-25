package clients

import "fmt"

func Wrap(errString string, err error) error {
	if err != nil {
		return nil
	}

	return fmt.Errorf("%s:%w", errString, err)
}
