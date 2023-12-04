// routes.go contém todas as rotas da API.
package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vidacalura/AFEB/internal/models"
	"github.com/vidacalura/AFEB/internal/services"
	"github.com/vidacalura/AFEB/internal/utils"
)

func NewRouter() *gin.Engine {
	env := &models.Env{
		DB: utils.ConectarBD(),
	}

	models.E = env

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	jog := r.Group("/api/jogadores")
	{
		jog.GET("/ranking", services.MostrarRankingJogadores)
		jog.GET("/:codJog", services.MostrarJogador)
		jog.POST("/", services.CadastrarJogador)
		jog.PUT("/", services.EditarCadastroJogador)
		jog.DELETE("/:codJog", services.ExcluirJogador)
	}

	notc := r.Group("/api/noticias")
	{
		notc.GET("/", services.MostrarTodasNoticias)
		notc.GET("/feed", services.MostrarFeedNoticias)
		notc.GET("/:codNotc", services.MostrarNoticia)
		notc.POST("/", services.CriarNoticia)
		notc.PUT("/", services.EditarNoticia)
		notc.DELETE("/:codNotc", services.ExcluirNoticia)
	}

	torn := r.Group("/api/torneios")
	{
		torn.GET("/", services.MostrarTodosTorneios)
		torn.GET("/:codTorn", services.MostrarTorneio)
		torn.POST("/", services.CriarTorneio)
		torn.PUT("/", services.EditarTorneio)
		torn.DELETE("/:codTorn", services.ExcluirTorneio)
	}

	trof := r.Group("/api/trofeus")
	{
		trof.GET("/jogador/:codJog", services.MostrarTrofeusJogador)
		trof.GET("/:codTrof", services.MostrarDadosTrofeu)
		trof.POST("/", services.CriarTrofeu)
		trof.PUT("/", services.EditarTrofeu)
		trof.DELETE("/:codTrof", services.ExcluirTrofeu)
	}

	usu := r.Group("/api/usuarios")
	{
		usu.GET("/", services.MostrarTodosUsuarios)
		usu.GET("/:username", services.MostrarUsuario)
		usu.POST("/", services.CriarUsuario)
		usu.PUT("/", services.EditarUsuario)
		usu.DELETE("/:username", services.ExcluirUsuario)
	}

	return r
}
