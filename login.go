package kickcom

import (
	"context"
	"fmt"
	"net/http"
)

// TBDLogin is a To-Be-Developed (TBD) function that will do the logging in (when will be complete).
// DO NOT USE THIS FUNCTION.
func (k *Kick) TBDLogin(
	ctx context.Context,
	email string,
	password string,
) error {
	// See: https://github.com/fb-sean/kick-website-endpoints/blob/main/APP_LOGIN_FLOW.md

	token, err := k.getToken(ctx)
	if err != nil {
		return fmt.Errorf("unable to get a token: %w", err)
	}

	//k.mobileLogin(ctx)

	_ = token
	panic("not implemented")
}

// KickTokenProviderReply is a reply to a token request.
type KickTokenProviderReply struct {
	Enabled                   bool   `json:"enabled"`
	NameFieldName             string `json:"nameFieldName"`
	UnrandomizedNameFieldName string `json:"unrandomizedNameFieldName"`
	ValidFromFieldName        string `json:"validFromFieldName"`
	EncryptedValidFrom        string `json:"encryptedValidFrom"`
}

func (k *Kick) getToken(ctx context.Context) (*KickTokenProviderReply, error) {
	return Request[KickTokenProviderReply](ctx, k, http.MethodGet, RouteKickTokenCreate, nil, nil, NoBody)
}

// MobileLoginRequest is a request for mobile logging in.
type MobileLoginRequest struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	OneTimePassword    string `json:"one_time_password"`
	KickTokenX         string `json:"_kick_token_xxxxxxxxxx"` // ?
	KickTokenValidFrom string `json:"_kick_token_valid_from"`
	IsMobileRequest    bool   `json:"isMobileRequest"`
}

// MobileLoginReply is a reply to mobile logging in request.
type MobileLoginReply struct {
	TwoFARequired bool `json:"2fa_required"`
}

/*func (k *Kick) mobileLogin(ctx context.Context) (*MobileLoginReply, error) {
	return Request[MobileLoginReply](ctx, k, http.MethodPost, "mobile/login", nil, NoBody)
}*/
