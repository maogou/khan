package demo

import (
	"smallBot/internal/pkg/help"
	"smallBot/internal/sdk/khan"

	"github.com/rs/zerolog/log"

	"github.com/urfave/cli/v2"
)

func Demo(sdk *khan.Khan) *cli.Command {
	return &cli.Command{
		Name:  "demo",
		Usage: "测试khan的接口",

		Action: func(cCtx *cli.Context) error {
			return do(cCtx, sdk)
		},
	}
}

func do(cCtx *cli.Context, sdk *khan.Khan) error {
	log.Info().Msg("i am demo!")

	//037ecf6351b795aebb07e43fd8451a3f

	key := "037ecf6351b795aebb07e43fd8451a3f"
	encrypted, err := help.AesEncrypt("hello world", []byte(key))

	if err != nil {
		log.Error().Err(err).Msg("aes encrypt error")
		return err
	}

	decrypted, err := help.AesDecrypt(encrypted, []byte(key))

	if err != nil {
		log.Error().Err(err).Msg("aes decrypt error")
		return err
	}

	log.Info().Msg("encrypted: " + decrypted)
	return nil
}
