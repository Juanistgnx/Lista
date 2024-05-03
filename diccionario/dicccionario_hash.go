package diccionario

type estadoNodo int

const (
	VACIO = estadoNodo(iota)
	OCUPADO
	BORRADO
)

const LARGOINICIAL int = 13

type nodoHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoNodo
}

type hashCerrado[K comparable, V any] struct {
	contenido []nodoHash[K, V]
	tamaño    int
	cantidad  int
	borrados  int
}

type iterDiccionarioHash[K comparable, V any] struct {
	tabla    *[]nodoHash[K, V]
	posicion int
	//Este iterador se debe mover hasta que encuentre algo
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	nuevo := new(hashCerrado[K, V])
	nuevo.contenido = make([]nodoHash[K, V], LARGOINICIAL)
	nuevo.tamaño = LARGOINICIAL
	nuevo.cantidad, nuevo.borrados = 0, 0
	return nuevo
}

//Todas las funciones van a tener
/* func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
*/
//Primitivas del Hash
func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

}
func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {

}
func (hash *hashCerrado[K, V]) Obtener(clave K) V {

}
func (hash *hashCerrado[K, V]) Borrar(clave K) V {

}
func (hash *hashCerrado[K, V]) Cantidad() int {

}
func (hash *hashCerrado[K, V]) Iterar(func(clave K, dato V) bool) {

}
func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

}

// Primitivas del iterador
func (iter *iterDiccionarioHash[K, V]) HaySiguiente() bool {

}
func (iter *iterDiccionarioHash[K, V]) VerActual() (K, V) {

}
func (iter *iterDiccionarioHash[K, V]) Siguiente() {
}
