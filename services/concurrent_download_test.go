package services

import "testing"

func TestDownload(t *testing.T) {
	type args struct {
		URL string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name : "Valid URL",
			args : args{
				URL : "https://upload.wikimedia.org/wikipedia/commons/3/3f/Fron_big.jpg",
			},
			wantErr : false,
		},
		{
			name : "Valid URL",
			args : args{
				URL : "https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG",
			},
			wantErr : false,
		},
		{
			name : "Invalid URL",
			args : args{
				URL : "https://imas.pexels.com/photos/414612/pexels-phot612.jpeg",
			},
			wantErr : true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Download(tt.args.URL)
			if (err != nil) && (tt.wantErr == false) {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			} else if (err == nil) && (tt.wantErr == true) {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}