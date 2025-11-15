package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	// Reset todos for this test
	todos = []Todo{}
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/todos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[]`, w.Body.String())
}

func TestAddTodo(t *testing.T) {
	// Reset todos for this test
	todos = []Todo{}
	router := setupRouter()

	newTodo := Todo{Text: "Buy groceries", Completed: false}
	jsonTodo, _ := json.Marshal(newTodo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/todos", bytes.NewBuffer(jsonTodo))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdTodo Todo
	err := json.Unmarshal(w.Body.Bytes(), &createdTodo)
	assert.NoError(t, err)
	assert.NotZero(t, createdTodo.ID)
	assert.Equal(t, newTodo.Text, createdTodo.Text)
	assert.Equal(t, newTodo.Completed, createdTodo.Completed)

	// Verify that the todo was added to the list
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/todos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var todoList []Todo
	err = json.Unmarshal(w.Body.Bytes(), &todoList)
	assert.NoError(t, err)
	assert.Len(t, todoList, 1)
	assert.Equal(t, createdTodo.ID, todoList[0].ID)
	assert.Equal(t, createdTodo.Text, todoList[0].Text)
	assert.Equal(t, createdTodo.Completed, todoList[0].Completed)
}
