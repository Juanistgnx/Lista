package lista_test

import (
	TDAlista "lista/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const PANICO = "La lista esta vacia"

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
