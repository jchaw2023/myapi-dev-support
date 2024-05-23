package biz_err

type BzsErr struct {
	Msg string
}

func NewBzsErr(msg string) *BzsErr {
	return &BzsErr{msg}
}

func (e *BzsErr) Error() string {
	return e.Msg
}

type AuthErr struct {
	Msg string
}

func NewAuthErr(msg string) *AuthErr {
	return &AuthErr{msg}
}
func (e *AuthErr) Error() string {
	return e.Msg
}
