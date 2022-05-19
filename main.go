package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "problems.csv", "set csv file")
	timeLimit := flag.Int("time", 30, "set time limit")
	flag.Parse()

	//Open file
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	//Close file
	defer csvfile.Close()

	r := csv.NewReader(csvfile)

	correct := 0
	questions := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		ansChannel := make(chan string)
		go func() {
			fmt.Printf("Question %v: %v = ", questions+1, line[0])
			var answer string
			fmt.Scan(&answer)
			ansChannel <- strings.TrimSpace(answer)
		}()

		select {
		case <-timer.C:
			os.Exit(1)
		case answer := <-ansChannel:
			if answer == line[1] {
				correct++
			}
		}

		questions++
	}

	fmt.Printf("Math Quiz: Your score was: %v/%v!\n", correct, questions)

}

/*
Part 1
-> Topic: csvs topic: flags topic: opening files topic: strings
- Programa que lê um arquivo CSV - OK
- Flag onde o usuário pode modificar o arquivo - OK
- Analisar quantas questões estão certas e o total - OK
- Arquivo com a primeira coluna de pergunta e segunda de resposta na mesma linha - OK
Part 2
-> Topic: goroutines topic: channels topic: timers
- Adicionar um timer de 30 segundos e encerrar se não responder. - OK
- Mesmo que a resposta seja errada, passa para a próxima pergunta - OK
- string trimming (string package) - OK
- Nova flag para reordenar as perguntas toda vez que o quiz começar
*/
