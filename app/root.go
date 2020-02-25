/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"nexusSync/core"
	"nexusSync/util"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var cfg = new(core.Config)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "nexusSyncShell",
	Short:             "nexusSyncShell",
	Long:              `nexusSyncShell`,
	DisableAutoGenTag: true, // disable displaying auto generation tag in cli docs
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initFlags() {
	flagSet := rootCmd.Flags()
	// url & output
	flagSet.StringVarP(&cfg.RemoteUrl, "remoteUrl", "r", "", "")
	flagSet.StringVarP(&cfg.Option, "option", "X", "", "format:POST/GET\n"+
		"upload:POST, download:GET")
	flagSet.StringVarP(&cfg.Auth, "auth", "u", "", "basic auth user info,format: usr:pwd")
	flagSet.IntVarP(&cfg.Process, "proc", "p", 0, "the proc num, but no bigger than the proc of executing machine")
	flagSet.StringVarP(&cfg.RemoteDir, "remoteDir", "d", "", "the dir that you want to download/upload of nexus")
	flagSet.StringVarP(&cfg.LocalDir, "localDir", "l", "", "the dir that you want to download/upload of local machine")

}
func init() {
	initFlags()
}

func run() error {
	err := checkParam(cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var syncService = new(core.NexusSyncService)
	option := cfg.Option
	option = strings.TrimSpace(option)
	switch option {
	case util.Post:
		fmt.Println("start upload")
		syncService.StartUpload(cfg)
	case util.Get:
		fmt.Println("start download")
		syncService.StartDownload(cfg)
	default:
		fmt.Println("option error, format:GET / POST")
	}
	return nil
}

func checkParam(config *core.Config) error {
	reg := `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`
	if flag, _ := regexp.MatchString(reg, config.RemoteUrl); !flag {
		return fmt.Errorf("param remoteUrl(-r) is not valid")
	}
	//设置线程数，建议少于CPU核数,如果超过核数将默认为核数大小
	if cfg.Process <= 0 || cfg.Process < runtime.NumCPU() {
		runtime.GOMAXPROCS(cfg.Process)
		fmt.Printf("process will use the default config, process = %d\n", runtime.NumCPU())
	}
	//验证user信息
	if len(config.Auth) < 1 {
		return fmt.Errorf("param Auth(-u) is not valid\n" +
			"format:usr:pwd")
	}
	auth := strings.Split(config.Auth, ":")
	if len(auth) < 2 {
		return fmt.Errorf("auth info error, plearse split with `:`")
	}
	config.Usr = auth[0]
	config.Pwd = auth[1]

	//验证localDir和规范格式
	config.LocalDir = util.FormatLocalPath(config.LocalDir)
	localDir := config.LocalDir
	if err := util.PathExists(localDir); err != nil {
		return fmt.Errorf("param localDir(-l) is not valid")
	}

	//规范格式
	if len(config.RemoteDir) <= 0 {
		return fmt.Errorf("param RemoteDir(-d) is not valid,standard format is /a/b")
	}
	config.RemoteDir = util.FormatNexusPathSeparator(config.RemoteDir)

	return nil
}
