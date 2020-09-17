package storage_test

import (
	"bytes"
	"fmt"
	"go-rest-skeleton/infrastructure/storage"
	"go-rest-skeleton/pkg/json_formatter"
	"go-rest-skeleton/pkg/util"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestUploadFile_Success(t *testing.T) {
	conf := InitConfig()

	dbConn, errDBConn := DBConnSetup(conf.DBTestConfig)
	if errDBConn != nil {
		t.Fatalf("want non error, got %#v", errDBConn)
	}

	dbService, errDBService := DBServiceSetup(conf.DBTestConfig)
	if errDBService != nil {
		t.Fatalf("want non error, got %#v", errDBService)
	}

	_, errSeeds := SeedStorageCategories(dbConn)
	if errSeeds != nil {
		t.Fatalf("want non error, got %#v", errSeeds)
	}

	storageService, errStorageService := StorageServiceSetup(conf.MinioConfig, dbService.DB)
	if errStorageService != nil {
		t.Fatalf("want non error, got %#v", errStorageService)
	}

	filePath := fmt.Sprintf("%s/tests/file/image.jpeg", util.RootDir())
	fileOpen, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Cannot open file: %s\n", err)
	}

	requestBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(requestBody)
	fileWriter, errCreateFormFile := multipartWriter.CreateFormFile("file", "image.jpeg")
	if errCreateFormFile != nil {
		t.Fatalf("want non error, got %#v", errCreateFormFile)
	}
	_, err = io.Copy(fileWriter, fileOpen)
	if err != nil {
		t.Errorf("Cannot copy file: %s\n", err)
	}
	_ = fileOpen.Close()
	_ = multipartWriter.Close()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/upload", func(c *gin.Context) {
		file, errGetFile := c.FormFile("file")
		if errStorageService != nil {
			t.Fatalf("want non error, got %#v", errGetFile)
		}

		fileName, _, errException, _ := storageService.Storage.UploadFile(file, "avatar")
		if errException != nil {
			t.Fatalf("want non error, got %#v", errException)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    fileName,
			"message": "OK",
		})
	})

	c.Request, _ = http.NewRequest("POST", "/upload", requestBody)
	c.Request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	r.ServeHTTP(w, c.Request)
	response := json_formatter.ResponseDecoder(w.Body)

	assert.NotNil(t, response["data"])
}

func TestGetFile_Success(t *testing.T) {
	conf := InitConfig()

	dbConn, errDBConn := DBConnSetup(conf.DBTestConfig)
	if errDBConn != nil {
		t.Fatalf("want non error, got %#v", errDBConn)
	}

	dbService, errDBService := DBServiceSetup(conf.DBTestConfig)
	if errDBService != nil {
		t.Fatalf("want non error, got %#v", errDBService)
	}

	_, errSeeds := SeedStorageCategories(dbConn)
	if errSeeds != nil {
		t.Fatalf("want non error, got %#v", errSeeds)
	}

	storageService, errStorageService := StorageServiceSetup(conf.MinioConfig, dbService.DB)
	if errStorageService != nil {
		t.Fatalf("want non error, got %#v", errStorageService)
	}

	filePath := fmt.Sprintf("%s/tests/file/image.jpeg", util.RootDir())
	fileOpen, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Cannot open file: %s\n", err)
	}

	requestBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(requestBody)
	fileWriter, errCreateFormFile := multipartWriter.CreateFormFile("file", "image.jpeg")
	if errCreateFormFile != nil {
		t.Fatalf("want non error, got %#v", errCreateFormFile)
	}
	_, err = io.Copy(fileWriter, fileOpen)
	if err != nil {
		t.Errorf("Cannot copy file: %s\n", err)
	}
	_ = fileOpen.Close()
	_ = multipartWriter.Close()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/upload", func(c *gin.Context) {
		file, errGetFile := c.FormFile("file")
		if errStorageService != nil {
			t.Fatalf("want non error, got %#v", errGetFile)
		}

		fileName, _, errException, _ := storageService.Storage.UploadFile(file, "avatar")
		if errException != nil {
			t.Fatalf("want non error, got %#v", errException)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    fileName,
			"message": "OK",
		})
	})

	c.Request, _ = http.NewRequest("POST", "/upload", requestBody)
	c.Request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	r.ServeHTTP(w, c.Request)
	response := json_formatter.ResponseDecoder(w.Body)

	fileUUID := response["data"].(string)
	url, errGetURL := storageService.Storage.GetFile(fileUUID)
	if errGetURL != nil {
		t.Fatalf("want non error, got %#v", errGetURL)
	}
	assert.NotNil(t, url)
}

func TestFormatFileName(t *testing.T) {
	fileName := "test-filename.jpeg"
	formattedFileName := storage.FormatFileName(fileName)
	assert.NotEqualValues(t, fileName, formattedFileName.String())
}
