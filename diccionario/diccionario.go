package diccionario

type Diccionario[K comparable, V any] interface {

	//Guardar recibe la clave y su dato correspondiente y los guarda en el diccionario.
	Guardar(clave K, dato V)

	//Pertenece recibe una clave y devuelve true en caso de que esta se encuentre en el diccionario.
	Pertenece(clave K) bool

	//Obtener recibe una clave que esta en el diccionario y devuelve su correspondiente dato.
	//En caso de no estar la clave en el diccionario entra en pànico con un mensaje
	//"La clave no pertenece al diccionario".
	Obtener(clave K) V

	//Borrar recibe una clave que esta en el diccionario y elimina el valor y a propia clave de
	//diccionario.En caso de no estar la clave en el diccionario entra en pànico con un mensaje
	//"La clave no pertenece al diccionario".
	Borrar(clave K) V

	//Cantidad devuelve el numero de claves que tiene guardadas el diccionario.
	Cantidad() int

	//Iterar atravieza el diccionario  aplicando una funcion que recibe una clave y su dato sobre los
	//elementos que recorre. Si sobre el elemento en el que està la funciòn devuelve  true,se seguirà
	//al siguiente elemento,en caso contrario, se detendrà la iteraciòn. En caso de aplicar la funciòn
	//sobre todos los elementos y llegar al final de la lista se detendrà la iteraciòn.
	Iterar(func(clave K, dato V) bool)

	//Iterador devuelve un IterDiccionario del diccionario.
	Iterador() IterDiccionario[K, V]
}
type IterDiccionario[K comparable, V any] interface {

	//HaySiguiente devuelve false si no quedan elementos a iterar en donde esta iterador.
	//En caso contrario,devueve true.
	HaySiguiente() bool

	//VerActual devuelve el dato y la clave del elemento sobre el que esta el iterador.En caso de no
	//haber mas elementos,entra en panico con un mensaje "El iterador termino de iterar".
	VerActual() (K, V)

	//Siguiente mueve el iterador al siguiente elemento del diccionario. En caso de no haber mas
	//elementos,entra en panico con un mensaje "El iterador termino de iterar".
	Siguiente()
}
