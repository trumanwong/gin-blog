package exceptions

type Exception struct {
	Message string
}

func (this *Exception) Error() string {
	return this.Message
}