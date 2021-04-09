# Go client for Apache Ozone

This repository contains an experimental, **proof-of-concept** Golang client for Apache Ozone.

It's not ready yet for using in production / non-production.

The repository contains the following sub-modules

 * api: the location for the generic golang api for Apache Ozone
 * cli: standalone executable tool for main operations (similar to the original `ozone sh`)
 * lib: proof-of-concept shared C library
 * python: example python script uses the shared C library
 * HA failovoer is missing

Status:

 * api
   * main OM metadata operations worked well, but not all the fields are implemented
   * data read / write are implemented on some level but needs further work
   * security can be supported by the used Hadoop RPC implementation, but not tested
 * fuse
   * first working POC, files are successfully listed and files can be read
   * write is not implemented at all
   * requires major work
  * shared lib / python: very basic example, poc

## Testing with cli:

```
cd cli
./build.sh
./ozone-go --om localhost volume create vol1
```

Or you can install it:

```
cd cli
go install
ozone -om 127.0.0.1 volume create vol1
```

## Testing Fuse file system

```
cd fuse
./build.sh
./ozone-fuse/ozone-fuse --om localhost --volume vol1 --bucket bucket1 /tmp/bucket1
```

## Testing the python binding

Create the shared library:

```
go build -o ozone.so -buildmode=c-shared lib/lib.go
```

Modify parameters of `python/test.py` (om address) and run it.
