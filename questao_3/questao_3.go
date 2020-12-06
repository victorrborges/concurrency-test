package main

import (
	"log"
	"sync"
)

/*
	Definição dos Semáforos
*/

var mutex = newSemaphore(1, 1)    // Mutex para executar operações de incremento nas variáveis drinking/wantsToLeave
var canLeave = newSemaphore(0, 1) // Indica se o estudante pode sair

/*
	Funções do Aluno: Beber e Sair =D
*/

func drink() { log.Print("aluno bebendo") }
func leave() { log.Print("aluno saiu") }

/*
	Solução do Problema
*/

func student() {

	/*
		O mutex se faz necessário para limitar o acesso as variáveis
		do número de estudantes bebendo e o número de estudantes remediados;

		Antes de beber, o aluno checa se há um estudante bebendo e um remediado,
		em caso positivo, ele deixa o estudante remediado sair e decrementa o
		contador de estudantes remediados;

		Depois de beber, o estudante checa:

			Se há um estudante bebendo, ele irá precisar esperar ele terminar;

			Se não há mais nenhum estudante bebendo, o estudante do caso acima
			terminou de beber, assim, os dois podem sair e a variável de estudantes
			remediados é decrementada 2 vezes;

			No último caso, ele apenas pode sair.

	*/

	mutex.wait()
	drinking++
	if drinking == 2 && wantsToLeave == 1 {
		canLeave.signal(1)
		wantsToLeave--
	}
	mutex.signal(1)

	drink()

	mutex.wait()
	drinking--
	wantsToLeave++
	if drinking == 1 && wantsToLeave == 1 {
		mutex.signal(1)
		canLeave.wait()
	} else if drinking == 0 && wantsToLeave == 2 {
		canLeave.signal(1)
		wantsToLeave--
		wantsToLeave--
		mutex.signal(1)
	} else {
		wantsToLeave--
		mutex.signal(1)
	}

	leave()
}

/*
	Main
	obs: desconsiderar, não implementei corretamente (com os Sleeps randomicos)
*/

const nStudents = 10

var drinking = 0     // Contador do número de estudantes comendo
var wantsToLeave = 0 // Contador do número de estudantes remediados

var wg sync.WaitGroup

func main() {
	wg.Add(nStudents)
	for i := 0; i < nStudents; i++ {
		go student()
	}
	wg.Wait()
}
