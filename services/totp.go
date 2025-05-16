package services

import (
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

type TOTPSetup struct {
	Secret string
	QRPNG  []byte
	URL    string
}

func GenerateTOTP(email string) (*TOTPSetup, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "blog-server",
		AccountName: email,
	})
	if err != nil {
		return nil, err
	}

	png, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return &TOTPSetup{
		Secret: key.Secret(),
		QRPNG:  png,
		URL:    key.URL(),
	}, nil
}

func ValidateTOTP(secret, token string) bool {
	return totp.Validate(token, secret)
}
