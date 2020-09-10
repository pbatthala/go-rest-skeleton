package exception

import "errors"

// ErrorTextNoRecordInsertedToRedis is an error representing there is no record inserted to redis.
var ErrorTextNoRecordInsertedToRedis = errors.New("api.msg.error.common.no_record_inserted_to_redis")

// ErrorTextInternalServerError is an error representing internal server error.
var ErrorTextInternalServerError = errors.New("api.msg.error.common.internal_server_error")

// ErrorTextAnErrorOccurred is an error representing an error occurred.
var ErrorTextAnErrorOccurred = errors.New("api.msg.error.common.an_error_occurred")

// ErrorTextUnauthorized is an error representing unauthorized request.
var ErrorTextUnauthorized = errors.New("api.msg.error.common.unauthorized")

// ErrorTextForbidden is an error representing forbidden request.
var ErrorTextForbidden = errors.New("api.msg.error.common.forbidden")

// ErrorTextBadRequest is an error representing bad request.
var ErrorTextBadRequest = errors.New("api.msg.error.common.bad_request")

// ErrorTextUnprocessableEntity is an error representing unprocessable entity.
var ErrorTextUnprocessableEntity = errors.New("api.msg.error.common.unprocessable_entity")

// ErrorTextNotFound is an error representing request not found.
var ErrorTextNotFound = errors.New("api.msg.error.common.not_found")

// ErrorTextFileTooLarge is an error representing that received file size too large.
var ErrorTextFileTooLarge = errors.New("api.msg.error.common.file_too_large")

// ErrorTextInvalidPrivateKey is an error representing invalid private key.
var ErrorTextInvalidPrivateKey = errors.New("api.msg.error.common.invalid_private_key")

// ErrorTextInvalidPublicKey is an error representing invalid public key.
var ErrorTextInvalidPublicKey = errors.New("api.msg.error.common.invalid_public_key")

// ErrorTextRefreshTokenIsExpired is an error representing refresh token is expired.
var ErrorTextRefreshTokenIsExpired = errors.New("api.msg.error.common.refresh_token_expired")

// ErrorTextRoleNotFound is an error representing role not found in database.
var ErrorTextRoleNotFound = errors.New("api.msg.error.role.not_found")

// ErrorTextRoleInvalidUUID is an error representing UUID not found in database.
var ErrorTextRoleInvalidUUID = errors.New("api.msg.error.role.invalid_uuid")

// ErrorTextStorageCategoryNotFound is an error representing storage_category not found in database.
var ErrorTextStorageCategoryNotFound = errors.New("storage_category.not_found")

// ErrorTextStorageFileNotFound is an error representing storage_file not found in database.
var ErrorTextStorageFileNotFound = errors.New("storage_file.not_found")

// ErrorTextUserNotFound is an error representing user not found in database.
var ErrorTextUserNotFound = errors.New("api.msg.error.user.not_found")

// ErrorTextUserInvalidUUID is an error representing UUID not found in database.
var ErrorTextUserInvalidUUID = errors.New("api.msg.error.user.invalid_uuid")

// ErrorTextUserInvalidPassword is an error representing hashed password not match with stored in database.
var ErrorTextUserInvalidPassword = errors.New("api.msg.error.user.invalid_password")

// ErrorTextUserInvalidUsernameAndPassword is an error representing hashed password not match with stored in database.
var ErrorTextUserInvalidUsernameAndPassword = errors.New("api.msg.error.user.invalid_email_and_password")

// ErrorTextUserEmailNotRegistered is an error representing email already is not exists in database.
var ErrorTextUserEmailNotRegistered = errors.New("api.msg.error.user.email_not_registered")

// ErrorTextUserEmailAlreadyTaken is an error representing email already exists in database.
var ErrorTextUserEmailAlreadyTaken = errors.New("api.msg.error.user.email_already_taken")

// ErrorTextUserPreferenceInvalidUUID is an error representing UUID not found in database.
var ErrorTextUserPreferenceInvalidUUID = errors.New("user_preference.invalid_uuid")

// ErrorTextUploadCannotOpenFile is an error representing uploaded file can not opened by system.
var ErrorTextUploadCannotOpenFile = errors.New("api.msg.error.common.cannot_open_file")

// ErrorTextUploadInvalidSize is an error representing uploaded file size is greater than allowed maximum size.
var ErrorTextUploadInvalidSize = errors.New("api.msg.error.common.upload_invalid_size")

// ErrorTextUploadInvalidFileType is an error representing uploaded file has invalid file type.
var ErrorTextUploadInvalidFileType = errors.New("api.msg.error.common.upload_invalid_file_type")

// ErrorTextPerPage is an error representing request per page over the limit.
var ErrorTextPerPage = errors.New("api.msg.error.common.per_page")

// ErrorCodeIFAUGA001 is an error represent represent given authorization via request headers
// is not valid.
// Or does not send authorization.
const ErrorCodeIFAUGA001 = "IFAUGA001"

// ErrorCodeIFAUGA002 is an error represent request with not supported authentication type.
const ErrorCodeIFAUGA002 = "IFAUGA002"

// ErrorCodeIFAUGA003 is an error represent decoded Basic Auth does not content
// pair of username and password.
const ErrorCodeIFAUGA003 = "IFAUGA003"

// ErrorCodeIFAUGA004 is an error represent username and password from given
// Basic Auth not registered in the DB.
const ErrorCodeIFAUGA004 = "IFAUGA004"

// ErrorCodeIFAUGA005 is an error represent invalid JWT Token.
const ErrorCodeIFAUGA005 = "IFAUGA005"

// ErrorCodeITMIPO001 is an error represent UUID on context is not exists.
const ErrorCodeITMIPO001 = "ITMIPO001"

// ErrorCodeITMIPO002 is an error represent authenticated user perfom unauthorized request.
// User does not have permission to perform this request.
const ErrorCodeITMIPO002 = "ITMIPO002"