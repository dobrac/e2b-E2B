// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

// Defines values for EnvironmentState.
const (
	Building EnvironmentState = "Building"
	Done     EnvironmentState = "Done"
	Failed   EnvironmentState = "Failed"
)

// Defines values for Template.
const (
	Ansys   Template = "Ansys"
	Bash    Template = "Bash"
	Go      Template = "Go"
	Nodejs  Template = "Nodejs"
	Python3 Template = "Python3"
	Rust    Template = "Rust"
)

// EnvironmentState defines model for EnvironmentState.
type EnvironmentState string

// EnvironmentStateUpdate defines model for EnvironmentStateUpdate.
type EnvironmentStateUpdate struct {
	State EnvironmentState `json:"state"`
}

// Error defines model for Error.
type Error struct {
	// Error code
	Code int32 `json:"code"`

	// Error
	Message string `json:"message"`
}

// NewEnvironment defines model for NewEnvironment.
type NewEnvironment struct {
	Deps     []string `json:"deps"`
	Template Template `json:"template"`
}

// NewSession defines model for NewSession.
type NewSession struct {
	// Identifier of a code snippet which which is the environment associated
	CodeSnippetID string `json:"codeSnippetID"`

	// Option determining if the session is a shared persistent edit session
	EditEnabled *bool `json:"editEnabled,omitempty"`
}

// Session defines model for Session.
type Session struct {
	// Identifier of the client
	ClientID string `json:"clientID"`

	// Identifier of a code snippet which which is the environment associated
	CodeSnippetID string `json:"codeSnippetID"`

	// Information if the session is a shared persistent edit session
	EditEnabled bool `json:"editEnabled"`

	// Identifier of the session
	SessionID string `json:"sessionID"`
}

// Template defines model for Template.
type Template string

// ApiKeyOpt defines model for apiKeyOpt.
type ApiKeyOpt = string

// ApiKeyReq defines model for apiKeyReq.
type ApiKeyReq = string

// CodeSnippetID defines model for codeSnippetID.
type CodeSnippetID = string

// SessionID defines model for sessionID.
type SessionID = string

// N400 defines model for 400.
type N400 = Error

// N401 defines model for 401.
type N401 = Error

// N500 defines model for 500.
type N500 = Error

// DeleteEnvsCodeSnippetIDParams defines parameters for DeleteEnvsCodeSnippetID.
type DeleteEnvsCodeSnippetIDParams struct {
	ApiKey *ApiKeyReq `form:"api_key,omitempty" json:"api_key,omitempty"`
}

// PatchEnvsCodeSnippetIDParams defines parameters for PatchEnvsCodeSnippetID.
type PatchEnvsCodeSnippetIDParams struct {
	ApiKey *ApiKeyReq `form:"api_key,omitempty" json:"api_key,omitempty"`
}

// PostEnvsCodeSnippetIDJSONBody defines parameters for PostEnvsCodeSnippetID.
type PostEnvsCodeSnippetIDJSONBody = NewEnvironment

// PostEnvsCodeSnippetIDParams defines parameters for PostEnvsCodeSnippetID.
type PostEnvsCodeSnippetIDParams struct {
	ApiKey *ApiKeyReq `form:"api_key,omitempty" json:"api_key,omitempty"`
}

// PutEnvsCodeSnippetIDStateJSONBody defines parameters for PutEnvsCodeSnippetIDState.
type PutEnvsCodeSnippetIDStateJSONBody = EnvironmentStateUpdate

// PutEnvsCodeSnippetIDStateParams defines parameters for PutEnvsCodeSnippetIDState.
type PutEnvsCodeSnippetIDStateParams struct {
	ApiKey *ApiKeyReq `form:"api_key,omitempty" json:"api_key,omitempty"`
}

// GetSessionsParams defines parameters for GetSessions.
type GetSessionsParams struct {
	ApiKey *ApiKeyReq `form:"api_key,omitempty" json:"api_key,omitempty"`
}

// PostSessionsJSONBody defines parameters for PostSessions.
type PostSessionsJSONBody = NewSession

// PostSessionsParams defines parameters for PostSessions.
type PostSessionsParams struct {
	ApiKey *ApiKeyOpt `form:"api_key,omitempty" json:"api_key,omitempty"`
}

// DeleteSessionsSessionIDParams defines parameters for DeleteSessionsSessionID.
type DeleteSessionsSessionIDParams struct {
	ApiKey *ApiKeyReq `form:"api_key,omitempty" json:"api_key,omitempty"`
}

// PostSessionsSessionIDRefreshParams defines parameters for PostSessionsSessionIDRefresh.
type PostSessionsSessionIDRefreshParams struct {
	ApiKey *ApiKeyOpt `form:"api_key,omitempty" json:"api_key,omitempty"`
}

// PostEnvsCodeSnippetIDJSONRequestBody defines body for PostEnvsCodeSnippetID for application/json ContentType.
type PostEnvsCodeSnippetIDJSONRequestBody = PostEnvsCodeSnippetIDJSONBody

// PutEnvsCodeSnippetIDStateJSONRequestBody defines body for PutEnvsCodeSnippetIDState for application/json ContentType.
type PutEnvsCodeSnippetIDStateJSONRequestBody = PutEnvsCodeSnippetIDStateJSONBody

// PostSessionsJSONRequestBody defines body for PostSessions for application/json ContentType.
type PostSessionsJSONRequestBody = PostSessionsJSONBody
