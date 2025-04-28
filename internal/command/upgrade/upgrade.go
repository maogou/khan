package upgrade

import (
	"maogou/khan/internal/constant"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v2"
)

func Upgrade() *cli.Command {
	cmd := &cli.Command{
		Name:        "upgrade",
		Usage:       "升级khan到最新版本",
		Description: "使用示例: khan upgrade",

		Action: func(cCtx *cli.Context) error {
			return execute()
		},
	}

	return cmd
}

func execute() error {
	arg := "GO111MODULE=" + constant.GO111MODULE + " GOPROXY=" + constant.GOPROXY + " go install " + constant.GithubRepo
	execCmd := exec.Command("sh", "-c", arg)

	if runtime.GOOS == constant.Windows {
		arg = "set GOPROXY=" + constant.GOPROXY + " && go install " + constant.GithubRepo
		execCmd = exec.Command("cmd.exe", "/c", arg)

	}

	if err := execCmd.Run(); err != nil {
		return err
	}

	return nil
}
