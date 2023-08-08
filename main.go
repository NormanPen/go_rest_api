package main

import (
	"fmt"

	"github.com/NormanPen/l298nhbridge"
)

func main() {
	fmt.Println("Hello, World!")

	// Initialisieren der Motoren und GPIO-Pins
	l298nhbridge.Init()

	// Motoren vorwärts bewegen
	l298nhbridge.SetMotors(0.5, 0.5) // Hier kannst du die Geschwindigkeiten einstellen
	l298nhbridge.SetMotors(0.5, 0.5)
	l298nhbridge.SetMotors(0.5, 0.5)
	l298nhbridge.SetMotors(0.5, 0.5)
	l298nhbridge.SetMotors(0.5, 0.5)
	l298nhbridge.SetMotors(0.5, 0.5)
	// Warte eine Weile, während die Motoren laufen
	// Hier kannst du deine eigene Logik einfügen
	// z.B.: time.Sleep(time.Second * 5)

	// Stoppe die Motoren
	l298nhbridge.StopMotors()

	// Beende die GPIO-Ressourcen
	l298nhbridge.Exit()
}
