package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spike510/url-shortener/internal/generator"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

type Handler struct {
	baseUrl   string
	generator *generator.CodeGenerator
}

func NewHandler(baseUrl string, generator *generator.CodeGenerator) *Handler {
	if strings.HasSuffix(baseUrl, "/") {
		baseUrl = strings.TrimRight(baseUrl, "/")
	}
	return &Handler{baseUrl: baseUrl}
}

func (h *Handler) Shorten(c *gin.Context) {

	var req shortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
		return
	}

	// TODO: verify if code is unique in storage
	code, err := h.generator.GenerateCode(6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate code"})
		return
	}

	res := shortenResponse{Code: code, ShortURL: h.baseUrl + "/" + code}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) Redirect(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}

	// TODO: retrieving url for code from storage
	orig := "http://onet.pl"

	// Basic safety: ensure URL has scheme
	if !strings.HasPrefix(orig, "http://") && !strings.HasPrefix(orig, "https://") {
		orig = "http://" + orig
	}
	c.Redirect(http.StatusFound, orig)
}
