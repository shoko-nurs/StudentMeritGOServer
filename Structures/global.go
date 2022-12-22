package Structures





type RESPONSE struct{
	Data any
	Ep map[string]string
	Message string
}


type ERROR struct{
	Err string
}

func (C ERROR) Error() string{
	return C.Err
}

