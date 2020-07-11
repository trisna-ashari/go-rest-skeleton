package exception

import "errors"

// ErrorTextNoRecordInsertedToRedis is an error representing there is no record inserted to redis.
var ErrorTextNoRecordInsertedToRedis = errors.New("api.msg.error.no_record_inserted_to_redis")

// ErrorTextInternalServerError is an error representing internal server error.
var ErrorTextInternalServerError = errors.New("api.msg.error.internal_server_error")

// ErrorTextAnErrorOccurred is an error representing an error occurred.
var ErrorTextAnErrorOccurred = errors.New("api.msg.error.an_error_occurred")

// ErrorTextUnauthorized is an error representing unauthorized request.
var ErrorTextUnauthorized = errors.New("api.msg.error.unauthorized")

// ErrorTextForbidden is an error representing forbidden request.
var ErrorTextForbidden = errors.New("api.msg.error.forbidden")

// ErrorTextBadRequest is an error representing bad request.
var ErrorTextBadRequest = errors.New("api.msg.error.bad_request")

// ErrorTextUnprocessableEntity is an error representing unprocessable entity.
var ErrorTextUnprocessableEntity = errors.New("api.msg.error.unprocessable_entity")

// ErrorTextNotFound is an error representing request not found.
var ErrorTextNotFound = errors.New("api.msg.error.not_found")

// ErrorTextFileTooLarge is an error representing that received file size too large.
var ErrorTextFileTooLarge = errors.New("api.msg.error.file_too_large")

// ErrorTextInvalidPrivateKey is an error representing invalid private key.
var ErrorTextInvalidPrivateKey = errors.New("api.msg.error.invalid_private_key")

// ErrorTextInvalidPublicKey is an error representing invalid public key.
var ErrorTextInvalidPublicKey = errors.New("api.msg.error.invalid_public_key")

// ErrorTextRefreshTokenIsExpired is an error representing refresh token is expired.
var ErrorTextRefreshTokenIsExpired = errors.New("api.msg.error.refresh_token_expired")

// ErrorTextRoleNotFound is an error representing role not found in database.
var ErrorTextRoleNotFound = errors.New("role.not_found")

// ErrorTextUserNotFound is an error representing user not found in database.
var ErrorTextUserNotFound = errors.New("user.not_found")
