package pkg

import (
	"log"
)

/*
Безопасность:

Включить сигнализацию
Выключить сигнализацию
*/

type DefendSystem struct {
	defendSystemController bool
}

func (ds *DefendSystem) OnnDefend() {
	log.Println("[Безопасность] включена")
	ds.defendSystemController = true
}
func (ds *DefendSystem) OffDefend() {
	log.Println("[Безопасность] выключена")
	ds.defendSystemController = false
}
