#!/usr/bin/env bash

version=3.1
githash=`git log --pretty=format:'%h' -n 1`

cat << EOF > ./version.go
package main

//DO NOT EDIT
//Build script overwrites this file.
//DO NOT EDIT

func getFactoCordVersion() (string, string) {
    return "${version}", "${githash}"
}
EOF

go build -o ./build/ServerManager main.go version.go
