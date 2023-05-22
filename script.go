package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	// Constantes de eventos de teclado
	KeyDownEvent  = 0x0000
	KeyUpEvent    = 0x0002
	ExtendedKey   = 0x0001
	KeyEvent      = 0x0000
	UnicodeEvent  = 0x0004
	ScancodeEvent = 0x0008
)

// Estructura de evento de teclado
type KeyboardEvent struct {
	Type         uint16
	Repeat       uint8
	ScanCode     uint8
	ExtendedKey  uint8
	Reserved     uint8
	VirtualKey   uint32
	VirtualScan  uint32
	Unicode      uint32
}

// Función para simular una pulsación de tecla
func keyPress(keyCode int) {
	// Obtener el identificador de la consola actual
	console, _ := syscall.Open("CONIN$", syscall.O_RDWR, 0)

	// Crear un evento de teclado
	event := KeyboardEvent{
		Type:        KeyEvent,
		ScanCode:    uint8(keyCode),
		VirtualKey:  uint32(keyCode),
		VirtualScan: uint32(keyCode),
	}

	// Enviar el evento de teclado a la consola actual
	syscall.Syscall(syscall.SYS_IOCTL, console, uintptr(syscall.TIOCSTI), uintptr(unsafe.Pointer(&event)))

	// Cerrar la consola
	syscall.Close(console)
}

func main() {
	fmt.Println("Presiona la tecla 'A' para simular una pulsación de tecla.")
	fmt.Println("Presiona la tecla 'Q' para salir.")

	// Leer la entrada del teclado en un bucle infinito
	for {
		var input string
		fmt.Scanln(&input)

		// Comprobar la tecla ingresada
		switch input {
		case "a", "A":
			// Simular la pulsación de la tecla 'A' (código ASCII: 65)
			keyPress(65)
		case "q", "Q":
			// Salir del programa si se presiona la tecla 'Q'
			fmt.Println("Saliendo...")
			os.Exit(0)
		default:
			fmt.Println("Tecla no válida.")
		}
	}
}
