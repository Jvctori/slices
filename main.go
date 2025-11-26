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

	// Multidimensionals slices:
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		// Aumentando a quantidade de elementos na slice Multidimensionals
		// ...usando o indice do primeiro for, e acrescentando mais um...
		// ... ou seja, primeira iteração a quantidade de elemetnos será 1
		// ... segunda será 2... e a ultima será 3
		// resultando em: [[0][1, 2][2, 3, 4]]

		// innerLen representa a quantidade de elementos criados por elementos na slice
		// representa o tamanho das subslices
		// primeira iteração = 1, pois innerLen = 0 + 1
		// segunda iteração = 2, pois innerLen = 1 + 1
		// terceira iteração = 3, pois innerLen = 2 + 1
		innerLen := i + 1

		// cria uma slice para cada elemento na slice twoD
		// aumentando em um a quantidade de elemento na proxima slice
		// innerLen = i + 1
		// ... a cada iteração do loop externo
		twoD[i] = make([]int, innerLen)

		// primeira iteração do loop de fora(innerLen=1):
		// primerira iteração do loop de dentro:
		// o primeiro e unico valor será 0, pois i=0 e j=0
		// loop interno acaba

		// segunda iteração do loop de fora:
		// primeira iteração do loop de dentro:
		// o primeiro valor da segunda slice será 1
		// pois i=1 e j=0
		// segunda iteração do loop de dentro(innerlen=2):
		// o segundo valor será 2 pois i=1 e j=1
		// loop interno acaba

		// terceira iteração do loop de fora(innerLen=3):
		// primeira iteração:
		// preenche o primeiro valor do terceiro elemento com:
		// 2, pois i=2 e j=0
		// segunda iteração:
		// preenche o segundo valor do terceiro elemento com:
		// 3, pois i=2 e j=1
		// terceira iteração e ultima:
		// preenche o terceiro valor do terceiro elemento com:
		// 4, pois i=2 e j=2
		// loop interno acaba junto com o externo, pois já pecorreu toda a slcie

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// slices operators
	// slice[low:high]
	sliceT := sliceCp[2:5]
	fmt.Println("sliceCp:", sliceCp)
	fmt.Println("sliceT := sliceCp[2:5]:")
	fmt.Println(sliceT)
	// observe que o elemento de indice 5 não é incluído
	// ou seja se quisesse incluí-lo deveria ser [2:6]
	// o elemento de indice 2 foi incluído, isso so acontece com o ultimo elemento que vocẽ deseja incluir!

	// Capacity: cap()
	// quanto a slice pode ser expandida
	// len = 3, mas cap=5
	// len(sliceCp) = 7, cap(sliceCp) = 7
	// ou seja cap(sliceT) = 5, pois é cap(sliceCp) - indice inicial (elementos iniciais declarados [2:5])
	// ou seja, atualmente possui 3 elementos, mas podendo expandir para 5

	fmt.Println("The capacity of the sliceCp is:", cap(sliceCp))
	fmt.Println("The len of sliceCp is:", len(sliceCp))

	fmt.Println("The capacity of the sliceT is:", cap(sliceT))
	fmt.Println("The len of sliceT is:", len(sliceT))
}
