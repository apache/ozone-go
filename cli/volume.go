// Package main TODO
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

import (
    "github.com/apache/ozone-go/api"
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"

    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

const volumeOpName = "volume"

// VolumeCmd TODO
var VolumeCmd = &cobra.Command{
    Use:     "volume",
    Aliases: []string{"v", "vol"},
    Short:   "Ozone volume related operations",
    Long:    `Ozone volume operate like: list/create/delete`,
}

var volumeCreateCommand = &cobra.Command{
    Use:     "create",
    Aliases: []string{"crt", "mk"},
    Short:   "Create ozone volumes",
    RunE: func(cmd *cobra.Command, args []string) error {
        return creatVolumes(args)
    },
}

var volumeListCommand = &cobra.Command{
    Use:     "list",
    Aliases: []string{"l", "ls"},
    Short:   "List ozone volumes",
    RunE: func(cmd *cobra.Command, args []string) error {
        return listVolumes()
    },
}

var volumeInfoCommand = &cobra.Command{
    Use:   "info",
    Short: "Info ozone volumes",
    RunE: func(cmd *cobra.Command, args []string) error {
        return infoVolumes(args)
    },
}

var volumeDeleteCommand = &cobra.Command{
    Use:     "delete",
    Aliases: []string{"d", "rm"},
    Short:   "Delete ozone volumes",
    RunE: func(cmd *cobra.Command, args []string) error {
        return deleteVolumes(args)
    },
}

func deleteVolumes(args []string) error {
    panic("")
}

func listVolumes() error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    volumes, err := ozoneClient.ListVolumes()
    if err != nil {
        logrus.Errorf("List all volumes error: %v", err)
        return err
    }
    for _, volume := range volumes {
        println(volume.String())
    }
    return nil
}

func infoVolumes(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        volume, err := ozoneClient.InfoVolume(arg)
        if err != nil {
            logrus.Errorf("Info volume %s error: %v", arg, err)
            return err
        }
        println(volume.String())
    }
    return nil
}

func creatVolumes(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        err := ozoneClient.CreateVolume(address.Volume)
        if err != nil {
            logrus.Errorf("Create %s error: %v", arg, err)
            return err
        }
    }
    return nil
}

func init() {
    VolumeCmd.AddCommand(volumeCreateCommand)
    VolumeCmd.AddCommand(volumeInfoCommand)
    VolumeCmd.AddCommand(volumeListCommand)
	VolumeCmd.AddCommand(volumeDeleteCommand)
}
