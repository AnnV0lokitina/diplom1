package external

// Error Keep error and its description.
type Error struct {
	Label string
	Err   error
}

func NewError(orig error, label string) *Error {
	return &Error{
		Label: label,
		Err:   orig,
	}
}

// Possible error descriptions.
const (
	ErrorTypeConnection         = "CONNECTION"
	ErrorTypeRegister           = "REGISTER"
	ErrorTypeLogin              = "LOGIN"
	ErrorTypeSaveToStorage      = "SAVE TO STORAGE"
	ErrorTypeRestoreFromStorage = "RESTORE FROM STORAGE"
)

func (e *Error) description() string {
	switch e.Label {
	case ErrorTypeConnection:
		return "Cannot connect to store"
	case ErrorTypeRegister:
		return "Registration failed"
	case ErrorTypeLogin:
		return "Login failed"
	case ErrorTypeSaveToStorage:
		return "Save to storage failed"
	case ErrorTypeRestoreFromStorage:
		return "Restore from storage failed"
	}
	return "Unknown error"
}

func (e *Error) Error() string {
	return "External error: " + e.description()
}
