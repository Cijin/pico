package main

import (
	"machine"
	"runtime"
	"time"
)

func main() {
	esp := machine.UART0
	esp.Configure(machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.GP16,
		RX:       machine.GP17,
	})

	host := machine.UART1
	host.Configure(machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.GP8,
		RX:       machine.GP9,
	})

	go func() {
		led := machine.LED
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
		for {
			led.Set(!led.Get())
			time.Sleep(time.Second)
		}
	}()

	for {
		for esp.Buffered() > 0 {
			b, _ := esp.ReadByte()
			host.WriteByte(b)
		}
		for host.Buffered() > 0 {
			b, _ := host.ReadByte()
			esp.WriteByte(b)
		}

		runtime.Gosched()
	}
}
