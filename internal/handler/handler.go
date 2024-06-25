package handler

import (
	"crypto/subtle"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sinameshkini/webshare/internal/utils"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

type Handler struct {
}

var (
	readWriteDir string

	//h *Handler
)

func Init(e *echo.Echo, readOnly, readWrite string, port int, username, password string) (err error) {
	//h = &Handler{}

	info, err := os.Stat(readOnly)
	if err != nil {
		return
	}
	if !info.IsDir() {
		return fmt.Errorf("Error: %s is not a directory\n", readOnly)
	}

	// basic auth
	if username != "" && password != "" {
		e.Use(middleware.BasicAuth(func(usr, pass string, c echo.Context) (bool, error) {
			// Be careful to use constant time comparison to prevent timing attacks
			if subtle.ConstantTimeCompare([]byte(usr), []byte(username)) == 1 &&
				subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1 {
				return true, nil
			}
			return false, nil
		}))

		fmt.Printf("Access protected by Basic Auth (Username: %s\tPassword: %s)\n", username, password)
	}

	if readWrite != "" {
		info, err := os.Stat(readWrite)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return fmt.Errorf("Error: %s is not a directory\n", readWrite)
		}

		readWriteDir = readWrite

		e.GET("/w", uploadForm)
		e.POST("/upload", uploadFile)
		fmt.Printf("Enter this address in other host in local network for upload file: http://%s:%d/w \n", utils.GetIP(), port)
	}

	e.Static("/", readOnly)
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   readOnly,
		Browse: true,
	}))
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "status: running")
	})

	fmt.Printf("Enter this address in other host in local network: http://%s:%d \n", utils.GetIP(), port)

	return nil
}

// Handler
func uploadForm(c echo.Context) error {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Upload File</title>
		</head>
		<body>
			<h1>Upload File</h1>
			<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="file" />
				<input type="submit" value="Upload" />
			</form>
		</body>
		</html>`
	return c.HTML(http.StatusOK, html)
}

func uploadFile(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := createFile(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<p>File uploaded successfully!</p>")
}

func createFile(filename string) (*os.File, error) {
	// Create destination file making sure the path is writeable.
	filePath := fmt.Sprintf("%s/%s", readWriteDir, filename)
	logrus.Infof("creating %s file\n", filePath)
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return dst, nil
}
