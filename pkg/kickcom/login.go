package kickcom

import (
	"context"
	"fmt"
	"net/http"
)

func (k *Kick) Login(
	ctx context.Context,
	email string,
	password string,
) error {
	// See: https://github.com/fb-sean/kick-website-endpoints/blob/main/APP_LOGIN_FLOW.md

	token, err := k.getToken(ctx)
	if err != nil {
		return fmt.Errorf("unable to get a token: %w", err)
	}

	k.mobileLogin(ctx)

	_ = token
	panic("not implemented")
}

type KickTokenProviderReply struct {
	Enabled                   bool   `json:"enabled"`
	NameFieldName             string `json:"nameFieldName"`
	UnrandomizedNameFieldName string `json:"unrandomizedNameFieldName"`
	ValidFromFieldName        string `json:"validFromFieldName"`
	EncryptedValidFrom        string `json:"encryptedValidFrom"`
}

func (k *Kick) getToken(ctx context.Context) (*KickTokenProviderReply, error) {
	return Request[KickTokenProviderReply](ctx, k, http.MethodGet, "kick-token-provider", struct{}{})
}

type MobileLoginRequest struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	OneTimePassword    string `json:"one_time_password"`
	KickTokenX         string `json:"_kick_token_xxxxxxxxxx"` // ?
	KickTokenValidFrom string `json:"_kick_token_valid_from"`
	IsMobileRequest    bool   `json:"isMobileRequest"`
}

type MobileLoginReply struct {
	TwoFARequired bool `json:"2fa_required"`
}

func (k *Kick) mobileLogin(ctx context.Context) (*MobileLoginReply, error) {
	return Request[MobileLoginReply](ctx, k, http.MethodPost, "mobile/login", struct{}{})
}
