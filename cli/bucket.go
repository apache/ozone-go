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

    "github.com/spf13/cobra"
)

const bucketOpName = "bucket"

// BucketCmd TODO
var BucketCmd = &cobra.Command{
    Use:   "bucket",
    Short: "Ozone bucket related operations",
    Long:  `Ozone bucket operate like: list/create/delete`,
}

var bucketCreateCommand = &cobra.Command{
    Use:     "create",
    Aliases: []string{"crt", "mk"},
    Short:   "Create ozone buckets",
    RunE: func(cmd *cobra.Command, args []string) error {
        return createBuckets(args)
    },
}

var bucketListCommand = &cobra.Command{
    Use:     "list",
    Aliases: []string{"l", "ls"},
    Short:   "List ozone buckets",
    RunE: func(cmd *cobra.Command, args []string) error {
        return listBuckets(args)
    },
}

var bucketInfoCommand = &cobra.Command{
    Use:   "info",
    Short: "Info ozone buckets",
    RunE: func(cmd *cobra.Command, args []string) error {
        return infoBuckets(args)
    },
}

var bucketDeleteCommand = &cobra.Command{
    Use:     "delete",
    Aliases: []string{"d", "rm"},
    Short:   "Delete ozone buckets",
    RunE: func(cmd *cobra.Command, args []string) error {
        return deleteBuckets(args...)
    },
}

func deleteBuckets(args ...string) error {
    panic("")
}

func createBuckets(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        err := ozoneClient.CreateBucket(address.Volume, address.Bucket)
        if err != nil {
            return err
        }
    }
    return nil
}

func infoBuckets(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        bucket, err := ozoneClient.InfoBucket(address.Volume, address.Bucket)
        if err != nil {
            return err
        }
        println(bucket.String())
    }
    return nil
}

func listBuckets(args []string) error {
    ozoneClient, err := api.CreateOzoneClient(config.OmAddress)
    if err != nil {
        return err
    }
    for _, arg := range args {
        address := common.OzoneObjectAddressFromPath(arg)
        buckets, err := ozoneClient.ListBucket(address.Volume)
        if err != nil {
            return err
        }
        for _, bucket := range buckets {
            println(bucket.String())
        }
    }
    return nil
}

func init() {
    BucketCmd.AddCommand(bucketCreateCommand)
    BucketCmd.AddCommand(bucketInfoCommand)
    BucketCmd.AddCommand(bucketListCommand)
	BucketCmd.AddCommand(bucketDeleteCommand)
}
