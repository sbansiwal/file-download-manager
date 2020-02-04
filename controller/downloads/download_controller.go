package downloads

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/sunil-bansiwal/file_download_manager/model/downloads"
	"github.com/sunil-bansiwal/file_download_manager/services"
	"github.com/sunil-bansiwal/file_download_manager/utils/errors"
	"net/http"
)

func CheckDownloadStatus(c *gin.Context) {
	downloadID := string(c.Param("id"))

	downloadedFile, err := services.CheckDownloadStatus(downloadID)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, downloadedFile)
}

func NewDownload(c *gin.Context){

	var newDownloadFiles downloads.NewDownloadFiles

	err := c.ShouldBindJSON(&newDownloadFiles)

	if err != nil {
		restErr := errors.NewBadRequestError("invalid JSON body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	if newDownloadFiles.Type != "serial" && newDownloadFiles.Type != "concurrent" {
		restErr := errors.NewBadRequestError("unknown type of download")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	if newDownloadFiles.Type == "concurrent" {

		DownloadID := xid.New().String()
		c.JSON(http.StatusOK, DownloadID)
		_, downloadErr := services.NewDownload(newDownloadFiles, string(DownloadID))

		if downloadErr != nil {
			c.JSON(http.StatusBadRequest, downloadErr)
		}

	} else{

		result, downloadErr := services.NewDownload(newDownloadFiles, "")

		if downloadErr != nil {
			c.JSON(http.StatusBadRequest, downloadErr)
			return
		}

		c.JSON(http.StatusOK, result.Id)
	}

}
