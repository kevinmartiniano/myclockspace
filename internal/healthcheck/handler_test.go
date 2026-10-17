package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	responseWriter := httptest.NewRecorder()

	NewHealthcheckHandler().Index(responseWriter, request)

	if responseWriter.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseWriter.Code)
	}

	expectedBody := `{"message":"Ok!"}`
	actualBody := strings.TrimSpace(responseWriter.Body.String())

	if actualBody != expectedBody {
		t.Errorf("Expected body %s, but got %q", expectedBody, actualBody)
	}
}
