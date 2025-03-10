package internal

import (
	"os"
	"smallBot/internal/command/auth"
	"smallBot/internal/command/demo"
	"smallBot/internal/command/license"
	"smallBot/internal/command/login"
	"smallBot/internal/command/serve"
	"smallBot/internal/command/upgrade"
	"smallBot/internal/command/version"
	"smallBot/internal/config"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/db"
	"smallBot/internal/sdk/khan"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/rs/zerolog"
)

var (
	conf        config.Config
	sdk         *khan.Khan
	client      = resty.New()
	validate    = validator.New()
	description = `免责说明:
	一.本框架仅供学习和技术研究使用,不得用于任何商业或非法行为,否则后果自负.
	二.本框架的作者保留随时修改、更新、删除或终止本工具的权利,无需事先通知或承担任何义务.
	三.本框架的使用者应遵守相关法律法规，尊重微信的版权和隐私,不得侵犯微信或其他第三方的合法权益,不得从事任何违法或不道德的行为.
	四.本框架的使用者在下载安装、运行或使用本工具时,即表示已阅读并同意本免责声明.如有异议,请立即停止使用本工具,并删除所有相关文件.
	五.本框架的作者不对本工具的安全性、完整性、可靠性、有效性、正确性或适用性做任何明示或暗示的保证,也不对本工具的使用或滥用造成的任何直接或间接的损失、责任、索赔、要求或诉讼承担任何责任.`
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		if _, after, found := strings.Cut(file, "build"); found {
			return strings.TrimLeft(after, "../") + ":" + strconv.Itoa(line)
		}
		return file + ":" + strconv.Itoa(line)
	}
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05.000"})
}

func NewKhanCommand() *cli.App {
	conf = config.MustLoadConfig("")
	rdb := db.MustInitRedis(&conf)
	sdk = khan.NewKhanSdk(&conf, client, validate, rdb)

	robot := &cli.App{
		Name:        "khan",
		Usage:       "khan-超简单超稳定的机器人",
		Copyright:   "(c) 2025 khanwchat632@gmail.com",
		Version:     constant.VERSION,
		Description: description,
		Commands: []*cli.Command{
			serve.Start(conf, sdk),
			version.Info(),
			auth.Verify(conf),
			login.Login(conf, sdk),
			upgrade.Upgrade(),
			license.Create(sdk),
			demo.Demo(sdk),
		},
	}

	return robot
}
