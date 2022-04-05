package main

import "strings"

func ValidarIngresoTeclado(nombre string, apellido string, email string, userTickets uint, ticketsRestantes uint) (bool, bool, bool) {
	validarNombre := len(nombre) >= 2 && len(apellido) >= 2
	validarEmail := strings.Contains(email, "@")
	validarTickets := userTickets > 0 && userTickets <= ticketsRestantes
	return validarNombre, validarEmail, validarTickets
}
