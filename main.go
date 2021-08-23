package main

import (
	"fmt"
	"uni/repository"
	"uni/utils"
)

func main() {
	repository.Init()

	utilsInterface := utils.NewCuorseInterfaceUtils()
	idPointer, err := utilsInterface.CreateCuorse("sp")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("your new cuorse id is : ", *idPointer)

	dataPointer, err := utilsInterface.GetCuorseData(*idPointer)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}


	fmt.Println("your requested cuorse data: ", *dataPointer)

}
