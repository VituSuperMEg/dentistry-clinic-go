package handler

import (
	"context"
	"dentistry-clinic/internal/application"
	"dentistry-clinic/internal/infrastructure/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitHandlerPaciente(r *gin.Engine, db *mongo.Database) {

	pacienteRepo := repository.NovoPacienteRepository(db)
	pacienteService := application.NovoPacienteService(pacienteRepo)

	r.GET("/pacientes", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pacientes.html", nil)
	})

	r.POST("/pacientes", func(c *gin.Context) {
		nome := c.PostForm("nome")
		cpf := c.PostForm("cpf")
		telefone := c.PostForm("telefone")

		_, err := pacienteService.RegistrarPaciente(context.Background(), nome, cpf, telefone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		pacientes, err := pacienteService.ListarPacientes(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "lista_pacientes.html", gin.H{"pacientes": pacientes})
	})

	r.GET("/pacientes/lista", func(c *gin.Context) {
		pacientes, err := pacienteService.ListarPacientes(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "lista_pacientes.html", gin.H{"pacientes": pacientes})
	})
}
