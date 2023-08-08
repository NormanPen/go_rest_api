package main

import (
	"fmt"
	"github.com/NormanPen/l298nhbridge"
)

func main() {
	fmt.Println("Hello, Gopher!")

	// Initialisieren der Motoren und GPIO-Pins
	l298nhbridge.Init()

	// Motoren vorwärts bewegen
	l298nhbridge.SetMotors(0.5, 0.5) // Hier kannst du die Geschwindigkeiten einstellen
	// ...

	// Stoppe die Motoren
	l298nhbridge.StopMotors()

	// Beende die GPIO-Ressourcen
	l298nhbridge.Exit()
}
