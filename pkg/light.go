package pkg

import (
	"log"
)

/*
Освещение:

Включить свет
Выключить свет
*/

type Light struct {
	lightController bool
}

func (l *Light) OnnLight() {
	log.Println("[Свет] включен")
	l.lightController = true
}
func (l *Light) OffLight() {
	log.Println("[Свет] выключен")
	l.lightController = false
}
