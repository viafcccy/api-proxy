/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"api-proxy/config"
	"api-proxy/internal/forward"
	"fmt"

	"github.com/spf13/cobra"
)

// 以下变量优先使用命令行参数，如果没有使用配置文件
var targetHost string

// forwardCmd represents the forward command
var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := config.GetConfig()
		if err != nil {
			panic(fmt.Errorf("config load fail: %v", err))
		}

		if targetHost == "" {
			targetHost = cfg.Forward.TargetHost
		}

		forward.Run(targetHost, cfg.Forward.ServerPort)
	},
}

func init() {
	rootCmd.AddCommand(forwardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forwardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forwardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	forwardCmd.Flags().StringVarP(&targetHost, "target", "t", "", "target host")
}
