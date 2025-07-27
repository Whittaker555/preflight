package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadResult accepts a JSON cost analysis result and saves it to disk.
func UploadResult(c *gin.Context) {
	var result map[string]interface{}
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to marshal JSON"})
		return
	}

	if err := os.MkdirAll("results", 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create results directory"})
		return
	}

	filename := "results/result-" + time.Now().Format("20060102150405") + ".json"
	if err := os.WriteFile(filename, data, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to save result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "saved", "path": filename})
}
