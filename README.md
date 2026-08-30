# Notes

## Links:
1. Examples: https://github.com/raspberrypi/pico-examples
2. Pico-SDK: https://github.com/raspberrypi/pico-sdk


# Raspberry Pi Pico Debug Probe Setup (macOS Apple Silicon)

## Prerequisites

Complete the C/C++ SDK setup first. This guide assumes `pico-sdk` and the ARM toolchain are already installed.

## 1. Wire the Hardware

Connect the debug probe to the target Pico.

- Connect SWCLK, SWDIO, and GND from the probe to the target Pico debug port.
- If you use an official Raspberry Pi Debug Probe, use the JST cable for this connection.
- If you want serial output, connect the probe UART TX/RX pins to the target Pico UART pins.

Do not mix up the SWD wiring and the UART wiring. The two connections use different pins.

## 2. Install OpenOCD

Run this command.

```
brew install openocd
```

Verify that the version supports `rp2040.cfg` and CMSIS-DAP.

## 3. Test the Connection

Run this command.

```
openocd -f interface/cmsis-dap.cfg -f target/rp2040.cfg
```

If the command connects without errors, the probe communicates with the target board.

If the command fails, check the debug probe firmware. See Section 5.

## 4. Flash the Target Board

Run this command from the build directory.

```
openocd -f interface/cmsis-dap.cfg -f target/rp2040.cfg -c "program build/your_project.uf2 verify reset exit"
```

Replace `your_project.uf2` with the actual file name.

### Alternative: VS Code

1. Install the Raspberry Pi Pico extension in VS Code.
2. Use the extension to drive OpenOCD.
3. Use the Flash button to flash the board.
4. Use the extension to set breakpoints and step through code.

## 5. Troubleshooting

If OpenOCD cannot see the probe, or the connection hangs, update the debug probe firmware.

The probe firmware update is a UF2 flash, the same method as for any other Pico target.

| Problem | Cause | Fix |
|---|---|---|
| Board does not mount | Wrong cable | Use a cable that supports data, not only power |
| Build fails with missing headers | Submodules not initialized | Run `git submodule update --init` in the SDK directory |
| CMake cannot find the SDK | Environment variable not set | Set `PICO_SDK_PATH` in `~/.zshrc` and run `source ~/.zshrc` |
| OpenOCD cannot connect | Old probe firmware | Update the probe firmware |
| SWD connection fails | Wrong wiring | Check for a mixup between SWD and UART pins |

## 6. Confirm Your Probe Type

Two probe types exist.

- The official Raspberry Pi Debug Probe.
- A second Pico flashed with debugprobe firmware.

The wiring and the firmware update steps differ between the two types.
