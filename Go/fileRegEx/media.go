package main

func mediaVoti(stud studente) float64{
	somma := 0
	for i := 0; i < len(stud.Voti); i++ {
		if stud.Voti[i] == 32 {
			somma += 30
		} else {
			somma += stud.Voti[i]	
		}
	}
	return float64(somma)/float64(len(stud.Voti))
}
