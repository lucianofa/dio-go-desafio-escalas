package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	k := ebulicaoK
	c := k - 273.0 

	fmt.Printf("Temperatura de ebulição da água em Kelvin (ºK) = %g , Temperatura de ebulição da água em Celsius (ºC) = %g.\n", k, c)

}

