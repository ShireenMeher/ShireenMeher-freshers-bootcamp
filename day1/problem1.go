package main

import (
	"encoding/json"
	"fmt"
)

type matrix struct{
	rows int
	cols int
	elements [][]int
}

func (m matrix) number_of_rows() int{
	return m.rows
}

func (m matrix) number_of_cols() int{
	return m.cols
}

func (m matrix) set_elements(i,j,element int){
	m.elements[i][j]=element
}


func (m *matrix) print_matrix(){
	data , err := json.MarshalIndent(m,""," ")

	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
func (m *matrix) add_matrices(addMatrix matrix) [][]int{

	for i:=0 ; i<len(m.elements) ; i++ {
		for j:=0 ; j<len(m.elements[0]) ; j++ {
			m.elements[i][j] = m.elements[i][j] + addMatrix.elements[i][j]
		}
	}

	return m.elements
}
func main(){
	var e =[][]int {
		{1,2,3},
		{4,5,6},{7,8,9}}

	var e2 =[][]int {
		{2,2,2},
		{2,2,2},{2,2,2}}
	m:= matrix{
		3,
		3,
		e,
	}

	m2:= matrix{
		3,
		3,
		e2,
	}
	fmt.Println(m.number_of_cols())
	fmt.Println(m.number_of_rows())
	m.set_elements(1,1,6)
	fmt.Println(m.elements[0][0])
	fmt.Println(m.add_matrices(m2))
	m2.print_matrix()
}