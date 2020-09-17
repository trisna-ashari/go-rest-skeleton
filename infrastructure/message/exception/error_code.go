package exception

const (
	// ErrorCodeIFAUGA001 is an error represent represent given authorization via request headers
	// is not valid.
	// Or does not send authorization.
	ErrorCodeIFAUGA001 = "IFAUGA001"

	// ErrorCodeIFAUGA002 is an error represent request with not supported authentication type.
	ErrorCodeIFAUGA002 = "IFAUGA002"

	// ErrorCodeIFAUGA003 is an error represent decoded Basic Auth does not content
	// pair of username and password.
	ErrorCodeIFAUGA003 = "IFAUGA003"

	// ErrorCodeIFAUGA004 is an error represent username and password from given
	// Basic Auth not registered in the DB.
	ErrorCodeIFAUGA004 = "IFAUGA004"

	// ErrorCodeIFAUGA005 is an error represent invalid JWT Token.
	ErrorCodeIFAUGA005 = "IFAUGA005"

	// ErrorCodeITMIPO001 is an error represent UUID on context is not exists.
	ErrorCodeITMIPO001 = "ITMIPO001"

	// ErrorCodeITMIPO002 is an error represent authenticated user perfom unauthorized request.
	// User does not have permission to perform this request.
	ErrorCodeITMIPO002 = "ITMIPO002"
)
