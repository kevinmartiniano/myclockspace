package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	responseWriter := httptest.NewRecorder()

	Index(responseWriter, request)

	if responseWriter.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseWriter.Code)
	}

	expectedBody := `{"message":"Ok!"}`
	actualBody := responseWriter.Body.String()

	if actualBody != expectedBody {
		t.Errorf("Expected body %s, but got %s", expectedBody, actualBody)
	}
}
