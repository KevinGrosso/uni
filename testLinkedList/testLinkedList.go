package main

import (
		"linkedList"
		"fmt"
)

func main() {
	var lista linkedList.List

	lista = linkedList.AddFront("foglia", lista)
	lista = linkedList.AddFront("dentice", lista)
	lista = linkedList.AddFront("abete", lista)

	lista = linkedList.AddTail("lavatrice", lista)

	lista = linkedList.AddInOrder("elefante", lista)
	lista = linkedList.AddInOrder("aaron", lista)

	linkedList.ItPrintAllElementsNL(lista)
	linkedList.ItPrintAllElementsArr(lista)

	lista = linkedList.AddAtIndex("gallo", 0, lista)

	len := linkedList.CountAllElements(lista)
	lista = linkedList.AddAtIndex("zanzara", len, lista)
	lista = linkedList.AddAtIndex("forza", 4, lista)
	linkedList.ItPrintAllElementsArr(lista)

	lista = linkedList.RemoveAtIndex(0, lista)
	linkedList.ItPrintAllElementsArr(lista)
	lista = linkedList.RemoveAtIndex(3, lista)
	linkedList.ItPrintAllElementsArr(lista)
	last := linkedList.CountAllElements(lista) - 1
	lista = linkedList.RemoveAtIndex(last, lista)
	linkedList.ItPrintAllElementsArr(lista)


	linkedList.RecPrintAllElementsNL(lista)
	linkedList.ItPrintAllElementsArr(lista)

	fmt.Println("Elementi nella lista:", linkedList.CountAllElements(lista))
}
