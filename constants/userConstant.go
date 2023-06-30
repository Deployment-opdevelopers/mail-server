package constants

// request related
const (
	HTTP_METHOD_GET  = "GET"
	HTTP_METHOD_POST = "POST"
)

// api status
const (
	API_FAILED_STATUS  = "Fail"
	API_SUCCESS_STATUS = "Success"
)

// error messages
const (
	INVALID_PASSWORD                       = "incorrect password"
	INVALID_USER_ID                        = "invalid userId"
	INVALID_REQUEST                        = "invalid json request body"
	INVALID_MAIL_ID                        = "invalid email address provided"
	INVALID_TOKEN                          = "invalid token"
	ERROR_BAD_REQUEST                      = "bad request"
	ERROR_IN_HASHING_PASSWORD              = "error while hashing password"
	ERROR_IN_STORING_UNIQUE_USER           = "the user already exists"
	ERROR_AUTH                             = "error authenticating user request"
	WRONG_PROVIDER                         = "wrong provider"
	NOT_AUTHERIZED                         = "unautherized access"
	REFRESH_TOKEN_EXPIRED                  = "refresh token expired"
	ERROR_EXTRACTING_POST_METADATA         = "Error extracting token metadata"
	REFRESH_EXPIRED                        = "refresh expired"
	ERROR_IN_VERIFYING_CODE                = "Verification code provided is Invalid. Please look in your mail for the code"
	INVALID_REQUEST_FOR_LOGIN              = "Please Signup first in the platform"
	INVALID_REQUEST_NOT_ADMIN              = "you don't have permission to access this resource"
	ERROR_USER_NOT_FOUND                   = "this email is not registered with us"
	SERVER_ERROR                           = "server error"
	PASSWORD_MUST_BE_AT_LEAST_6_CHARACTERS = "password must be at least 6 characters"
	PASSWORD_TOO_SHORT                     = "password too short"
	ERROR_IN_UPDATING_PROFILE              = "error in updating profile"
	GET_IN_TOUCH_ALREADY_EXIST             = "you have already sent a request to get in touch"
	EMAIL_DOES_NOT_BELONG_TO_YOU           = "email does not belong to you"
)

// response messages
const (
	CREATE_USER_MESSAGE                = "User Created Successfully"
	LOGIN_USER_MESSAGE                 = "User Logged In Successfully"
	LOGOUT_MESSAGE                     = "User Logged Out Successfully"
	VERIFY_MAIL                        = "Please verify your email account using the confirmation code send to your mail"
	ERROR_IN_SENDING_MAIL              = "error in sending mail"
	ERROR_IN_GETTING_VERIFICATION_DATA = "Unable to verify mail. Please try again later"
	CODE_EXPIRED                       = "Confirmation code has expired. Please try generating a new code"
	USER_VERIFICATION_SUCCESS          = "user mail verification succeeded"
	SEND_EMAIL_SUCCESS                 = "email sent successfully"
	ADMIN_MESSAGE                      = "successfully made user admin"
	USER_PROFILE_UPDATE_SUCCESS        = "user profile updated successfully"
	USER_FEED_BACK_SUCCESS             = "user feedback submitted successfully"
	USER_GET_IN_TOUCH_SENT_SUCCESS     = "user get in touch received successfully"
	USER_SET_NEW_PASSWORD_SUCCESS      = "user set new password successfully"
)

// code alphabets
const (
	CODE_ALPHABET_SHORT = "abcdef0123456789"
	CODE_ALPHABET_LONG  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// required request body fields
var (
	SEND_EMAIL_REQUIRED_FIELDS = []string{"userName", "email", "type", "otp" ,"company", "adminId", "password", "subject", "lenderName", "lenderEmail"}
	SEND_EMAIL_OPTIONAL_FIELDS = []string{}
)

// info messages
const (
	INFO_CACHE_DISABLED = "cache disabled"
)

// ContentType
const (
	CONTENT_TYPE_JSON        = "application/json"
	CONTENT_TYPE_XML_ENCODED = "application/x-www-form-urlencoded"
)

const (
	DISCORD_URL = "https://discord.com/api/webhooks/"
	BLOG_URL    = "localhost:8080/blog/"
)
