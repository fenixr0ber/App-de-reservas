package main

//App para ingreso y reserva de tickets. Comienzo de la aplicacion.
import (
	"fmt"
	"sync"
	"time"
)

var nombreApp = "App de reserva" //solo para variables, no sirve para constantes.
const totalTickets = 100

var ticketsRestantes uint = 100

// var bookings [100]string array, para darle límite al array
var reservas = make([]DatoUsuario, 0) //utilizo estructura, iniciando lista de mapas //([]map[string]string, 0)

type DatoUsuario struct {
	nombre          string
	apellido        string
	email           string
	numeroDeTickets uint
}

var wg = sync.WaitGroup{} //espera a la funcion que comienza con GO para que termine de procesar

func main() {

	saludoUsuarios()

	for {

		nombre, apellido, email, ticketsUsuario := escanearTeclado()

		validarNombre, validarEmail, validarTickets := ValidarIngresoTeclado(nombre, apellido, email, ticketsUsuario, ticketsRestantes)

		if validarNombre && validarEmail && validarTickets {

			reservaTicket(ticketsUsuario, nombre, apellido, email)

			wg.Add(1)
			go enviarTickets(ticketsUsuario, nombre, apellido, email) //go routine, concurrency (crea una linea paralela de
			//reproduccion de codigo para que la principal no espere), luego de
			//utilizarla la elimina solo.

			nombre := obtenerNombre()
			fmt.Printf("El nombre del usuario que reservó es: %v\n", nombre)

			if ticketsRestantes == 0 {
				fmt.Println("Ya no quedan tickets disponibles.")
				break
			}

		} else {
			if !validarNombre {
				fmt.Println("El nombre o apellido que ingreso es muy corto")
			}
			if !validarEmail {
				fmt.Println("El email ingresado no contiene '@'")
			}
			if !validarTickets {
				fmt.Println("El numero ingresado de tickets es inválido")
			}

		}
	}
	wg.Wait()
}

func saludoUsuarios() {
	fmt.Printf("Bienvenido a %v\n", nombreApp)
	fmt.Printf("Tenemos un total de %v tickets y %v todavia están disponibles.\n", totalTickets, ticketsRestantes)
	fmt.Println("Obtenga aqui sus Tickets")
}

func obtenerNombre() []string {
	nombres := []string{}
	for _, reserva := range reservas { // "_" variable definida, pero no utilizada

		nombres = append(nombres, reserva.nombre)
	}

	return nombres
}

func escanearTeclado() (string, string, string, uint) {
	var nombre string
	var apellido string
	var email string
	var ticketsUsuario uint

	fmt.Println("Ingrese su nombre: ")
	fmt.Scan(&nombre)

	fmt.Println("Ingrese su apellido: ")
	fmt.Scan(&apellido)

	fmt.Println("Ingrese su email: ")
	fmt.Scan(&email)

	fmt.Println("Ingrese numero de tickets a reservar: ")
	fmt.Scan(&ticketsUsuario)

	return nombre, apellido, email, ticketsUsuario
}

func reservaTicket(ticketsUsuario uint, nombre string, apellido string, email string) {
	ticketsRestantes = ticketsRestantes - ticketsUsuario

	//creando mapa de usuario
	var datoUsuario = DatoUsuario{
		nombre:          nombre,
		apellido:        apellido,
		email:           email,
		numeroDeTickets: ticketsUsuario,
	}

	reservas = append(reservas, datoUsuario)
	fmt.Printf("La lista de reservas es: %v\n", reservas)

	fmt.Printf("Gracias %v %v por reservar %v tickets. Recibirá confirmación en el email: %v\n", nombre, apellido, ticketsUsuario, email)
	fmt.Printf("%v tickets quedan disponibles para %v\n", ticketsRestantes, nombreApp)
}

func enviarTickets(ticketsUsuario uint, nombre string, apellido string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets para %v %v", ticketsUsuario, nombre, apellido)
	fmt.Println("···················")
	fmt.Printf("Enviando tickets:\n %v \nal email %v\n", ticket, email)
	fmt.Println("···················")
	wg.Done()
}
