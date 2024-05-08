package main

//Use estemain para hacer pruebas mas rapido, lo dejo para que despues se use en el test

import (
	"fmt"
	"strings"
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
	diccionario1 := TDADiccionario.CrearABB[string, int](strings.Compare)
	diccionario1.Guardar("hola", 10)
	diccionario1.Guardar("mundo", 30)
	diccionario1.Guardar("XD", 45)
	diccionario1.Guardar("nashe", 34)
	cont := 0
	prueba := func(string, int) bool {
		cont++
		return true
	}
	diccionario1.Iterar(prueba)
	fmt.Printf("Los elementos son %d\n", cont)
	//Probamos que iterar funca
	diccionario2 := TDADiccionario.CrearABB[int, int](comparar_entero)
	diccionario2.Guardar(7, 6)
	diccionario2.Guardar(10, 6)
	diccionario2.Guardar(4, 4)
	diccionario2.Guardar(5, 5)
	if diccionario2.Borrar(7) == 6 {
		fmt.Println("nashe")
	}
	if diccionario2.Obtener(4) == 4 && diccionario2.Obtener(5) == 5 {
		fmt.Println("recontra nashe")
	}
	//Pasa una prueba el borrar ,estoy muy a la izquierda
	diccionario3 := TDADiccionario.CrearABB[int, int](comparar_entero)
	diccionario3.Guardar(7, 6)
	diccionario3.Guardar(10, 6)
	diccionario3.Guardar(3, 3)
	diccionario3.Guardar(5, 5)
	diccionario3.Guardar(4, 4)
	if diccionario3.Borrar(7) == 6 {
		fmt.Println("nashe 2")
	}
	if diccionario3.Obtener(4) == 4 && diccionario3.Obtener(5) == 5 {
		fmt.Println("recontra nashe 2")
	}
	//Pasa otra el borrar,estoy muy a la izquierda y mi nodo remplazo tiene un hijo

	diccionario4 := TDADiccionario.CrearABB[int, int](comparar_entero)
	diccionario4.Guardar(7, 6)
	diccionario4.Guardar(10, 6)
	diccionario4.Guardar(3, 3)
	diccionario4.Guardar(5, 5)
	diccionario4.Guardar(4, 4)
	if diccionario4.Borrar(3) == 3 {
		fmt.Println("nashe 3")
	}
	if diccionario4.Obtener(4) == 4 && diccionario4.Obtener(5) == 5 {
		fmt.Println("recontra nashe 3")
	}
	//Pasa otra prubea,estoy muy a la izquird y mi nodo de remplazo es el propio hijo de mi nodo a borrar y tiene un hijo

	//No tiene mucho sentido probar si estoy a la derecha,como tiene dos hijos,como mucho es su propio hijo el remplazo cosa que ya estuve probando

	diccionario5 := TDADiccionario.CrearABB[int, int](comparar_entero)
	diccionario5.Guardar(7, 6)
	diccionario5.Guardar(10, 6)
	diccionario5.Guardar(3, 3)
	diccionario5.Guardar(5, 5)
	diccionario5.Guardar(4, 4)
	if diccionario5.Borrar(5) == 5 {
		fmt.Println("nashe 4")
	}
	if diccionario5.Obtener(4) == 4 && diccionario5.Obtener(3) == 3 {
		fmt.Println("recontra nashe 4")
	}
	diccionario5.Borrar(3)
	diccionario5.Borrar(7)
	diccionario5.Borrar(4)
	diccionario5.Borrar(10)
	if diccionario5.Cantidad() == 0 {
		fmt.Println("Super nashe 4")
	}

	//Pruebo el caso donde el nodo a borrar solo tiene un hijo

	diccionario6 := TDADiccionario.CrearABB[int, int](comparar_entero)
	diccionario6.Guardar(7, 6)
	diccionario6.Guardar(10, 6)
	diccionario6.Guardar(3, 3)
	diccionario6.Guardar(5, 5)
	if diccionario6.Borrar(3) == 3 {
		fmt.Println("nashe 5")
	}
	if diccionario6.Obtener(7) == 6 && diccionario6.Obtener(5) == 5 {
		fmt.Println("recontra nashe 5")
	}
	//Pasa otra prubea,estoy muy a la izquird y mi nodo de remplazo es el propio hijo de mi nodo a borrar y  no tiene un hijo

	//Aprovecho y pruebo rapido que se pueda volver a usar si borre todo lo que tenia
	diccionario6.Borrar(7)
	if diccionario6.Pertenece(10) && diccionario6.Pertenece(5) {
		fmt.Println("ESto ts bn")
	}
	if diccionario6.Obtener(10) == 6 {
		fmt.Println("Algo ta bn")
	}
	diccionario6.Borrar(10)
	fmt.Println("hasta aca tamo 2")
	diccionario6.Borrar(5)
	if diccionario6.Cantidad() != 0 {
		fmt.Printf("Algo ta mal")
	}
	diccionario6.Guardar(67, 8)
	if diccionario6.Cantidad() != 1 {
		fmt.Println("Algo ta mal 2")
	}
	if !diccionario6.Pertenece(67) {
		fmt.Println("Algo ta mal 3")
	}
	if diccionario6.Obtener(67) == 8 {
		fmt.Println("Ta todo bn")
	}
}
