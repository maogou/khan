package auth

import (
	"crypto/ed25519"
	"encoding/base64"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/license/licenseutil"

	"github.com/rs/zerolog/log"

	"github.com/pterm/pterm"

	"github.com/urfave/cli/v2"
)

func Verify() *cli.Command {
	return &cli.Command{
		Name:  "verify",
		Usage: "授权许可证验证",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "key",
				Value:    "",
				Usage:    "颁发的软件许可证",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			key := cCtx.String("key")

			if res, err := parse(key); err != nil {
				log.Error().Err(err).Msg("授权许可证验证失败")
				return err
			} else {
				log.Info().Msg("许可证信息:")
				return printLicense(res)
			}
		},
	}
}

func parse(key string) (*license.License, error) {
	var (
		publicKey ed25519.PublicKey
		result    *license.License
		err       error
	)
	if publicKey, err = licenseutil.ReadPublicKeyFromStr(constant.PublicKey); err != nil {
		return nil, err
	}

	var bKey = make([]byte, base64.StdEncoding.DecodedLen(len(key)))
	if _, err = base64.StdEncoding.Decode(bKey, []byte(key)); err != nil {
		return nil, err
	}

	if result, err = license.DecodeStr(string(bKey), publicKey); err != nil {
		return nil, license.ErrInvalidSignature
	}

	return result, nil
}

func printLicense(license *license.License) error {
	return pterm.DefaultTable.WithHasHeader().
		WithBoxed().
		WithData(
			pterm.TableData{
				{"客户标识", license.Cus},
				{"微信号", license.Sub},
				{"发证单位", license.Iss},
				{"签发时间", license.Iat.Format("2006-01-02 15:04:05")},
				{"过期时间", license.Exp.Format("2006-01-02 15:04:05")},
			},
		).Render()
}
