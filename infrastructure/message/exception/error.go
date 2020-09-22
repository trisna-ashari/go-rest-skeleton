package exception

import (
	"errors"
)

// Common errors.
var (
	// ErrorTextNoRecordInsertedToRedis is an error representing there is no record inserted to redis.
	ErrorTextNoRecordInsertedToRedis = errors.New("api.msg.error.common.no_record_inserted_to_redis")

	// ErrorTextInternalServerError is an error representing internal server error.
	ErrorTextInternalServerError = errors.New("api.msg.error.common.internal_server_error")

	// ErrorTextAnErrorOccurred is an error representing an error occurred.
	ErrorTextAnErrorOccurred = errors.New("api.msg.error.common.an_error_occurred")

	// ErrorTextUnauthorized is an error representing unauthorized request.
	ErrorTextUnauthorized = errors.New("api.msg.error.common.unauthorized")

	// ErrorTextForbidden is an error representing forbidden request.
	ErrorTextForbidden = errors.New("api.msg.error.common.forbidden")

	// ErrorTextBadRequest is an error representing bad request.
	ErrorTextBadRequest = errors.New("api.msg.error.common.bad_request")

	// ErrorTextUnprocessableEntity is an error representing unprocessable entity.
	ErrorTextUnprocessableEntity = errors.New("api.msg.error.common.unprocessable_entity")

	// ErrorTextNotFound is an error representing request not found.
	ErrorTextNotFound = errors.New("api.msg.error.common.not_found")

	// ErrorTextFileTooLarge is an error representing that received file size too large.
	ErrorTextFileTooLarge = errors.New("api.msg.error.common.file_too_large")

	// ErrorTextInvalidPrivateKey is an error representing invalid private key.
	ErrorTextInvalidPrivateKey = errors.New("api.msg.error.common.invalid_private_key")

	// ErrorTextInvalidPublicKey is an error representing invalid public key.
	ErrorTextInvalidPublicKey = errors.New("api.msg.error.common.invalid_public_key")

	// ErrorTextRefreshTokenIsExpired is an error representing refresh token is expired.
	ErrorTextRefreshTokenIsExpired = errors.New("api.msg.error.common.refresh_token_expired")

	// ErrorTextPerPage is an error representing request per page over the limit.
	ErrorTextPerPage = errors.New("api.msg.error.common.per_page")
)

// Errors for role.
var (
	// ErrorTextRoleNotFound is an error representing role not found in database.
	ErrorTextRoleNotFound = errors.New("api.msg.error.role.not_found")

	// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
	ErrorTextRoleInvalidUUID = errors.New("api.msg.error.role.invalid_uuid")
)

// Errors for user.
var (
	// ErrorTextUserNotFound is an error representing user not found in database.
	ErrorTextUserNotFound = errors.New("api.msg.error.user.not_found")

	// ErrorTextUserInvalidUUID is an error representing UUID not found in database.
	ErrorTextUserInvalidUUID = errors.New("api.msg.error.user.invalid_uuid")

	// ErrorTextUserInvalidPassword is an error representing hashed password not match with stored in database.
	ErrorTextUserInvalidPassword = errors.New("api.msg.error.user.invalid_password")

	// ErrorTextUserInvalidUsernameAndPassword is an error representing hashed password not match with stored in database.
	ErrorTextUserInvalidUsernameAndPassword = errors.New("api.msg.error.user.invalid_email_and_password")

	// ErrorTextUserEmailNotRegistered is an error representing email already is not exists in database.
	ErrorTextUserEmailNotRegistered = errors.New("api.msg.error.user.email_not_registered")

	// ErrorTextUserEmailAlreadyTaken is an error representing email already exists in database.
	ErrorTextUserEmailAlreadyTaken = errors.New("api.msg.error.user.email_already_taken")

	// ErrorTextUserPreferenceInvalidUUID is an error representing UUID not found in database.
	ErrorTextUserPreferenceInvalidUUID = errors.New("api.msg.error.user.preference.invalid_uuid")

	// ErrorTextUserForgotPasswordTokenNotFound is an error representing Token not found in database.
	ErrorTextUserForgotPasswordTokenNotFound = errors.New("api.msg.error.user.forgot_password.token_not_found")
)

// Errors for storage.
var (
	// ErrorTextStorageCategoryNotFound is an error representing storage_category not found in database.
	ErrorTextStorageCategoryNotFound = errors.New("api.msg.error.storage.category.not_found")

	// ErrorTextStorageFileNotFound is an error representing storage_file not found in database.
	ErrorTextStorageFileNotFound = errors.New("api.msg.error.storage.file.not_found")

	// ErrorTextStorageUploadCannotOpenFile is an error representing uploaded file can not opened by system.
	ErrorTextStorageUploadCannotOpenFile = errors.New("api.msg.error.storage.file.cannot_open_file")

	// ErrorTextStorageUploadInvalidSize is an error representing uploaded file size is greater than allowed maximum size.
	ErrorTextStorageUploadInvalidSize = errors.New("api.msg.error.storage.file.invalid_file_size")

	// ErrorTextStorageUploadInvalidFileType is an error representing uploaded file has invalid file type.
	ErrorTextStorageUploadInvalidFileType = errors.New("api.msg.error.storage.file.invalid_file_type")
)
