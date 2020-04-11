package handlers

// TOTPRegistrationAction is the string representation of the action for which the token has been produced.
const TOTPRegistrationAction = "RegisterTOTPDevice"

// U2FRegistrationAction is the string representation of the action for which the token has been produced.
const U2FRegistrationAction = "RegisterU2FDevice"

// ResetPasswordAction is the string representation of the action for which the token has been produced.
const ResetPasswordAction = "ResetPassword"

const authPrefix = "Basic "

const AuthorizationHeader = "Proxy-Authorization"
const remoteUserHeader = "Remote-User"
const remoteGroupsHeader = "Remote-Groups"

var protoHostSeparator = []byte("://")

const (
	// Forbidden means the user is forbidden the access to a resource
	Forbidden authorizationMatching = iota
	// NotAuthorized means the user can access the resource with more permissions.
	NotAuthorized authorizationMatching = iota
	// Authorized means the user is authorized given her current permissions.
	Authorized authorizationMatching = iota
)

const operationFailedMessage = "Operation failed."
const authenticationFailedMessage = "Authentication failed. Check your credentials."
const userBannedMessage = "Please retry in a few minutes."
const unableToRegisterOneTimePasswordMessage = "Unable to set up one-time passwords." //nolint:gosec
const unableToRegisterSecurityKeyMessage = "Unable to register your security key."
const unableToResetPasswordMessage = "Unable to reset your password."
const mfaValidationFailedMessage = "Authentication failed, please retry later."
