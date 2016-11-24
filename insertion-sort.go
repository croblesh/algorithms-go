package main

import "fmt"

func main(){

	numbersToBeSorted := [7]int {6, 4, 1, 10, 5, 2, 22}

	for j := 1; j < len(numbersToBeSorted); j++ {

		key := numbersToBeSorted[j]
		i := j - 1

		for i >= 0 && numbersToBeSorted[i] > key {
			numbersToBeSorted[i+1] = numbersToBeSorted[i]
			i = i - 1
		}

		numbersToBeSorted[i+1] = key

	}

	fmt.Println(numbersToBeSorted)

}
