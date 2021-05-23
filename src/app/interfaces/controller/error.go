package controller

type Error struct {
	Message string
}

func APIError(msg string) *Error {
	return &Error{
		Message: msg,
	}
}
