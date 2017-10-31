package main

import (
	"io/ioutil"
	"os"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	engine.POST("/topdf", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		if data, err := toPdf(string(body)); err != nil {
			c.Error(err)
		} else {
			c.Data(200, "application/pdf", data)
		}
	})
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	engine.Run(":" + port)
}
func toPdf(page string) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	pdfg.Dpi.Set(72)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(page)))
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}
	return pdfg.Bytes(), nil
}
