package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	responseWriter := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseWriter)

	context.Request, _ = http.NewRequest(http.MethodGet, "/healthcheck", nil)

	Index(context)

	if responseWriter.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseWriter.Code)
	}

	expectedBody := `{"message":"Ok!"}`
	actualBody := responseWriter.Body.String()

	if actualBody != expectedBody {
		t.Errorf("Expected body %s, but got %s", expectedBody, actualBody)
	}
}
