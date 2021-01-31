package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type data struct{
	question string
	answer string
}

func main() {
	problem, err := os.Open("problems.csv")
	if err!=nil{
		fmt.Println("Could not open problems.csv")
	}
	columns, err := csv.NewReader(problem).ReadAll()
	if err!=nil{
		fmt.Println("Could not read problems.csv")
	}
	point := 0
	for _, column := range columns{
		emp := data{
			question: column[0],
			answer: column[1],
		}
		fmt.Println("what is " + emp.question + "?")
		var ans string
		ch := make(chan int)
		go func(){
			fmt.Scanln(&ans)
			ch <- 1
		}()
		select{
		case <- ch:
			if ans == emp.answer{
				point += 1
			}
		case <-time.After(3*time.Second):
			if ans == ""{
				ans = " "
			}
		}
	}
	fmt.Printf("Total Point: %d\n", point)
}