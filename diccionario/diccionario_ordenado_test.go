package diccionario_test

import (
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

// No dar bola,solo es para que me queden los import con el go mod
func TestListaVacia(t *testing.T) {
	t.Log("Valido que la lista vacia se comporte como tal")
	diccionario := TDADiccionario.Guardar(dsdd, ddd)
	lista := TDAlista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista no esta sabiendo ver si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.VerPrimero() }, "La lista no deberia poder ver el primero si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.VerUltimo() }, "La lista no deberia poder ver el ultimo si esta vacia")
	require.PanicsWithValue(t, PANICO, func() { lista.BorrarPrimero() }, "La lista no deberia borrar el primero si esta vacia")
}
