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
	//Close file
	defer file.Close()

	r := csv.NewReader(file)

	ansCorrect := 0
	totalQuestions := 0

	//Set timer

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Question %v: %v = ", totalQuestions+1, line[0])

		var answer string
		fmt.Scan(&answer)

		if answer == line[1] {
			ansCorrect++
		}

		totalQuestions++

	}

	fmt.Printf("Math Quiz! Your score was: %v/%v!\n", ansCorrect, totalQuestions)

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
