package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

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

	display.ClearDisplay()
	white := color.RGBA{255, 255, 255, 255}
	// Draw text to the screen buffer
	// X=10, Y=30 (Y is the bottom baseline of the text)
	tinyfont.WriteLine(display, &proggy.TinySZ8pt7b, 10, 30, "Hello, Maker Pi!", white)

	display.Display()

	for {
		time.Sleep(time.Hour)
	}
}
