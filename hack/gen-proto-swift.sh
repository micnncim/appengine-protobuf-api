#!/bin/bash

set -eu

proto_path=$1
proto_out=$2

swift_path=$proto_path/swift
mkdir $swift_path
for file in $(find $proto_path -type f)
do
    filename=$(basename $file)
    swift_file=$swift_path/$filename
    cp $file $swift_file
    sed -i -e '/import .*gogo\/protobuf.*/d' -e 's/ \[.*\]//g' $swift_file
done

mkdir -p $proto_out
protoc \
  --swift_out=$proto_out \
  --proto_path=$swift_path \
  --swift_opt=Visibility=Public \
  $swift_path/*.proto

echo "generated proto for swift"
echo $(find $proto_out -type f)

rm -rf $swift_path
