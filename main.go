package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Question interface {
	Expected() int
}

type SumQuestion struct {
	S1, S2 int
}

func (q SumQuestion) String() string {
	return fmt.Sprintf("%v + %v", q.S1, q.S2)
}

func (q SumQuestion) Expected() int {
	return q.S1 + q.S2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		s1 := rand.Intn(10)
		s2 := rand.Intn(10)
		q := SumQuestion{s1, s2}
		fmt.Printf("%v = ", q)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		var actual int
		_, err = fmt.Sscanf(text, "%d", &actual)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if actual != q.Expected() {
			fmt.Printf("Incorrect! %v = %v\n", q, q.Expected())
			continue
		} else {
			fmt.Printf("Correct! ")
		}
	}
}
