#!/bin/bash

if [ -z "$1" ]; then
  echo "usage: $0 <app name>"
  exit 1
fi

(set -x
  rm -f *.{txt,bash,fish,zsh,ps1}
  go build -ldflags "-X main.name=$1"
)

for name in bash fish zsh powershell; do
  ext=$name
  if [ "$ext" = "powershell" ]; then
    ext="ps1"
  fi
  (set -x;
    $1 completion $name > $1.$ext
    ./tree completion $name > tree.$ext
  )
done