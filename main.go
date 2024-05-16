package main

//Use estemain para hacer pruebas mas rapido, lo dejo para que despues se use en el test

import (
	TDADiccionario "tdas/diccionario"
)

func comparar_entero(n1, n2 int) int {
	if n1 > n2 {
		return 1
	}
	if n1 < n2 {
		return -1
	}
	return 0
}
func main() {
	diccionario := TDADiccionario.CrearABB[int, int](comparar_entero)
	desde := 200
	hasta := 300
	iterador := diccionario.IteradorRango(&desde, &hasta)
	iterador.Siguiente()
}
