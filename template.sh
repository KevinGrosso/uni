#!/bin/bash
mkdir "$1"
cd "$1"
touch "$1.go"
echo $'package main \n\nimport (\n\t"fmt"\n)\n\nfunc main() {\n\n}' >> "$1.go"