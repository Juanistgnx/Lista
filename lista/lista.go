package lista

type Lista[T any] interface {

	//EstaVacia devuelve true si la lista no contiene elementos,false en caso contrario
	EstaVacia() bool

	//InsertarPrimero agrega un nuevo elemento a la lista en la primera posiciòn de la misma
	InsertarPrimero(T)

	//InsertarUltimo agrega un nuevo elemento a la lista en la ultima posiciòn de la misma
	InsertarUltimo(T)

	//BorrarPrimero elimina el primer elemento de la lista. Si tiene elementos,se quita el primero y
	//devuelve su valor. Si la lista esta vacìa, entra en pánico con un mensaje "La lista esta vacia"
	BorrarPrimero() T

	//VerPrimero obtiene el valor del primer elemento de la lista. Si está vacía, entra en pánico con
	//un mensaje "La lista esta vacia".
	VerPrimero() T

	//VerUltimo obtiene el valor del ultimo elemento de la lista. Si está vacía, entra en pánico con
	//un mensaje "La lista esta vacia".
	VerUltimo() T

	//Largo devuelve la cantidad de elementos que tiene la lista
	Largo() int

	//Iterar atravieza la lista desde el principio aplicando la funciòn "visitar" sobre los elementos
	//que recorre. Si sobre el elemento en el que està la funciòn "visitar" devuelve true,se seguirà
	//al siguiente elemento,en caso contrario, se detendrà la iteraciòn. En caso de aplicar la funciòn
	//sobre todos los elementos y llegar al final de la lista se detendrà la iteraciòn
	Iterar(visitar func(T) bool)

	//Iterador devuelve un IteradorLista de la lista. Este es inicializado al principio de la lista
	//,es decir,se encuentra sobre el primer elemento de la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	//VerActual devuelve el dato del elemento sobre el que esta el iterador.En caso de estar al final
	//de la lista,entra en panico con un mensaje "El iterador termino de iterar"
	VerActual() T

	//HaySiguiente devuelve false si estoy sobre el final de la lista,es decir,no existen mas elementos
	//en donde esta iterador.En caso contrario,devueve true
	HaySiguiente() bool

	//Siguiente mueve el iterador al siguiente elemento de la lista. En caso de estar al final de la
	//lista,entra en panico con un mensaje "El iterador termino de iterar"
	Siguiente()

	//Insertar agrega un nuevo elemento a la lista en la posiciòn en la que se encuentra el iterador
	Insertar(T)

	//Borrar elimina el elemento sobre el que esta el iterador. Si el iterador no esta al final de la
	//lista,elimina el elemento y devuelve su valor.En caso contrario,entra en panico con un mensaje
	//"El iterador termino de iterar"
	Borrar() T
}
