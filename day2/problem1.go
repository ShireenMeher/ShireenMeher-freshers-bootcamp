package main

import (
	"fmt"
	"sort"
)

func main()  {
	slc := make([]string,0)
	var n int
	fmt.Scanln(&n)

	for i:=0;i<n;i++{
		var temp string
		fmt.Scanln(&temp)
		slc=append(slc, temp)
	}

	mapp:=make(map[string]int)
	for i:=0;i<n;i++{
		go frequency(slc[i],mapp)
	}

	sorted_keys:=make([]string,len(mapp))
	for k:= range mapp{
		sorted_keys=append(sorted_keys,k)
	}
	sort.Strings(sorted_keys)

	for _, k := range sorted_keys {
		fmt.Println(k," : ", mapp[k])
	}
}

func frequency(str string, mapp map[string]int){
	var n int=len(str)

	for i:=0;i<n;i++{
		_,prs := mapp[string(str[i])]
		if prs{
			mapp[string(str[i])] = mapp[string(str[i])]+1
			continue
		}
		mapp[string(str[i])]=1
	}
}