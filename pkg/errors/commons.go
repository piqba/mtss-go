package errors

var (
	// errors configs
	ErrReadConfig      = NewError("config: error to load yaml file")
	ErrUnmarshalConfig = NewError("config: error to unmarsahl yaml file")
)
