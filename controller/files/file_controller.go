package files

import (
	"github.com/gin-gonic/gin"
	"github.com/sunil-bansiwal/file_download_manager/model/downloads"
	"net/http"
)

func GetDownloadedFiles(c *gin.Context)  {
	fileDB := downloads.FilesDB

	for id, _ := range fileDB {
		c.JSON(http.StatusOK, id)
	}

}
