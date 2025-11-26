package main

import (
	"fmt"
	"slices"
)

// Estudos de slice
func main() {

	// var sliceName[]ElementType
	// note que não declaramos uma quantidade de elementos entre []
	var numbers []int                   // slice nil // slice declarada sem valores
	var numbers1 = []int{1, 2, 3, 4, 5} // slice declarada com valores definidos

	numbers2 := []int{8, 9, 10} // forma mais "enxuta" de declarar uma slice

	fmt.Println("Numbers:", numbers)
	fmt.Println("Numbers1:", numbers1)
	fmt.Println("numbers2:", numbers2)

	slice := make([]int, 5) // slice definida com uma quantidade de elementos usando make()

	fmt.Println("Slice:", slice)

	slice[0] = 2020

	fmt.Println("Slice after alteration:", slice)

	// array:
	a := [5]int{1, 2, 3, 4, 5}

	// Pegando uma parte da array e criando uma slice:
	slice1 := a[1:4]

	fmt.Println("Array:", a)
	fmt.Println("Slice1:", slice1)

	// Alterando valores de elementos da slice
	slice1[0] = 20312

	fmt.Println("SLice1 after alteration:", slice1)
	slice1 = append(slice1, 10, 20, 30, 40)

	fmt.Println("Slice1 after add some numbs:", slice1)

	// FAZENDO COPIA DE UMA SLICE:
	// para copiarmos uma slice precisamops declarar

	// uma slice declarada de tamanho definido preenche os valores com 0:
	fmt.Println("Exemplo:")

	sliceCp := make([]int, len(slice1))

	fmt.Println("sliceCp antes de ter os seus valores copiados de slice1:", sliceCp)

	// copy copia os valores
	// copy(slice_que_vai_receber_os_valores, slice_que_vai_ser_copiada)
	// copy(src, dest)
	copy(sliceCp, slice1)

	fmt.Println("Copia de slice1 -> sliceCopy:", sliceCp)

	// Nil slices, ou seja, slices sem valores
	// slice nil
	var nilSlice []int

	fmt.Println(nilSlice)

	for i, v := range slice1 {
		fmt.Println("Index:", i)
		fmt.Println("Value:", v)
	}
	fmt.Println(slice1[0])

	slice1[0] = 15

	fmt.Println(slice1[0])

	// Comparando slices
	// slices.Equal(slice_a_ser_comparada, slice_de_comparação)
	// slices.Equal(s1, s2) bool
	if slices.Equal(slice1, sliceCp) {
		fmt.Println("Slice1 is equal to sliceCp")
	} else {
		fmt.Println("Slice1 is not equal to sliceCp")
	}

	sliceBool := slices.Equal(slice1, sliceCp)
	fmt.Println("É verdade que slice1 é igual a sliceCp?", sliceBool)
}
