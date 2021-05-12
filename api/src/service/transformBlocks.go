package service

import (
	"fmt"
)

func StartTransformBlocks(returnChan chan string) {
	fmt.Println("StartTransformBlocks()")
	for true {
		ans := <-returnChan
		fmt.Println("Transformed: ", ans)
	}
}
