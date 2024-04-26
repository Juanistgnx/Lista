package lista_test

import (
	TDAlista "lista/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista no esta sabiendo ver si esta vacia")
}
