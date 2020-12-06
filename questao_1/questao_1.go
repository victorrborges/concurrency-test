package main

import (
	"log"
	"sync"
)

/*
	Definição dos Semáforos
*/

var boardQueue = newSemaphore(0, 1)   // Fila de embarque
var unboardQueue = newSemaphore(0, 1) // Fila de desembarque
var mutex = newSemaphore(1, 1)        // Mutex para executar operações de incremento nas variáveis boardCounter/unboardCounter
var carIsFull = newSemaphore(0, 1)    // Indica se o carro está cheio
var carIsEmpty = newSemaphore(0, 1)   // Indica se o carro está vazio

/*
	Funções da thread Carro
*/

func load()   { log.Print("o carro pode CARREGAR") }
func run()    { log.Print("o carro está CORRENDO") }
func unload() { log.Print("o carro pode DESCARREGAR") }

/*
	Funções das threads Passageiras
*/

func board()   { log.Print("um passageiro embarcou") }
func unboard() { log.Print("um passageiro desembarcou") }

/*
	Solução do Problema
*/

func car() {
	for {
		/*
			O carro começa avisando que pode carregar passageiros;
			Sequencialmente, executa a função de signal C vezes para fila de embarque;
			Por fim, espera até que esteja em sua capacidade máxima.
		*/

		load()
		boardQueue.signal(carCapacity)
		carIsFull.wait()

		/*
			Quando o carro fica cheio, executa a funçao de correr.
		*/

		run()

		/*
			Após executar a funçao de correr, o carro avisa que pode descarregar;
			Executa a função de signal C vezes para a fila de desembarque;
			Por fim, espera até que esteja completamente vazio.
		*/

		unload()
		unboardQueue.signal(carCapacity)
		carIsEmpty.wait()
	}
}

func passenger() {
	/*
		Primeiramente, o passageiro espera até que possa embarcar
		(se norteando pelo estado da fila de embarque) e embarca
	*/

	boardQueue.wait()
	board()

	/*
		O mutex se faz necessário para limitar o acesso a variável
		do número de pessoas que embarcaram, impedindo o acesso concorrente;

		A variável do número de pessoas que embarcaram é incrementada;

		Se o carro está em sua capacidade máxima, um sinal é enviado para
		o semáforo de controle que indica se o carro está cheio e a variável
		do número de pessoas que embarcaram é reiniciada;
	*/

	mutex.wait()
	boardCounter++
	if boardCounter == carCapacity {
		carIsFull.signal(1)
		boardCounter = 0
	}
	mutex.signal(1)

	/*
		O passageiro espera até que possa desembarcar e desembarca
	*/

	unboardQueue.wait()
	unboard()

	/*
		O mutex se faz necessário para limitar o acesso a variável
		do número de pessoas que desembarcaram, impedindo o acesso concorrente;

		A variável do número de pessoas que desembarcaram é incrementada;

		Após o carro ser completamente esvaziado, um sinal é enviado para
		o semáforo de controle que indica se o carro está vazio e a variável
		do número de pessoas que desembarcaram é reiniciada;
	*/

	mutex.wait()
	unboardCounter++
	if unboardCounter == carCapacity {
		carIsEmpty.signal(1)
		unboardCounter = 0
	}
	mutex.signal(1)
	wg.Done()
}

/*
	Main
*/

const carCapacity = 4  // Capacidade do carro
const passengers = 40  // Número de passageiros que irão pegar carona
var boardCounter = 0   // Contador do número de passageiros que embarcaram
var unboardCounter = 0 // Contador do número de passageiros que desembarcaram

var wg sync.WaitGroup

func main() {
	go car()
	wg.Add(passengers)
	for i := 0; i < passengers; i++ {
		go passenger()
	}
	wg.Wait()
}
