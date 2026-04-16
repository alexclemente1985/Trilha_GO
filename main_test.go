package main

import (
	"alunos/controllers"
	"alunos/database"
	"alunos/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //para um resultado mais enxuto
	rotas := gin.Default()
	return rotas
}

// func TestFalhador(t *testing.T){
// 	t.Fatalf("Teste falhou de propósito, não se preocupe")
// }

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/tester", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockResposta := `{"API diz:":"E aí tester, tudo certin?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, string(respostaBody))

	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	// }
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()         //cria um aluno para o teste
	defer DeletaAlunoMock() //apaga o aluno de teste no final do processo

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Tester",
		CPF:  "12345678901",
		RG:   "123456789",
	}

	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock) //pega o corpo a resposta em JSON e joga na variável

	assert.Equal(t, alunoMock.Nome, "Tester")
	assert.Equal(t, alunoMock.CPF, "12345678901")
	assert.Equal(t, alunoMock.RG, "123456789")

}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{
		Nome: "Tester",
		CPF:  "99945678901",
		RG:   "123456700",
	}

	alunoJSON, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(alunoJSON))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno

	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	//assert.Equal(t, http.StatusOK, resposta.Code)

	assert.Equal(t, alunoMockAtualizado.Nome, "Tester")
	assert.Equal(t, alunoMockAtualizado.CPF, "99945678901")
	assert.Equal(t, alunoMockAtualizado.RG, "123456700")
	assert.NotEqual(t, alunoMockAtualizado.CPF, "12345678901")
	assert.NotEqual(t, alunoMockAtualizado.RG, "123456789")

}
