package main

import (
	"alunos/controllers"
	"alunos/database"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine{
	rotas := gin.Default()
	return rotas
}

// func TestFalhador(t *testing.T){
// 	t.Fatalf("Teste falhou de propósito, não se preocupe")
// }

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T){
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/tester",nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockResposta := `{"API diz:":"E aí tester, tudo certin?"}`
	respostaBody,_ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, string(respostaBody))

	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	// }
}

func TestListandoTodosOsAlunosHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET","/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}