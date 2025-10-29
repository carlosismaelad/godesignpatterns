package generator

import (
	"math/rand"
	"time"

	"carlosismaelad.com/godesignpatterns/generator/model"
)

// LogGenerator is a Generator that sends logs to a channel
func LogGenerator() <-chan model.Log {
	out := make(chan model.Log)

	go func() {
		defer close(out)
		levels := []string{"INFO", "WARN", "ERROR"}
		messages := []string{
			"Usuário fez login",
			"Conexão com banco estabelecida",
			"Falha ao processar pagamento",
			"Cache limpo com sucesso",
			"Token expirado",
		}

		for {
			// Generates a random log
			log := model.Log{
				Level:   levels[rand.Intn(len(levels))],
				Message: messages[rand.Intn(len(messages))],
				Time:    time.Now(),
			}

			out <- log // send the log to the channel
			time.Sleep(time.Second) // 1s pause
		}
	}()

	return out
}