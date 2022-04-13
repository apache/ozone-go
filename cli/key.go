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
    "errors"
    "fmt"
    "github.com/apache/ozone-go/api"
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"
    "github.com/apache/ozone-go/api/utils"
    "os"
    "path"
    "strings"

    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

const keyOpName = "key"

var printType = "json"
var force = false
var limit = int32(100)
var prefix = "/"

// KeyCmd TODO
var KeyCmd = &cobra.Command{
    Use:   "key",
    Short: "Ozone key related operations",
    Long:  `Ozone key operate like: list/put/get/delete`,
}

var keyListCommand = &cobra.Command{
    Use:     "list",
    Aliases: []string{"l", "ls"},
    Short:   "List ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return listKeys(args)
    },
}

func listKeys(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        if keys, err := ozoneClient.ListKeys(address.Volume, address.Bucket, prefix, limit); err != nil {
            return err
        } else {
            for _, key := range keys {
                println(key.String())
            }
        }
    }
    return nil
}

var keyInfoCommand = &cobra.Command{
    Use:   "info",
    Short: "Info ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return infoKeys(args)
    },
}

func infoKeys(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        key, err := ozoneClient.InfoKey(address.Volume, address.Bucket, address.Key)
        if err != nil {
            return err
        }
        if key == nil {
            log.Error(args, " keys not found")
        } else {
            println(key.String())
        }
    }
    return nil
}

var touchzCommand = &cobra.Command{
    Use:   "touchz",
    Short: "touch zero size ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return touchzKeys(args)
    },
}

func touchzKeys(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        key, err := ozoneClient.TouchzKey(address.Volume, address.Bucket, address.Key)
        if err != nil {
            return err
        }
        println(key.String())
    }
    return nil
}

var keyPutCommand = &cobra.Command{
    Use:   "put",
    Short: "Put ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return putKeys(args)
    },
}

func putKeys(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    obj, localFile := args[0], args[1]
    address := common.OzoneObjectAddressFromPath(obj)
    if force {
        if _, err := ozoneClient.InfoKey(address.Volume, address.Bucket, address.Key); err != nil {
            if res, err := ozoneClient.DeleteKey(address.Volume, address.Bucket, address.Key); err != nil {
                return fmt.Errorf("force delete %s result %s error:%v", obj, res, err)
            }
        }
    }
    fileInfo, err := os.Stat(localFile)
    if err != nil {
        return err
    }
    f, err := os.Open(localFile)
    defer f.Close()
    if err != nil {
        return err
    }
    return ozoneClient.PutKey(address.Volume, address.Bucket, address.Key, fileInfo.Size(), f)
}

var keyGetCommand = &cobra.Command{
    Use:   "get",
    Short: "Get ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return getKeys(args)
    },
}

// getKeys todo: add glob patterns, process multi get
func getKeys(args []string) error {
    var dstPath string
    var err error
    obj := args[0]
    if len(args) == 1 {
        if dstPath, err = os.Getwd(); err != nil {
            return err
        }
        dstPath = dstPath + "/" + path.Base(obj)
    } else {
        dstPath = args[len(args)-1]
    }
    dstPath = path.Clean(dstPath)
    if force && utils.Exists(dstPath) {
        os.Remove(dstPath)
    }
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }

    address := common.OzoneObjectAddressFromPath(obj)
    f, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    defer f.Close()
    if err != nil {
        return err
    }
    err = ozoneClient.GetKey(address.Volume, address.Bucket, address.Key, f)
    if err != nil {
        if strings.Contains(err.Error(), "key not found") {
            log.Error(address.Key, " ", err.Error())
            return nil
        }
        return err
    }
    return nil

}

var keyFastPutCommand = &cobra.Command{
    Use:   "fastput",
    Short: "Fast put ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return fastPutKeys(args)
    },
}

func fastPutKeys(args []string) error {
    return errors.New("panic")
}

var keyFastGetCommand = &cobra.Command{
    Use:   "fastget",
    Short: "Fast get ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return fastGetKeys(args)
    },
}

func fastGetKeys(args []string) error {
    panic("")
}

var keyCatCommand = &cobra.Command{
    Use:   "cat",
    Short: "Cat ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return catKeys(args)
    },
}

// catKeys cat key to stdout
func catKeys(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        err := ozoneClient.GetKey(address.Volume, address.Bucket, address.Key, os.Stdout)
        if err != nil {
            if strings.Contains(err.Error(), "key not found") {
                log.Error(address.Key, " ", err.Error())
                return nil
            }
            return err
        }
    }
    return nil
}

var keyRenameCommand = &cobra.Command{
    Use:     "rename",
    Aliases: []string{"mv"},
    Short:   "Rename ozone key",
    RunE: func(cmd *cobra.Command, args []string) error {
        return renameKey(args)
    },
}

func renameKey(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    src, dst := args[0], args[1]
    address := common.OzoneObjectAddressFromPath(src)
    if force {
        if _, err := ozoneClient.InfoKey(address.Volume, address.Bucket, address.Key); err != nil {
            ozoneClient.DeleteKey(address.Volume, address.Bucket, address.Key)
        }
    }
    if unRenamed, err := ozoneClient.RenameKey(address.Volume, address.Bucket, address.Key, dst); err != nil {
        return err
    } else if len(unRenamed) == 0 {
        return errors.New("Not rename completely，unRenamed: " + unRenamed)
    }
    return nil
}

var keyCopyCommand = &cobra.Command{
    Use:     "copy",
    Aliases: []string{"cp"},
    Short:   "Copy ozone key",
    RunE: func(cmd *cobra.Command, args []string) error {
        return copyKey(args)
    },
}

func copyKey(args []string) error {
    panic("")
}

var keyDeleteCommand = &cobra.Command{
    Use:     "delete",
    Aliases: []string{"d", "rm"},
    Short:   "Delete ozone keys",
    RunE: func(cmd *cobra.Command, args []string) error {
        return deleteKeys(args)
    },
}

func deleteKeys(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        if result, err := ozoneClient.DeleteKey(address.Volume, address.Bucket, address.Key); err != nil {
            return fmt.Errorf("delete result %s error:%v", result, err)
        }
    }
    return nil
}

func init() {
    KeyCmd.AddCommand(keyInfoCommand)
    KeyCmd.AddCommand(touchzCommand)
    KeyCmd.AddCommand(keyListCommand)
    KeyCmd.AddCommand(keyGetCommand)
    KeyCmd.AddCommand(keyPutCommand)
    KeyCmd.AddCommand(keyCatCommand)
    KeyCmd.AddCommand(keyRenameCommand)
    KeyCmd.AddCommand(keyCopyCommand)
    KeyCmd.AddCommand(keyFastGetCommand)
    KeyCmd.AddCommand(keyFastPutCommand)
    KeyCmd.AddCommand(keyDeleteCommand)
    keyListCommand.Flags().StringVarP(&printType, "printType", "t", "json", "print type")
    keyListCommand.Flags().Int32VarP(&limit, "limit", "n", 100, "number of list")
    keyListCommand.Flags().StringVarP(&prefix, "prefix", "p", "", "prefix to list")
    keyPutCommand.Flags().BoolVarP(&force, "force", "f", false, "force to to")
    keyGetCommand.Flags().BoolVarP(&force, "force", "f", false, "force to to")
}
