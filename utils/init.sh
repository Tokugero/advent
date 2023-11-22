#!/bin/bash

YEAR=2023
mkdir ./$YEAR
for i in {1..25}
do 
    mkdir -p ./2023/day$i
    cd ./2023/day$i
    echo "package main" > main.go
    go mod init day$i
    go mod tidy
    cd ../..
done

for i in $(ls $YEAR)
do
    go work use ./$YEAR/$i/
done