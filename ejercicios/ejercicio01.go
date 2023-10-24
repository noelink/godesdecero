package ejercicios

import "strconv"

func ConvierteEnteros(enteroConvertir string) (int, string) {

	if intConvertido, err := strconv.Atoi(enteroConvertir); err != nil && intConvertido > 100 {
		return intConvertido, "Es mayor a 100"
	} else {
		return intConvertido, "Es menor a 100"
	}

}
