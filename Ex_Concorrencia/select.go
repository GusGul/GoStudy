// Select funciona como um Switch para canais.
// Select permite que você aguarde várias operações em canais.
// Select é bloqueado até que um dos seus casos possa ser executado.
// Combinar goroutines e canais com select é um recurso poderoso do Go.
// Para nosso exemplo, selecionaremos em dois canais.
package main

// Cada canal receberá um valor após algum tempo, para simular, por exemplo
// o bloqueio de operações RPC em execuçãop em goroutines concorrentes.
import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		// Após 1 segundo, envie "um" para c1
		time.Sleep(time.Second * 1)
		//Observe que o tempo total de execução é de apenas ~2 segundos,
		//pois o 1 e o 2 segundos são Sleeps executados simultaneamente.
		c1 <- "um"
	}()
	go func() {
		// Após 2 segundos, envie "dois" para c2
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	for i := 0; i < 2; i++ {
		// Select aguarda simultaneamente os valores de c1 e c2.
		// Caso seja recebido um valor em um dos canais, o select
		// será executado.
		select {
		case msg1 := <-c1:
			fmt.Println("recebido ", msg1)
		case msg2 := <-c2:
			fmt.Println("recebido ", msg2)
		}
	}
}
