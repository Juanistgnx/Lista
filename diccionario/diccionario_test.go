package diccionario_test

import (
	TDADiccionario "diccionario/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	t.Log("Valido que la lista vacia se comporte como tal")

	lista := TDADiccionario.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista no esta sabiendo ver si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.VerPrimero() }, "La lista no deberia poder ver el primero si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.VerUltimo() }, "La lista no deberia poder ver el ultimo si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.BorrarPrimero() }, "La lista no deberia borrar el primero si esta vacia")
}
