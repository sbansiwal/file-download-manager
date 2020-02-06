package services

import (
	"github.com/sunil-bansiwal/file_download_manager/model/downloads"
	"testing"
	"time"
)

func TestConcurrentDownload(t *testing.T) {

	type args struct {
		urls []string
	}
	tests := []struct {
		name  string
		id string
		args  args
		isSuccess  bool
	}{
		// TODO: Add test cases.
		{
			name : "All files downloaded successfully",
			id : "1",
			args : args{
				[]string{
					"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
					"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG",
				},
			},
			isSuccess : true,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ConcurrentDownload(tt.args.urls)

			DownloadedFile := downloads.DownloadedFiles{
				Id			:    tt.id,
				StartTime	:    time.Now(),
				EndTime		:    time.Now(),
				Status		:    "",
				DownloadType: 	 "concurrent",
				Files		:    *got,
			}

			DownloadedFile.Save()
			result := downloads.FilesDB[DownloadedFile.Id]
			MappedURL := result.Files

			cnt := 0

			for _, check := range MappedURL {
				if check == "Unsuccessful" {
					cnt++
				}
			}

			if (cnt == 0) && (tt.isSuccess == false) {
				t.Errorf("ConcurrentDownload() got = Successful, want = Unsuccessful")
			} else if (cnt > 0) && (tt.isSuccess == true) {
				t.Errorf("ConcurrentDownload() got = Unsuccessful, want = Successful")

			}
		})
	}

}