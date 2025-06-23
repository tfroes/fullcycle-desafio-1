package cotacaoservice

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func BuscaCotacao(ctx context.Context) (*Cotacao, error) {
	req, error := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if error != nil {
		return nil, error
	}

	client := &http.Client{}
	resp, error := client.Do(req)

	if error != nil {
		return nil, error
	}

	body, error := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if error != nil {
		return nil, error
	}

	var c Cotacao
	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}

	return &c, nil
}
