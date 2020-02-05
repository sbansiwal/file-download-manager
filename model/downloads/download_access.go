package downloads

import (
	"fmt"
	"github.com/sunil-bansiwal/file_download_manager/utils/errors"
)

var FilesDB = make(map[string]*DownloadedFiles)

func (downloadedFile DownloadedFiles) Get() (*DownloadedFiles, *errors.RestErr) {
	result := FilesDB[downloadedFile.Id]

	if result == nil {
		restErr := errors.NewBadRequestError("Download ID doesn't exists")
		return nil, restErr
	}

	MappedUrl := result.Files

	cnt := 0

	for _, check := range MappedUrl {
		if check == "Unsuccessful" {
			cnt++
		}
	}

	if cnt == 0 {
		result.Status = "All files downloaded successfully"
	} else if cnt == 1 {
		result.Status = "1 file not downloaded"
	} else {
		result.Status = fmt.Sprintf("%d files not downloaded", cnt)
	}

	return result, nil
}

func (downloadedFile DownloadedFiles) Save() *errors.RestErr {
	check := FilesDB[downloadedFile.Id]

	if check != nil {
		restErr := errors.NewBadRequestError("Download ID already exists")
		return restErr
	}

	FilesDB[downloadedFile.Id] = &downloadedFile

	return nil
}