package geo

import "errors"

func CubeVolume(n int) (int, error) {

	if n == 0 {
		return 0, errors.New("zero cant be")
	} else {
		return n * n * n, nil
	}
}
