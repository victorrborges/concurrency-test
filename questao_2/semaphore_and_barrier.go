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

/*
	Código do Barrier
*/

type barrier struct {
	n, count                           int
	mutex, mainSemaphore, auxSemaphore semaphore
}

/*
	Função construtora de um Barrier
*/

func newBarrier(n int) *barrier {
	return &barrier{
		n:             n,                  // Número de threads
		count:         0,                  // Contador auxiliar
		mutex:         newSemaphore(1, 1), // Mutex para executar operações de incremento/decremento na variável count
		mainSemaphore: newSemaphore(0, 1), // Semáforo principal
		auxSemaphore:  newSemaphore(0, 1), // Semáforo auxiliar
	}
}

func (b *barrier) wait() {

	/*
		O mutex se faz necessário para limitar o acesso a
		variável count, impedindo o acesso concorrente;

		A variável count é incrementada;

		Caso essa seja a última thread, o count será
		igual ao número máximo de threads, assim, o
		semáforo é liberado para que as threads  possam
		chegar nos seus respectivos pontos críticos;

		Caso não seja a última thread, a thread é bloqueada
		até que todas as threads incrementem count.
	*/

	b.mutex.wait()
	b.count++
	if b.count == b.n {
		b.mainSemaphore.signal(b.n)
	}
	b.mutex.signal(1)
	b.mainSemaphore.wait()

	/*
		Segunda etapa auxiliar, responsável
		por decrementar o valor de count.
	*/

	b.mutex.wait()
	b.count--
	if b.count == 0 {
		b.auxSemaphore.signal(b.n)
	}
	b.mutex.signal(1)
	b.auxSemaphore.wait()
}
