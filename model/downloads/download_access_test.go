package downloads

import (
	"testing"
	"time"
)

func TestDownloadedFiles_Get(t *testing.T) {

	tests := []struct {
		name   string
		id 	   string
		want   bool
	}{
		// TODO: Add test cases.
		{
			name : "Download ID doesn't exists",
			id 	 : "123",
			want : false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloadedFile := DownloadedFiles{
				Id:           tt.id,
				StartTime:    time.Now(),
				EndTime:      time.Now(),
				Status:       "",
				DownloadType: "",
				Files:        nil,
			}
			got, _ := downloadedFile.Get()
			if (got != nil) && (tt.want == false) {
				t.Errorf("Get() got1 = %v, want %v", got, tt.want)
			} else if (got == nil) && (tt.want == true) {
				t.Errorf("Get() got1 = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestDownloadedFiles_Save(t *testing.T) {
	tests := []struct {
		name   string
		id 	   string
		want   bool
	}{
		// TODO: Add test cases.
		{
			name : "New download ID",
			id : "1234",
			want : true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloadedFile := DownloadedFiles{
				Id:           tt.id,
				StartTime:    time.Now(),
				EndTime:      time.Now(),
				Status:       "",
				DownloadType: "",
				Files:        nil,
			}

			got := downloadedFile.Save()

			if (got != nil) && (tt.want == true) {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			} else if (got == nil) && (tt.want == false) {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}