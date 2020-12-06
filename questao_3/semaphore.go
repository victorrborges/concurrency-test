package main

/*
	Código do Semáforo
*/

type semaphore chan struct{}

func (s semaphore) signal(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s semaphore) wait() { <-s }

/*
	Função construtora de um Semáforo que, além de definir o
	tamanho, inicializa o Semáforo com o valor determinado
*/

func newSemaphore(initialValue, size int) semaphore {
	s := make(semaphore, size)
	for i := 0; i < initialValue; i++ {
		s.signal(1)
	}
	return s
}
