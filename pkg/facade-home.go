package pkg

import "fmt"

//
/*
Фасад
Создайте фасад, который будет предоставлять простые методы для управления всеми
подсистемами.
Например, методы для "Уйти из дома" и "Вернуться домой", которые будут автоматически
управлять освещением,
климат-контролем и системой безопасности.
*/

type Facade struct {
	Light          *Light
	ClimateControl *ClimateControl
	DefendSystem   *DefendSystem
}

func (f *Facade) GetOutOfHome() {
	f.Light.OnnLight()
	f.DefendSystem.OnnDefend()
	f.ClimateControl.GetTemperature(15)
	fmt.Printf(
		"Все системы включены.\n [Защита:] %t\n [Свет:] %t\n [Температура:]%v°C\n Хорошего дня!\n", f.DefendSystem.defendSystemController, f.Light.lightController, f.ClimateControl.temperature)
}
func (f *Facade) GoHome() {
	f.Light.OffLight()
	f.DefendSystem.OffDefend()
	f.ClimateControl.GetTemperature(18)
	fmt.Printf(
		"Все системы включены,\n [Защита:] %t\n [Свет:] %t\n [Температура:]%v°C\n Добро пожаловать домой\n", f.DefendSystem.defendSystemController, f.Light.lightController, f.ClimateControl.temperature)
}
