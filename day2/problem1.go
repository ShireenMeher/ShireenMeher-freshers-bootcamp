package main

import "fmt"

func main()  {
	slc := make([]string,0)
	var n int
	fmt.Scanln(&n)

	for i:=0;i<n;i++{
		var temp string
		fmt.Scanln(&temp)
		slc=append(slc, temp)
	}

	mapp:=make(map[byte]int)
	for i:=0;i<n;i++{
		go frequency(slc[i],mapp)
	}

	for key,value := range mapp{
		fmt.Println(string(key)," : ", value)
	}
}

func frequency(str string, mapp map[byte]int){
	var n int=len(str)

	for i:=0;i<n;i++{
		_,prs := mapp[str[i]]
		if prs{
			mapp[str[i]] = mapp[str[i]]+1
			continue
		}
		mapp[str[i]]=10
	}
}