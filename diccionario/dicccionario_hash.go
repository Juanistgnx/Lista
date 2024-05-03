package diccionario

type diccionarioHash[K comparable, V any] struct {
	//Estructura del hash definir si es abierto o cerrado y como
}

type iterDiccionarioHash[K comparable, V any] struct {
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

}

//Todas las funciones van a tener
/* func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
*/
//Primitivas del Hash
func (hash *diccionarioHash[K, V]) Guardar(clave K, dato V) {

}
func (hash *diccionarioHash[K, V]) Pertenece(clave K) bool {

}
func (hash *diccionarioHash[K, V]) Obtener(clave K) V {

}
func (hash *diccionarioHash[K, V]) Borrar(clave K) V {

}
func (hash *diccionarioHash[K, V]) Cantidad() int {

}
func (hash *diccionarioHash[K, V]) Iterar(func(clave K, dato V) bool) {

}
func (hash *diccionarioHash[K, V]) Iterador() IterDiccionario[K, V] {

}

// Primitivas del iterador
func (iter *iterDiccionarioHash[K, V]) HaySiguiente() bool {

}
func (iter *iterDiccionarioHash[K, V]) VerActual() (K, V) {

}
func (iter *iterDiccionarioHash[K, V]) Siguiente() {
}
