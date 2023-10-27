package mapas

import "fmt"

func MostrarMapas() {
	//Aqui tambien podemos usar make como cuando hicimos el slice
	//                 clave    valor
	paises := make(map[string]string)

	paises["Mexico"] = "CDMX"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	//otra forma de crear un map
	//                clave   valor
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Cruz Azul":    3,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s tiene un puntaje de %d\n", equipo, puntaje)
	}
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Cruz Azul"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t\n", puntaje, existe)
}
