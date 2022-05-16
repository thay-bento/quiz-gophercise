package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

//Read file
func CSVfile(f string) {

}

var wg sync.WaitGroup

func main() {
	var file string
	flag.StringVar(&file, "file", "problems.csv", "set csv file")
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

	timer := time.NewTimer(5 * time.Second)
	ansChannel, q := make(chan string)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question %v: %v = ", questions+1, line[0])

		var answer string
		fmt.Scan(&answer)

		select {
		case ansChannel <- answer:
			if answer == line[1] {
				correct++
			}
		case <-timer.C:
			break
		}

		questions++
	}

	fmt.Printf("Math Quiz! Your score was: %v/%v!\n", correct, questions)

}

/*
Part 1
-> Topic: csvs topic: flags topic: opening files topic: strings
- Programa que lê um arquivo CSV
- Flag onde o usuário pode modificar o arquivo
- Analisar quantas questões estão certas e o total
- Arquivo com a primeira coluna de pergunta e segunda de resposta na mesma linha
Part 2
-> Topic: goroutines topic: channels topic: timers
- Adicionar um timer de 30 segundos e encerrar se não responder.
- Apertar Enter antes de começar o tempo a correr
- Mesmo que a resposta seja errada, passa para a próxima pergunta
- string trimming (string package)
- Nova flag para reordenar as perguntas toda vez que o quiz começar
*/
