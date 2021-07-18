package claims

// TokenClaimTag token claims tag which will be set in token
type TokenClaimTag string

const (
	// Claims the text sits instead of claims
	Claims = "claims"
	// Username username
	Username TokenClaimTag = "username"
	// TokenRoles roles tag
	TokenRoles TokenClaimTag = "roles"
	// ExpiryTime exp time
	ExpiryTime TokenClaimTag = "exp"
	// OSVersion os version
	OSVersion TokenClaimTag = "os_version"
	// APPVersion app version
	APPVersion TokenClaimTag = "app_version"
	// APIVersion api version
	APIVersion TokenClaimTag = "api_version"
	// Platform used platform
	Platform TokenClaimTag = "platform"
	// DeviceID device id
	DeviceID TokenClaimTag = "deviceId"
	// Issuer aka iss
	Issuer TokenClaimTag = "iss"
	// Subject aka sub
	Subject TokenClaimTag = "sub"
	// NotBefore nbf
	NotBefore TokenClaimTag = "nbf"
	// IssuedAt iat
	IssuedAt TokenClaimTag = "iat"
	// ClientIP client ip
	ClientIP TokenClaimTag = "ipv4"
	// Scopes scopes
	Scopes TokenClaimTag = "scopes"
	// Audience aka aud
	Audience TokenClaimTag = "aud"
	// Agent agent
	Agent TokenClaimTag = "agent"
	// UserSub user_sub
	UserSub TokenClaimTag = "user_sub"

	// Token token tag
	Token TokenClaimTag = "token"
	// RefreshToke refresh token tag
	RefreshToke TokenClaimTag = "refresh_token"
	// Exp expiry time
	Exp TokenClaimTag = "exp"
)
