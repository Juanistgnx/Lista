package diccionario_test

import (
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

const PANICO = "La clave no pertenece al diccionario"
const PANICOITER = "El iterador termino de iterar"

// No dar bola,solo es para que me queden los import con el go mod
func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Valido que el diccionario se comporte correctamente cuando no tiene elementos y guarde bien")
	diccionario := TDADiccionario.CrearABB[string, int](strings.Compare)

	require.Equal(t, 0, diccionario.Cantidad(), "No se le ingreso ningun elemento,por ende su cantidad deberia ser 0")
	require.PanicsWithValue(t, PANICO, func() { diccionario.Obtener("hola") }, "El diccioanario esta vacio,por ende no deberia registar esta o ninguna clave")
	require.PanicsWithValue(t, PANICO, func() { diccionario.Borrar("hola") }, "El diccionario esta vacio,por ende no deberia registar esta o ninguna clave")

	diccionario.Guardar("hola", 10)

	require.Equal(t, 1, diccionario.Cantidad(), "El diccionario deberia poder actualizar su cantidad cuando le ingreso una clave")
	require.True(t, diccionario.Pertenece("hola"), "El diccionario deberia registrar la clave recièn ingresada")
	require.Equal(t, 10, diccionario.Obtener("hola"), "El diccionario deberia poder devolver el dato asociado a su clave")

	diccionario.Guardar("mundo", 20)

	require.Equal(t, 2, diccionario.Cantidad(), "El diccionario deberia poder actualizar su cantidad cuando le ingreso una nueva clave")
	require.True(t, diccionario.Pertenece("mundo"), "El diccionario deberia registrar la clave recièn ingresada")
	require.Equal(t, 20, diccionario.Obtener("mundo"), "El diccionario deberia poder devolver el dato asociado a su clave")

	require.True(t, diccionario.Pertenece("hola"), "El diccionario no deberia perder la primer clave ingresada")
	require.Equal(t, 10, diccionario.Obtener("hola"), "El diccionario no deberia perder el dato asociado a la clave ingresada")

}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}
