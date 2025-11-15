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

func TestDeleteTodo(t *testing.T) {
	// Reset todos and add initial data for this test
	todos = []Todo{
		{ID: 1, Text: "Initial Todo 1", Completed: false},
		{ID: 2, Text: "Initial Todo 2", Completed: true},
	}
	router := setupRouter()

	// Test deleting an existing todo
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/todos/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verify todo is deleted
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/todos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var todoList []Todo
	err := json.Unmarshal(w.Body.Bytes(), &todoList)
	assert.NoError(t, err)
	assert.Len(t, todoList, 1)
	assert.Equal(t, int64(2), todoList[0].ID) // Only todo with ID 2 should remain

	// Test deleting a non-existent todo
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/api/todos/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
