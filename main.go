package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

//Read file
func CSVfile(f string) {
	//Open file
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)

	ansCorrect := 0
	ansIncorrect := 0
	nQuestions := 1
	totalQuestions := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Question %v: %v = ", nQuestions, record[0])
		var answer string

		fmt.Scan(&answer)
		if answer == record[1] {
			ansCorrect++
		} else {
			ansIncorrect++
		}
		totalQuestions++
	}
	fmt.Printf("Math Quiz! Your score was: %v/%v!\n", ansCorrect, totalQuestions)
	if ansCorrect > ansIncorrect {
		fmt.Println("Congratulations!")
	}
	fmt.Printf("Correct: %v\nIncorrect: %v\n", ansCorrect, ansIncorrect)

}

func main() {
	var file string
	flag.StringVar(&file, "file", "problems.csv", "set csv file")
	flag.Parse()

	//Function to Read File
	CSVfile(file)

}

/*
- Adicionar um timer de 30 segundos e encerrar se não responder.
- Apertar Enter antes de começar o tempo a correr
- Mesmo que a resposta seja errada, passa para a próxima pergunta
- string trimming (string package)
- Nova flag para reordenar as perguntas toda vez que o quiz começar
*/
