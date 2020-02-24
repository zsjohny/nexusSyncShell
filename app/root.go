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
	"runtime"
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
	flagSet.StringVarP(&cfg.RemoteDir, "remoteDir", "d", "/", "the dir that you want to download/upload of nexus")
	flagSet.StringVarP(&cfg.LocalDir, "localDir", "l", "", "the dir that you want to download/upload of local machine")

}
func init() {
	initFlags()
}

func run() error {
	fmt.Println("config", cfg)
	//设置线程数，必须少于CPU核数
	if cfg.Process <= 0 && cfg.Process < runtime.NumCPU() {
		runtime.GOMAXPROCS(cfg.Process)
	}
	var syncService = new(core.NexusSyncService)
	option := cfg.Option
	switch option {
	case util.Post:
		syncService.StartUpload(cfg)
	case util.Get:
		syncService.StartDownload(cfg)
	default:
		fmt.Errorf("option error, format:GET / POST")
	}
	return nil
}
