package downloads

import "time"

type DownloadedFiles struct {
	Id 			 string 			`json:"id"`
	StartTime 	 time.Time 			`json:"start_time"`
	EndTime 	 time.Time 			`json:"end_time"`
	Status 		 string				`json:"status"`
	DownloadType string				`json:"download_type"`
	Files 	  	 map[string]string	`json:"files"`

}

type NewDownloadFiles struct {
	Type   string `json:"type"`
	URLs []string `json:"urls"`
}




