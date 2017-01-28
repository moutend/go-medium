// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

// Token represents response from /v1/tokens.
type Token struct {
	TokenType    string   `json:"token_type": "`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	ExpiresAt    int      `json:"expires_at"`
}
