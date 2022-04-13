// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// import "github.com/apache/ozone-go"
import (
    "fmt"
    "os"
    "os/user"
    "github.com/apache/ozone-go/api/config"
    "github.com/apache/ozone-go/api/utils"
    "strings"

    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

// Version TODO
const Version = "v1.0.0"

// Date TODO
const Date = "20220218"

// Commit TODO
const Commit = "love"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "github.com/apache/ozone-go",
    Short: "Native Ozone command line client",
}

// ShCmd TODO
var ShCmd = &cobra.Command{
    Use:   "sh",
    Short: "Command line interface for object store operations",
}

// FsCmd TODO
var FsCmd = &cobra.Command{
    Use:   "fs",
    Short: "Run a filesystem command on Ozone file system. Equivalent to 'hadoop fs'",
}

// VersionCmd TODO
var VersionCmd = &cobra.Command{
    Use:   "version",
    Short: fmt.Sprintf("%s (%s, %s)", Version, Commit, Date),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(cmd.Short)
    },
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
    if err := config.InitLog(); err != nil {
        log.Errorf("init logrus config error: %v ", err)
        os.Exit(1)
    }
    if len(config.User) == 0 {
        if u, err := user.Current(); err == nil {
            config.User = u.Name
        } else {
            log.Error(fmt.Sprintf("get current user error: %v ", err))
            os.Exit(1)
        }
    }
    if len(config.OmAddress) > 0 {
        config.OmAddresses = strings.Split(config.OmAddress, ",")
    } else {
        if utils.Exists(config.ConfFilePath) {
            config.OzoneConfig.ConfigPath = config.ConfFilePath
            if err := config.OzoneConfig.LoadFromFile(); err != nil {
                log.Error(fmt.Sprintf("parse config file %s error:%v", config.ConfFilePath, err))
                os.Exit(1)
            } else {
                if len(config.OmAddress) == 0 {
                    config.OmAddresses = config.OzoneConfig.GetOmAddresses()
                    config.OmAddress = config.OzoneConfig.NextOmNode()
                }
            }
        } else {
            log.Error(fmt.Sprintf("ozone config file %s not found", config.ConfFilePath))
            os.Exit(1)
        }
    }
}

// Execute adds all child commands to the rootPath command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    RootCmd.AddCommand(VersionCmd)
    RootCmd.AddCommand(FsCmd)
    RootCmd.AddCommand(ShCmd)
    ShCmd.AddCommand(VolumeCmd)
    ShCmd.AddCommand(BucketCmd)
    ShCmd.AddCommand(KeyCmd)
    RootCmd.PersistentFlags().StringVar(&config.ConfFilePath, "config", "ozone-site.xml",
        "config file")
    RootCmd.PersistentFlags().StringVar(&config.OmAddress, "om_client", "",
        "Ozone manager host address(host:port)")
    RootCmd.PersistentFlags().StringVar(&config.LogLevel, "loglevel",
        "INFO", "log level, default INFO")
    RootCmd.PersistentFlags().StringVar(&config.User, "user", "", "current user")
	cobra.OnInitialize(InitConfig)
}

func main() {
	Execute()
}
