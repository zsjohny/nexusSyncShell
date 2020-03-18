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
	"time"
)

var cfg = new(core.Config)
var versionFlag bool

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
	flagSet.StringVarP(&cfg.RemoteUrl, "remoteUrl", "u", "", "[Required]format: https://nexus.huya.com")
	flagSet.StringVarP(&cfg.Option, "option", "X", "", "[Required]format:POST/GET\n"+
		"upload:POST, download:GET")
	flagSet.StringVarP(&cfg.Auth, "auth", "a", "", "[Required]basic auth user info,format: usr:pwd")
	flagSet.IntVarP(&cfg.Process, "proc", "p", 0, "[Not Required][default:1]the proc num, but no bigger than the proc of executing machine")
	flagSet.StringVarP(&cfg.RemoteDir, "remoteDir", "d", "", "[Required]the dir that you want to download/upload of nexus")
	flagSet.StringVarP(&cfg.LocalDir, "localDir", "l", "", "[Required]the dir that you want to download/upload of local machine")
	flagSet.StringVarP(&cfg.Repository, "repo", "r", "", "[Required]the repository of nexus")
	flagSet.BoolVarP(&versionFlag, "version", "v", false, "version" )

}
func init() {
	initFlags()
}

func run() error {
	t1 := time.Now()
	if versionFlag == true{
		fmt.Printf("nexusSync version is v%s\n", Version)
		return nil
	}
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
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	return nil
}

func checkParam(config *core.Config) error {
	reg := `(http(s)?:\/\/)?(www\.)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+(:\d+)*(\/\w+\.\w+)*$`
	if flag, _ := regexp.MatchString(reg, config.RemoteUrl); !flag {
		return fmt.Errorf("param remoteUrl(-u) is not valid\n" +
			"format: https://nexus.huya.com")
	}

	//仓库不能为空
	if len(cfg.Repository) <= 0 {
		return fmt.Errorf("param [repo(-r)] is not valid")
	}
	//设置线程数，建议少于CPU核数,如果超过核数将默认为核数大小
	if cfg.Process == 0 {
		runtime.GOMAXPROCS(1)
		fmt.Printf("process will use the default config, process = %d\n", 1)
	}
	if cfg.Process > runtime.NumCPU(){
		fmt.Printf("process will use the default config, process = %d\n", runtime.NumCPU())
	}
	//验证user信息
	if len(config.Auth) < 1 {
		return fmt.Errorf("param [auth(-u)] is not valid\n" +
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
		return fmt.Errorf("param [localDir(-l)] is not valid")
	}

	//规范格式
	if len(config.RemoteDir) <= 0 {
		return fmt.Errorf("param [remoteDir(-r)] is not valid,standard format is /a/b")
	}
	config.RemoteDir = util.FormatNexusPathSeparator(config.RemoteDir)

	return nil
}
