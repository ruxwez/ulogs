package main

import (
	"github.com/ruxwez/ulogs/v1"
)

func main() {
	logs := ulogs.NewInstance(&ulogs.Config{
		Name:       "ULogs Name",
		FilePrefix: "ULogs",
		DiscordURL: "https://discord.com/api/webhooks/1201485312694693923/qq-7erzCWkMDiCInHSdawd6jBe851DHadawdawdm1VRF0Meu6tg6ZLt-h_K7oirQhBtL-tSHdadawd9CLoFXtQ",
	})

	logs.Send(ulogs.TYPE_INFO, "Esto es un log de prueba")
}
