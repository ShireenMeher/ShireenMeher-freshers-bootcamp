package main

import "fmt"

type matrix struct{
	rows int
	cols int
	elements [10][10]int
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


func (m matrix) print_matrix(){

	fmt.Println(m)
}
func add_matrices(m1, m2 matrix) matrix{
	const n1  = 10
	const n2 = 10

	var e [n1][n2]int
	//e:= [n1][n2]int{0}
	m:= matrix{m1.rows,m1.cols,e}
	for i:=0;i<m1.rows;i++{
		for j:=0;j<m1.cols;j++{
			m.elements[i][j]=m1.elements[i][j]+m2.elements[i][j]
		}
	}

		return m
}
func main(){
	var e =[10][10]int {
		{1,2,3},
		{4,5,6},{7,8,9}}
	m:= matrix{
		3,
		3,
		e,
	}
	fmt.Println(m.number_of_cols())
	fmt.Println(m.number_of_rows())
	m.set_elements(1,1,6)
	fmt.Println(m.elements[0][0])
	fmt.Println(add_matrices(m,m))
	m.print_matrix()
}



