package tempconv

/*Упражнение 2.1. Добавьте в пакет tem pconv типы, константы и функции для работы с температурой по шкале
Кельвина, в которой нуль градусов соответствует температуре-273.15°С,
а разница температур в 1К имеет ту же величину, что и 1°С.*/

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

const (
	AbsoluteZeroC      Celsius = -273.15
	FreezingC          Celsius = 0
	BoilingC           Celsius = 100
	AbsoluteZeroKelvin Kelvin  = -273.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KtoC(k Kelvin) Celsius     { return Celsius(k) + AbsoluteZeroC }
