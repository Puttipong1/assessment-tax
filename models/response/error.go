package response

type Error struct {
	HttpStatus int    `json:"-"`
	Message    string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}
