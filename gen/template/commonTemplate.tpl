func As(err error) *errors.Error {
	if err == nil {
		return nil
	}
	var e *errors.Error
	if errors1.As(err, &e) {
		return e
	}
	return nil
}

func New(code int, reason, message string) *errors.Error {
	return errors.New(code, reason, message)
}
