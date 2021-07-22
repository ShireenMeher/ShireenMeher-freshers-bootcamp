package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sum int =0

func student_response(wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	var rdm2 int
	var rdm1 float32
	rdm1=rand.Float32()
	rdm2=rand.Intn(4)+1

	time.Sleep(time.Duration(rdm1))

	mtx.Lock()
	sum+=rdm2
	mtx.Unlock()

}

func main(){
	wg:= new(sync.WaitGroup)
	wg.Add(200)

	mtx:=new(sync.Mutex)

	for i:=0;i<200;i++{
		go student_response(wg,mtx)
	}

	wg.Wait()
	var avg_rating float64
	avg_rating= float64(sum)/float64(200)
	fmt.Println(avg_rating)
}