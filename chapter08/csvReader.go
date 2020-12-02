package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("test.csv")
	defer file.Close()

	csvReader := csv.NewReader(file)

	//for {
	//	line, err := csvReader.Read()
	//	if err !=nil {
	//		if err == io.EOF {
	//			break
	//		}
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(line)
	//}

	all, _ := csvReader.ReadAll()
	fmt.Println(all)

}
