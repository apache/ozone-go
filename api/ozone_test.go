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
package api

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
import "github.com/stretchr/testify/assert"

func randomName(prefix string) string {
	seq := rand.New(rand.NewSource(time.Now().Unix())).Int31()
	return prefix + fmt.Sprintf("%d", seq)
}

func TestOzoneClientVolumeCreateGet(t *testing.T) {
	client := CreateOzoneClient("localhost")

	volumeName := randomName("vol")
	err := client.CreateVolume(volumeName)
	assert.Nil(t, err)

	vol, err := client.GetVolume(volumeName)
	assert.Nil(t, err)

	assert.Equal(t, volumeName, vol.Name)
}

func TestOzoneClientBucketCreateGet(t *testing.T) {

	client := CreateOzoneClient("localhost")

	//volumeName := "vol1"
	volumeName := randomName("vol")
	bucketName := randomName("bucket")

	err := client.CreateVolume(volumeName)
	assert.Nil(t, err)

	time.Sleep(4 * time.Second)




	err = client.CreateBucket(volumeName, bucketName)
	assert.Nil(t, err)

	bucket, err := client.GetBucket(volumeName, bucketName)
	assert.Nil(t, err)

	assert.Equal(t, bucketName, bucket.Name)
	assert.Equal(t, volumeName, bucket.VolumeName)
}
