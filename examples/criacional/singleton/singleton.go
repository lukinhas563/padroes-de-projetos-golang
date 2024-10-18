package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var singleInstance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("Instância criada")
		singleInstance = &singleton{}
	})

	fmt.Println("Instância já criada")
	return singleInstance
}

func main() {
	single1 := GetInstance()
	single2 := GetInstance()

	fmt.Println(single1 == single2)
}
