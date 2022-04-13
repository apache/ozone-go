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
    "fmt"
    "github.com/apache/ozone-go/api"
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"
    "os"
    "text/tabwriter"

    "github.com/spf13/cobra"
)

const fsOpName = "fs"

var lsLimit = uint64(1024)
var startKey = ""
var humanReadable = false

var fsListCommand = &cobra.Command{
    Use:     "list",
    Aliases: []string{"l", "ls"},
    Short:   "List ozone files",
    RunE: func(cmd *cobra.Command, args []string) error {
        return listFiles(args)
    },
}

func listFiles(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromFsPath(arg)
        if address.IsRootPath() {
            if files, err := ozoneClient.ListVolumes(); err != nil {
                return err
            } else {
                fmt.Println(fmt.Sprintf("Found %d items", len(files)))
                lines := make([]string, 0)
                for _, f := range files {
                    lines = append(lines, f.FriendlyFileInfoString(humanReadable))
                }
                PrettyPrintFiles(lines)
            }
        } else if len(address.Volume) > 0 && len(address.Bucket) == 0 {
            if files, err := ozoneClient.ListBucket(address.Volume); err != nil {
                return err
            } else {
                fmt.Println(fmt.Sprintf("Found %d items", len(files)))
                lines := make([]string, 0)
                for _, f := range files {
                    lines = append(lines, f.FriendlyFileInfoString(humanReadable))
                }
                PrettyPrintFiles(lines)
            }
        } else {
            if files, err := ozoneClient.ListFiles(address.Volume, address.Bucket, address.Key, startKey, lsLimit); err != nil {
                return err
            } else {
                fmt.Println(fmt.Sprintf("Found %d items", len(files)))
                lines := make([]string, 0)
                for _, f := range files {
                    lines = append(lines, f.FriendlyFileInfoString(humanReadable))
                }
                PrettyPrintFiles(lines)
            }
        }
    }
    return nil
}

// PrettyPrintFiles TODO
func PrettyPrintFiles(lines []string) {
    tw := tabwriter.NewWriter(os.Stdout, 3, 8, 0, ' ', tabwriter.AlignRight|tabwriter.TabIndent)
    for _, line := range lines {
        _, _ = fmt.Fprintf(tw, "%s\n", line)
    }
    tw.Flush()
}

func init() {
    FsCmd.AddCommand(fsListCommand)
    fsListCommand.Flags().BoolVarP(&humanReadable, "humanReadable", "H", false, "human read for data size")
    fsListCommand.Flags().Uint64VarP(&lsLimit, "lsLimit", "n", 1024, "number of list")
    fsListCommand.Flags().StringVarP(&startKey, "startKey", "s", "", "startKey to list")
}
