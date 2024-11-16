package main

import (
	"dentistry-clinic/cmd/db"
	"dentistry-clinic/internal/handlers/auth"
	"dentistry-clinic/internal/handlers/dentista"
	handler "dentistry-clinic/internal/handlers/paciente"
	valorconsulta "dentistry-clinic/internal/handlers/valor_consulta"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := db.InitMongoDb()

	r := gin.Default()
	r.Use(cors.Default())
	auth.InitHanlderAuth(r)
	handler.InitHandlerPaciente(r, db)
	dentista.InitHandlerDentista(r, db)
	valorconsulta.InitHandlerValorConsulta(r, db)
	r.Run(":1001")
}
