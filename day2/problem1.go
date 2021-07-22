package main

import (
	"fmt"
	"sort"
	"sync"
)

type mymap struct{
	mapp map[string]int
	mtx sync.Mutex
}

func increament(mymap2 *mymap, a byte){
	mymap2.mtx.Lock()
	_,prs := mymap2.mapp[string(a)]

	if prs{
		mymap2.mapp[string(a)] = mymap2.mapp[string(a)]+1

	}
	if !prs {
		mymap2.mapp[string(a)] = 1
	}
	mymap2.mtx.Unlock()
}

func main()  {
	slc := make([]string,0)
	wg:= new(sync.WaitGroup)


	var n int

	fmt.Scanln(&n)
	wg.Add(n)
	for i:=0;i<n;i++{
		var temp string
		fmt.Scanln(&temp)
		slc=append(slc, temp)
	}

	mymap2:=mymap{make(map[string]int),*new(sync.Mutex)}
	for i:=0;i<n;i++{
		go frequency(slc[i],&mymap2,wg)
	}

	wg.Wait()

	sorted_keys:=make([]string,0)
	for k:= range mymap2.mapp{
		sorted_keys=append(sorted_keys,k)
	}
	sort.Strings(sorted_keys)

	for _, k := range sorted_keys {
		fmt.Println(k," : ", mymap2.mapp[k])
	}

}

func frequency(str string, mymap2 *mymap, wg *sync.WaitGroup){
	var n int=len(str)

	defer wg.Done()
	for i:=0;i<n;i++{
		increament(mymap2,str[i])

	}
}