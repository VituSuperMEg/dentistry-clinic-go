package handler

import (
	"context"
	"dentistry-clinic/internal/application"
	"dentistry-clinic/internal/domain/paciente"
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
		var paciente paciente.Paciente

		if err := c.BindJSON(&paciente); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		_, err := pacienteService.RegistrarPaciente(context.Background(), paciente.Nome, paciente.CPF, paciente.Telefone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Paciente Cadastrado com Sucesso!"})
	})

	r.DELETE("/pacientes/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID não informado"})
			return
		}
		err := pacienteService.Delete(context.Background(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Paciente excluído com Sucesso"})
	})

	r.GET("/pacientes/lista", func(c *gin.Context) {

		pacientes, err := pacienteService.ListarPacientes(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"pacientes": pacientes})
	})
}
