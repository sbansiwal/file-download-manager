package errors

var internalCode int = 4000

type RestErr struct {
	InternalCode int 	`json:"internal_code"`
	Message 	 string	`json:"message"`
}

func NewBadRequestError(message string) *RestErr {
	internalCode++
	return &RestErr{
		InternalCode :   internalCode,
		Message		 :	 message,
	}
}


