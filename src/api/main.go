package main


//Autores: Mauricio Clerici y Franco Nahuel Capurro
/*
	Utilizamos la libreria gobreaker(github.com/sony/gobreaker) de Sony,
y las modificaciones pueden verse en src/api/domain/myml/user_receivers.go.
Cabe aclarar que utilizamos el circuit breaker como parte del servicio y no
como un servicio aparte. Si se desea cambiar la configuracion, hay que modificar
el struct Setting dentro de la función init.
*/


import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myml/src/api/controllers/myml"
	"github.com/mercadolibre/myml/src/api/controllers/ping"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()

)



func main() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:id", myml.Myml)
	router.Run(port)

}
