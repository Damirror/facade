package pkg

import (
	"log"
)

/*
Климат-контроль:

Установить температуру
*/

type ClimateControl struct {
	temperature float64
}

func (cc *ClimateControl) GetTemperature(temperature float64) {
	cc.temperature = temperature
	log.Printf("Темература задана на отметку %v\n", temperature)
}
