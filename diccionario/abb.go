package diccionario

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}
type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}
type funcCmp[K comparable] func(K, K) int

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}
func (abb *abb[K, V]) Guardar(clave K, dato V) {

}

func (abb *abb[K, V]) Pertenece(clave K) bool {

}

func (abb *abb[K, V]) Obtener(clave K) V {

}

func (abb *abb[K, V]) Borrar(clave K) V {

}

func (abb *abb[K, V]) Cantidad() int {

}

func (abb *abb[K, V]) Iterar(func(clave K, dato V) bool) {

}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {

}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

}
