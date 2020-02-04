package app

import (
	"github.com/sunil-bansiwal/file_download_manager/controller/downloads"
	"github.com/sunil-bansiwal/file_download_manager/controller/files"
	"github.com/sunil-bansiwal/file_download_manager/controller/health"
)

func mapURL() {

	router.GET("/health", health.Health)

	router.GET("/downloads/:id", downloads.CheckDownloadStatus)

	router.GET("/files", files.GetDownloadedFiles)

	router.POST("/downloads", downloads.NewDownload)

}
