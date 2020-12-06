package main

import (
	"log"
	"sync"
)

/*
	Definição dos Semáforos
*/

var questionBarrier = newBarrier(4)    // Barrier utilizado na solução
var mutex = newSemaphore(1, 1)         // Mutex para executar operações de incremento nas variáveis compStuds/psiStuds
var compStudQueue = newSemaphore(0, 1) // Fila dos estudantes de Computação
var psiStudQueue = newSemaphore(0, 1)  // Fila dos estudantes de Psicologia

/*
	Funções Embarcar e Remar
*/

func board(p string) {
	log.Print("um estudante de ", p, " embarcou")
}

func paddle(p string) {
	log.Print("um estudante de ", p, " remou")
	wg.Done()
}

/*
	Solução do Problema

	As duas funções, tanto para estudantes de Computação como
	para estudantes de Psicologia, são simétricas, mudando
	apenas as variáveis que são utilizadas para cada uma;
*/

func compStud() {
	/*
		O mutex se faz necessário para limitar o acesso a variável
		do número de estudantes de Computação que embarcaram,
		impedindo o acesso concorrente;

		Cada thread checa se há um agrupamento válido (2 de cada curso ou
		todos do mesmo curso) para que possam embarcar. Em caso positivo,
		sinaliza as threads necessárias. No caso abaixo, haverá 4 sinalizações
		para a fila de estudantes de computação OU 2 sinalizações para a fila de
		estudantes de computação e 2 sinalizações para a fila de estudantes de
		psicologia;

		O Barrier vai manter o controle de quantas threads embarcaram: quando
		a última thread chegar, todas prosseguem para que o capitão possa
		executar a função de remar e liberar o mutex.

	*/
	isCaptain := false
	mutex.wait()
	compStuds++
	if compStuds == 4 {
		compStudQueue.signal(4)
		compStuds = 0
		isCaptain = true
	} else if compStuds == 2 && psiStuds >= 2 {
		compStudQueue.signal(2)
		psiStudQueue.signal(2)
		psiStuds -= 2
		compStuds = 0
		isCaptain = true
	} else {
		mutex.signal(1) // captain keeps the mutex
	}
	compStudQueue.wait()
	board("computação")
	questionBarrier.wait()
	if isCaptain {
		paddle("computação")
		mutex.signal(1) // captain releases the mutex
	}
}

func psiStud() {
	/*
		O mutex se faz necessário para limitar o acesso a variável
		do número de estudantes de Psicologia que embarcaram,
		impedindo o acesso concorrente;

		Cada thread checa se há um agrupamento válido (2 de cada curso ou
		todos do mesmo curso) para que possam embarcar. Em caso positivo,
		sinaliza as threads necessárias. No caso abaixo, haverá 4 sinalizações
		para a fila de estudantes de psicologia OU 2 sinalizações para a fila de
		estudantes de computação e 2 sinalizações para a fila de estudantes de
		psicologia;

		O Barrier vai manter o controle de quantas threads embarcaram: quando
		a última thread chegar, todas prosseguem para que o capitão possa
		executar a função de remar e liberar o mutex.

	*/
	isCaptain := false
	mutex.wait()
	psiStuds++
	if psiStuds == 4 {
		psiStudQueue.signal(4)
		psiStuds = 0
		isCaptain = true
	} else if psiStuds == 2 && compStuds >= 2 {
		psiStudQueue.signal(2)
		compStudQueue.signal(2)
		compStuds -= 2
		psiStuds = 0
		isCaptain = true
	} else {
		mutex.signal(1) // captain keeps the mutex
	}
	psiStudQueue.wait()
	board("psicologia")
	questionBarrier.wait()
	if isCaptain {
		paddle("psicologia")
		mutex.signal(1) // captain releases the mutex
	}
}

/*
	Main
*/

const nCompStuds = 10 // Número de estudantes de computação utilizados no exemplo
const nPsiStuds = 10  // Número de estudantes de psicologia utilizados no exemplo

var compStuds = 0 // Contador do número de estudantes de computação
var psiStuds = 0  // Contador do número de estudantes de psicologia

var wg sync.WaitGroup

func main() {
	wg.Add((nCompStuds + nPsiStuds) / 4)
	for i := 0; i < nCompStuds; i++ {
		go compStud()
	}
	for i := 0; i < nPsiStuds; i++ {
		go psiStud()
	}
	wg.Wait()
}
