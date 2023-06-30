package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo11.txt")

	//if the file is found it gives nil
	// if there is no such a file it will give error

	//demo1.txt should be in the example file not into the subfiles
	//why?
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok { //type assertion
			fmt.Println("there is no such a file", pErr.Path)
			return
			//there is a path error
		} else {
			fmt.Println("there is no such a file")
			return
		}

	} else {
		fmt.Println(f.Name())
	}

}
