package main

import (
	"client-server/internal/client/cotacaoapi"
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {

	fileCotacao, err := os.OpenFile("cotacao.txt", os.O_APPEND, 0644)
	//os.Create("cotacao.txt")

	if err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		panic("Erro ao criar arquivo de cotação")
	}

	defer fileCotacao.Close()

	for {

		var ctx, cancel = context.WithTimeout(context.Background(), 300*time.Millisecond)
		defer cancel()

		cotacao, err := cotacaoapi.GetCotacao(ctx)

		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("[Error] Requisição da api com Timeout!\n")
			continue
		}

		if err != nil {
			fmt.Printf("[Error] %s\n", err.Error())
			continue
		}

		fileCotacao.WriteString(fmt.Sprintf("Dólar: %s\n", cotacao.USDBRL.Bid))

		fmt.Printf("%s\n", cotacao.USDBRL.Bid)

		time.Sleep(500 * time.Millisecond)
	}
}
