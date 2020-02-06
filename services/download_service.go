package services

import (
	"github.com/rs/xid"
	"github.com/sunil-bansiwal/file_download_manager/model/downloads"
	"github.com/sunil-bansiwal/file_download_manager/utils/errors"
	"time"
)

func CheckDownloadStatus(downloadID string) (*downloads.DownloadedFiles, *errors.RestErr){
	downloadedFile := &downloads.DownloadedFiles{ Id : downloadID }

	downloadedFile, err := downloadedFile.Get()

	if err != nil {
		return nil, err
	}

	return downloadedFile, nil
}

func NewDownload(newDownloadFile downloads.NewDownloadFiles, DownloadID string) (*downloads.DownloadedFiles, *errors.RestErr){

	if newDownloadFile.Type == "serial" {

		startTime := time.Now()
		MappedUrl, _ := SerialDownload(newDownloadFile.URLs)
		endTime := time.Now()

		DownloadID = xid.New().String()

		var status string
		status = ""

		DownloadedFile := downloads.DownloadedFiles{
			Id			:    string(DownloadID),
			StartTime 	:    startTime,
			EndTime		:    endTime,
			Status		:    status,
			DownloadType: 	 "serial",
			Files		:    *MappedUrl,
		}

		err := DownloadedFile.Save()

		if err != nil {
			return nil, err
		} else {
			return &DownloadedFile, nil
		}
	} else{

		startTime := time.Now()
		MappedUrl, _:= ConcurrentDownload(newDownloadFile.URLs)
		endTime := time.Now()

		var status string

		DownloadedFile := downloads.DownloadedFiles{
			Id			:    DownloadID,
			StartTime	:    startTime,
			EndTime		:    endTime,
			Status		:    status,
			DownloadType: 	 "concurrent",
			Files		:    *MappedUrl,
		}

		err := DownloadedFile.Save()

		if err != nil {
			return nil,err
		} else{
			return &DownloadedFile, nil
		}
	}
}
