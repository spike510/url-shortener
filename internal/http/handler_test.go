package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spike510/url-shortener/internal/generator"
	"github.com/spike510/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupRouter(h *Handler) *gin.Engine {
	r := gin.Default()
	r.POST("/api/shorten", h.Shorten)
	r.GET("/:code", h.Redirect)
	return r
}

func TestShortenHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	storage := storage.NewInMemoryStorage()
	generator := generator.NewCodeGenerator()
	h := NewHandler("http://localhost:8080", generator, storage)
	router := setupRouter(h)

	t.Run("valid request", func(t *testing.T) {
		reqBody := `{"url":"http://example.com"}`
		req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusCreated, w.Code)
		var resp shortenResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		require.NoError(t, err)
		assert.NotEmpty(t, resp.Code)
		assert.Equal(t, "http://localhost:8080/"+resp.Code, resp.ShortURL)

		savedURL, err := storage.Get(resp.Code)
		require.NoError(t, err)
		assert.Equal(t, "http://example.com", savedURL)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		reqBody := `{"url":}`
		req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("missing URL", func(t *testing.T) {
		reqBody := `{}`
		req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestRedirectHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	storage := storage.NewInMemoryStorage()
	generator := generator.NewCodeGenerator()
	h := NewHandler("http://localhost:8080", generator, storage)
	router := setupRouter(h)

	code := "testcode"
	originalURL := "http://example.com"
	err := storage.Save(code, originalURL)
	require.NoError(t, err)

	t.Run("valid code", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/"+code, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusFound, w.Code)
		assert.Equal(t, originalURL, w.Header().Get("Location"))
	})

	t.Run("invalid code", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/invalidcode", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestRedirectHandler_MissingCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	storage := storage.NewInMemoryStorage()
	generator := generator.NewCodeGenerator()
	h := NewHandler("http://localhost:8080", generator, storage)
	router := setupRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}
