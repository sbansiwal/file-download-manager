package services

import (
	"reflect"
	"testing"
)

func TestSerialDownload(t *testing.T) {

	type args struct {
		urls []string
	}
	tests := []struct {
		name  string
		args  args
		isSuccess  bool
	}{
		// TODO: Add test cases.
		{
			name : "All files downloaded successfully",
			args : args{
				[]string{
					"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
					"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG",
				},
			},
			isSuccess : true,
		},
		{
			name : "1 or more unsuccessful downloads",
			args : args{
				urls : []string{
					"https://imas.pexels.com/photos/414612/pexels-phot612.jpeg",
					"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG",
				},
			},
			isSuccess : false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result := SerialDownload(tt.args.urls)
			if !reflect.DeepEqual(result, tt.isSuccess) {
				t.Errorf("SerialDownload() got = %v, want = %v", result, tt.isSuccess)
			}
		})
	}
}