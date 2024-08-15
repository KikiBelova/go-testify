package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=20&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code, "want status code %d, got %d", http.StatusOK, responseRecorder.Code)

	responseBody := responseRecorder.Body.String()
	gotList := strings.Split(responseBody, ",")

	assert.Len(t, gotList, totalCount)
}

func TestMainHandlerWhenRequestValid(t *testing.T) {
	totalCount := 3
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "want status code %d, got %d", http.StatusOK, responseRecorder.Code)

	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafes, totalCount)
}

func TestMainHandlerWhenCityInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=petersburg", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "want status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)

	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}
