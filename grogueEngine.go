package main

func checkDestination(x, y int, level *[]Cell, width int) bool {
	return (*level)[x+y*width].crossable

}
