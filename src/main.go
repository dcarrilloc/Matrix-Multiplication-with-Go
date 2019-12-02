package main

import "fmt"

func main(){
	// creating matrix
	mat1 := [3][3] float32 {{1,2,3},{4,5,6},{7,8,9}}
	mat2 := [3][3] float32 {{5,2,3},{7,0,1},{2,5,2}}
	var result[len(mat1)][len(mat1)] float32

	//calculating multiplication
	for i := 0; i < len(result) ; i++  {
		for j := 0; j < len(result) ; j++  {
			for k := 0; k < len(result); k++  {
				result[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	// printing result
	fmt.Print(result)
}
