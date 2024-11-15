package dentista

import (
	"context"
	"dentistry-clinic/internal/application"
	"dentistry-clinic/internal/domain/dentista"
	"dentistry-clinic/internal/infrastructure/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitHandlerDentista(r *gin.Engine, db *mongo.Database) {

	dentistaRepo := repository.NovoDentistaRepository(db)
	dentistaService := application.NovoDentistaService(dentistaRepo)

	r.GET("/dentista", func(ctx *gin.Context) {
		dentistas, err := dentistaService.Lista(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"dentistas": dentistas})
	})

	r.POST("/dentista", func(c *gin.Context) {
		var dentista dentista.Dentista
		if err := c.BindJSON(&dentista); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		_, err := dentistaService.Criar(context.Background(), dentista.Nome, dentista.CRO)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"Message": "Dentista Criado com Sucesso!"})
	})
}
