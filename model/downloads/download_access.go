package downloads

import "github.com/sunil-bansiwal/file_download_manager/utils/errors"

var FilesDB = make(map[string]*DownloadedFiles)

func (downloadedFile DownloadedFiles) Get() (*DownloadedFiles, *errors.RestErr) {
	result := FilesDB[downloadedFile.Id]

	if result == nil {
		restErr := errors.NewBadRequestError("Download ID doesn't exists")
		return nil, restErr
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