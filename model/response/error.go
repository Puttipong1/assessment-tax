package response

type Error struct {
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}
