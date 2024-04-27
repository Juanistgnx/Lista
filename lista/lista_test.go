package lista_test

import (
	TDAlista "lista/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const PANICO = "La lista esta vacia"

//PRUEBAS DE LISTA

//t.log() es una funcion para comentar las pruebas. La uso por que el comentario queda gurdado cuando ejecuto las pruebas
//Cualquier cosa no lo uses,comenta y yo despues lo paso

func TestListaVacia(t *testing.T) {
	t.Log("Valido que la lista vacia se comporte como tal")

	lista := TDAlista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista no esta sabiendo ver si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.VerPrimero() }, "La lista no deberia poder ver el primero si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.VerUltimo() }, "La lista no deberia poder ver el ultimo si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.BorrarPrimero() }, "La lista no deberia borrar el primero si esta vacia")
}

func TestListaUnElemento(t *testing.T) {

	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	require.EqualValues(t, 5, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 5, lista.BorrarPrimero())
	require.True(t, true, lista.EstaVacia())

	lista.InsertarUltimo(10)
	require.EqualValues(t, 10, lista.VerPrimero())
	require.EqualValues(t, 10, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 10, lista.BorrarPrimero())
}

func TestComportamientoLista(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(8)
	lista.InsertarUltimo(7)

	require.EqualValues(t, 4, lista.Largo())
	require.EqualValues(t, 10, lista.BorrarPrimero())
	require.EqualValues(t, 5, lista.BorrarPrimero())
	require.EqualValues(t, 8, lista.BorrarPrimero())
	require.EqualValues(t, 7, lista.BorrarPrimero())
	require.True(t, true, lista.EstaVacia())
}

func TestInt(t *testing.T) {
	enteros := []int{58, 65, 21, 46, 9765, 2342, 1}
	lista_int := TDAlista.CrearListaEnlazada[int]()
	for _, value := range enteros {
		lista_int.InsertarUltimo(int(value))
	}
	largo := len(enteros)
	for i := 0; i < largo; i++ {
		require.EqualValues(t, enteros[i], lista_int.BorrarPrimero())
	}
}

func TestStrings(t *testing.T) {
	cadenas := []string{"Hola", "no", "te", "mandes", "una", "cagada", ":)"}
	lista_string := TDAlista.CrearListaEnlazada[string]()
	for _, value := range cadenas {
		lista_string.InsertarUltimo(value)
	}
	largo := len(cadenas)
	for i := 0; i < largo; i++ {
		require.EqualValues(t, cadenas[i], lista_string.BorrarPrimero())
	}
}

func TestFlotantes(t *testing.T) {
	flotantes := []float32{0.1, 45.91829324, 2323.523, 342523452.2334, 453.2423, 22, 34234.352345, 32423.4554342, 2342.23542, 345233.4, 8669.4563567978, 46334.2}
	lista_float := TDAlista.CrearListaEnlazada[float32]()
	for _, value := range flotantes {
		lista_float.InsertarUltimo(float32(value))
	}
	largo := len(flotantes)
	for i := 0; i < largo; i++ {
		require.EqualValues(t, flotantes[i], lista_float.BorrarPrimero())
	}
}

func TestBooleanos(t *testing.T) {
	booleanos := []bool{true, false, true, true, false, true, false, false, false, true, false, false, true, true, true, true, false}
	lista_bool := TDAlista.CrearListaEnlazada[bool]()
	for _, value := range booleanos {
		lista_bool.InsertarUltimo(bool(value))
	}
	largo := len(booleanos)
	for i := 0; i < largo; i++ {
		require.EqualValues(t, booleanos[i], lista_bool.BorrarPrimero())
	}
}

func TestVolumen(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 20000; i++ {
		lista.InsertarPrimero(i)
	}
	for i := 0; i < 20000; i++ {
		require.EqualValues(t, 19999-i, lista.BorrarPrimero())
	}
	require.True(t, true, lista.EstaVacia())
}

func TestBorder(t *testing.T) {
	lista_prueba := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		lista_prueba.InsertarPrimero(int(i))
	}
	for j := 0; j < 10000; j++ {
		lista_prueba.BorrarPrimero()
	}
	require.True(t, lista_prueba.EstaVacia())
	for i := 0; i < 10; i++ {
		lista_prueba.InsertarPrimero(int(i))
	}
	require.EqualValues(t, 9, lista_prueba.VerPrimero())
}

//PRUEBAS ITERADOR

func TestIteradorInicio(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 11; i++ {
		lista.InsertarPrimero(i)
	}
	iterador := lista.Iterador()
	require.EqualValues(t, 10, iterador.VerActual())
}

func TestInsertarFinal(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 11; i++ {
		lista.InsertarPrimero(i)
	}
	iterador0 := lista.Iterador()
	for iterador0.HaySiguiente() {
		iterador0.Siguiente()
	}
	iterador0.Insertar(-1)
	require.EqualValues(t, -1, lista.VerUltimo())
	require.EqualValues(t, 12, lista.Largo())
}

func TestInsertarMedio(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 11; i++ {
		lista.InsertarUltimo(i)
	}
	iterador0 := lista.Iterador()
	for i := 0; i < 4; i++ {
		iterador0.Siguiente()
	}
	iterador0.Insertar(27)
	iterador1 := lista.Iterador()
	for i := 0; i < 4; i++ {
		iterador1.Siguiente()
	}
	require.EqualValues(t, 27, iterador1.VerActual())
	iterador1.Siguiente()
	require.EqualValues(t, 4, iterador1.VerActual())
	require.EqualValues(t, 12, lista.Largo())
}

func TestEliminarPrimero(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	require.EqualValues(t, 0, iterador.Borrar())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 9, lista.Largo())
}

func TestEliminarUltimo(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador1 := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
		iterador1.Siguiente()
	}
	require.EqualValues(t, 9, iterador1.Borrar())
	require.EqualValues(t, 8, lista.VerUltimo())
	require.EqualValues(t, 9, lista.Largo())
}

func TestEliminarElemento(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	for i := 1; i < 8; i++ {
		iterador.Siguiente()
	}
	dato := iterador.Borrar()
	presencia := false
	lista.Iterar(func(x int) bool {
		presencia = (x == dato)
		return !presencia
	})
	require.False(t, presencia)
}
