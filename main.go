package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	Hang time.Duration
)

func init() {
	flag.DurationVar(&Hang, "h", 0, "hang for this time after completing all questions")
}

type Question interface {
	Expected() int
}

type SumQuestion struct {
	S1, S2 int
}

func (q SumQuestion) String() string {
	return fmt.Sprintf("%v+%v", q.S1, q.S2)
}

func (q SumQuestion) Expected() int {
	return q.S1 + q.S2
}

type MulQuestion struct {
	S1, S2 int
}

func (q MulQuestion) String() string {
	return fmt.Sprintf("%vx%v", q.S1, q.S2)
}

func (q MulQuestion) Expected() int {
	return q.S1 * q.S2
}

func main() {
	flag.Parse()
	questions, correct, incorrect := make([]Question, 0, 200), 0, 0
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			questions = append(questions, SumQuestion{i, j}, MulQuestion{i, j})
		}
	}
	reader := bufio.NewReader(os.Stdin)
	for len(questions) > 0 {
		i := rand.Intn(len(questions))
		q := questions[i]
		var actual int
		for {
			fmt.Printf("%v=", q)
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}
			_, err = fmt.Sscanf(text, "%d", &actual)
			if err != nil {
				fmt.Println(err)
				continue
			}
			break
		}
		if actual != q.Expected() {
			incorrect++
			fmt.Printf("Incorrect! %v=%v\n", q, q.Expected())
			continue
		}
		questions[i], questions = questions[len(questions)-1], questions[:len(questions)-1]
		correct++
		fmt.Printf("Correct! (%v left) ", len(questions))
	}
	fmt.Printf("\n\nCorrect: %v, Incorrect: %v\n", correct, incorrect)
	time.Sleep(Hang)
}
