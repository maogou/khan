package version

import (
	"maogou/khan/internal/constant"
	"runti
	"time"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

type Version struct {
	GoVersion  string
	Os         string
	BuildDate  string
	CommitId   string
	CommitDate string
}

func Info() *cli.Command {
	return &cli.Command{
		Name:  "info",
		Usage: "获取当前发布打包的版本信息",
		Action: func(cCtx *cli.Context) error {
			bv := getVersionInfo()
			return printVersionInfo(bv)
		},
	}
}

func getVersionInfo() Version {
	info, ok := debug.ReadBuildInfo()
	var bv Version
	if ok {
		bv.GoVersion = info.GoVersion
		bv.BuildDate = time.Now().Format(time.DateTime)
		for _, item := range info.Settings {
			switch item.Key {
			case "vcs.revision":
				bv.CommitId = item.Value
			case "vcs.time":
				bv.CommitDate = formatUtcTime(item.Value)
			case "GOOS", "GOARCH":
				bv.Os += item.Value + " "
			}
		}
	}
	return bv
}

func formatUtcTime(utcTimeStr string) string {
	utcTime, err := time.Parse(time.RFC3339, utcTimeStr)
	if err != nil {
		return utcTimeStr
	}
	return utcTime.Local().Format(time.DateTime)
}

func printVersionInfo(bv Version) error {
	return pterm.DefaultTable.WithData(
		pterm.TableData{
			{"Go版本号", bv.GoVersion},
			{"构建系统", bv.Os},
			{"版本号", constant.VERSION},
			{"协议", "iPad 8.0.48"},
			{"分支", "master"},
			{"CommitId", bv.CommitId},
			{"CommitDate", bv.CommitDate},
			{"作者", "khanwchat632@gmail.com"},
		},
	).Render()
}
