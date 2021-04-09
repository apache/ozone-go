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

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

var version string
var commit string
var date string

type OzoneFs struct {
	pathfs.FileSystem
	ozoneClient *api.OzoneClient
	Volume      string
	Bucket      string
}

func (me *OzoneFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	if name == "" {
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	}
	keys, err := me.ozoneClient.OmClient.ListKeysPrefix(me.Volume, me.Bucket, name)
	if len(keys) == 0 {
		return nil, fuse.ENOENT
	}
	if err != nil {
		fmt.Println("Error with getting key: " + name + " " + err.Error())
		return nil, fuse.ENOENT
	}

	if *keys[0].KeyName == name {
		return &fuse.Attr{
			Mode: fuse.S_IFREG | 0644, Size: *keys[0].DataSize}, fuse.OK
	} else {
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	}
}

func (me *OzoneFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {

	keys, err := me.ozoneClient.ListKeysPrefix(me.Volume, me.Bucket, name)
	if err != nil {
		panic(err)
	}
	result := make([]fuse.DirEntry, 0)
	var lastDir = ""
	for _, key := range keys {
		keyName := key.Name
		relative := keyName[len(name):]
		if strings.HasPrefix(relative, "/") {
			relative = relative[1:]
		}
		levels := strings.Count(relative, "/")
		if levels > 0 {
			name := relative[0:strings.Index(relative, "/")]
			if name != lastDir {
				entry := fuse.DirEntry{Name: name, Mode: fuse.S_IFDIR}
				result = append(result, entry)
				lastDir = name
			}
		} else {
			entry := fuse.DirEntry{Name: relative, Mode: fuse.S_IFREG}
			result = append(result, entry)
		}

	}
	return result, fuse.OK

}

func (me *OzoneFs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	key, err := me.ozoneClient.OmClient.GetKey(me.Volume, me.Bucket, name)
	if err != nil {
		return nil, fuse.EACCES
	}
	return CreateOzoneFile(me.ozoneClient, key), fuse.OK
}


func main() {
	app := cli.NewApp()
	app.Name = "ozone-fuse"
	app.Usage = "Ozone fuse driver"
	app.Description = "FUSE filesystem driver for Apache Ozone"
	app.Version = fmt.Sprintf("%s (%s, %s)", version, commit, date)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "om",
			Required: true,
			Value:    "localhost",
			Usage:    "Host (or host:port) address of the OzoneManager",
		},
		cli.StringFlag{
			Name:     "volume",
			Required: true,
			Usage:    "Name of the volume to mount",
		},
		cli.StringFlag{
			Name:     "bucket",
			Required: true,
			Value:    "localhost",
			Usage:    "Name of the bucket to mount",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Turn on FUSE debug log",
		},
	}
	app.Action = func(c *cli.Context) error {
		client := api.CreateOzoneClient(c.String("om"))

		fs := &OzoneFs{FileSystem: pathfs.NewDefaultFileSystem(), ozoneClient: client, Volume: c.String("volume"), Bucket: c.String("bucket")}
		nfs := pathfs.NewPathNodeFs(fs, nil)
		opts := nodefs.Options{
			Debug: c.Bool("debug"),
		}
		server, _, err := nodefs.MountRoot(c.Args().Get(0), nfs.Root(), &opts)
		if err != nil {
			log.Fatalf("Mount fail: %v\n", err)
		}
		server.Serve()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
