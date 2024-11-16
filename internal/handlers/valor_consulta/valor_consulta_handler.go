package valorconsulta

import (
	"dentistry-clinic/internal/application"
	valorconsulta "dentistry-clinic/internal/domain/valor_consulta"
	"dentistry-clinic/internal/infrastructure/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitHandlerValorConsulta(r *gin.Engine, db *mongo.Database) {

	repoValor := repository.NovoValorconsultaRepository(db)
	repoService := application.NovoValorConsultaService(repoValor)

	r.GET("/valor-consulta", func(c *gin.Context) {
		valorConsultas, err := repoService.ListarConsultas(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, valorConsultas)
	})

	r.POST("/valor-consulta", func(c *gin.Context) {
		var valorConsulta valorconsulta.Valorconsulta
		if err := c.ShouldBindJSON(&valorConsulta); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err := repoService.Criar(c, float64(valorConsulta.Valor))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Consulta salva com Sucesso!"})
	})
}
