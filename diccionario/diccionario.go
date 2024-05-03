package diccionario

type Diccionario[K comparable, V any] interface {
	Guardar(clave K, dato V)
	Pertenece(clave K) bool
	Obtener(clave K) V
	Borrar(clave K) V
	Cantidad() int
	Iterar(func(clave K, dato V) bool)
	Iterador() IterDiccionario[K, V]
}
type IterDiccionario[K comparable, V any] interface {
	HaySiguiente() bool
	VerActual() (K, V)
	Siguiente()
}
