package main

import (
	"regexp"
	"os"
	"strconv"
	"errors"
	"bufio"
)

func parseLine(line string) (stud studente, err error) {
	nomeRE := regexp.MustCompile("([a-z]|[A-Z])([a-z]|[A-Z]|[0-9]|_)+")
	votoRE := regexp.MustCompile("[0-9][0-9]")
	
	stud.Nome = nomeRE.FindString(line)
	
	if stud.Nome == "" {
		err = errors.New("Errore: nome mancante")
		return
	}
	
	fineNome := nomeRE.FindStringIndex(line)[1]
	voti := votoRE.FindAllString(line[fineNome:], -1)
	
	for _, x := range voti {
		v, e := strconv.Atoi(x)
		if e != nil {
			err = errors.New("Impossibile convertire il voto " + x)
			return
		} else if v < 18 || v > 32 || v == 31 {
			err = errors.New("Il voto " + x + " è invalido")
			return
		}
		stud.Voti = append(stud.Voti, v)
	}
	
	if len(stud.Voti) == 0 {
		err = errors.New("Errore: ci deve essere almeno un voto per studente")
	}
	
	return
}

func parseFile(fileName string) (stud []studente, err error) {
	f, e := os.Open(fileName)
	if e != nil {
		err = e
		return
	}
	defer f.Close()
	
	b := bufio.NewScanner(f)
	for b.Scan() {
		line := b.Text()
		
		s, e := parseLine(line)
		if e != nil {
			err = e
			return
		}
		
		stud = append(stud, s)
	}
	return
}