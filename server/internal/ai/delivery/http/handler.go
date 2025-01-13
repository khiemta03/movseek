package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h handlerImpl) healthy(c *gin.Context) {
	req, err := http.NewRequest("GET", aiHealthyURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to communicate with AI service"})
		return
	}
	defer resp.Body.Close()

	aiResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read AI service response"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), aiResponse)
}

func (h handlerImpl) retriever(c *gin.Context) {
	req, err := h.processRetrieverRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	baseURL, err := url.Parse(aiRetrieverURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid base URL"})
		return
	}

	params := url.Values{}
	params.Add("llm_api_key", req.LLMAPIKey)
	params.Add("query", req.Query)
	params.Add("collection_name", req.CollectionName)
	params.Add("amount", strconv.Itoa(req.Amount))
	params.Add("threshold", strconv.FormatFloat(req.Threshold, 'f', -1, 64))
	baseURL.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to communicate with AI service"})
		return
	}
	defer resp.Body.Close()

	aiResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read AI service response"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), aiResponse)
}

func (h handlerImpl) navigate(c *gin.Context) {
	req, err := h.processNavigateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	baseURL := fmt.Sprintf("%s/?llm_api_key=%s&query=%s",
		aiNavigateURL,
		url.QueryEscape(req.LLMAPIKey),
		url.QueryEscape(req.Query))

	httpReq, err := http.NewRequest("POST", baseURL, strings.NewReader(""))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	httpReq.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to communicate with AI service"})
		return
	}
	defer resp.Body.Close()

	aiResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read AI service response"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), aiResponse)
}
