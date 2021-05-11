#!/usr/bin/env python
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from ctypes import *

lib = cdll.LoadLibrary("../lib/lib")

lib.CreateOmClient.argtypes = [c_char_p]
lib.CreateOmClient.restype = c_long
lib.PrintKey.argtypes = [c_long, c_char_p, c_char_p, c_char_p]

client = lib.CreateOmClient(b"localhost")
print(client)
lib.PrintKey(client, b"vol1", b"bucket1", b"key1")
