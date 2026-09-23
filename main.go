package main

import (
	"fmt"
	"image/color"
	"machine"
	"math"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

const b float64 = 4275.0
const r0 float64 = 100000.0

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
		SDA:       machine.GP0,
		SCL:       machine.GP1,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	machine.InitADC()
	sensor := machine.ADC{Pin: machine.ADC1}
	sensor.Configure(machine.ADCConfig{})

	for {
		v := float64(sensor.Get())
		r := r0 * (65535.0/v - 1.0)
		temp := 1.0/(math.Log(r/r0)/b+1.0/298.15) - 273.15

		display.ClearDisplay()
		white := color.RGBA{255, 255, 255, 255}

		str := fmt.Sprintf("Temp: %0.2f C\n", temp)
		tinyfont.WriteLine(display, &proggy.TinySZ8pt7b, 10, 30, str, white)

		display.Display()

		time.Sleep(time.Second * 10)
	}
}
