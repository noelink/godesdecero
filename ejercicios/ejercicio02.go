package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicador() string {
	var err error
	var numero int
	var texto string
	fmt.Println("Ingresa un numero a multiplicar")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto!! \n" + err.Error())
		}
		fmt.Println("Generando tabla de multiplicar para: ", numero)
		for i := 1; i <= 10; i++ {
			texto += fmt.Sprintf("%d * %d = %d\n", numero, i, numero*1)
		}
	}
	return texto
}
