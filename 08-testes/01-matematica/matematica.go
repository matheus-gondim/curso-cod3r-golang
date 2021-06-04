package matematica

//go test ./... => procura todos os testes abaixo do diretorio e executa

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular a media
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	// strconv.ParseFloat => converte uma string em float
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
