package main

import "linkedList"

func main() {
	var lista linkedList.List

	lista = linkedList.AddFront("foglia", lista)
	lista = linkedList.AddFront("dentice", lista)
	lista = linkedList.AddFront("abete", lista)

	lista = linkedList.AddTail("lavatrice", lista)

	lista = linkedList.AddInOrder("elefante", lista)
	lista = linkedList.AddInOrder("aaron", lista)

	linkedList.PrintAllElements(lista)
}
