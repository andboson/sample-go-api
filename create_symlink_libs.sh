#!/bin/bash

#needed only for idea+glide pakage manager

echo "creating symlinks for 3rd-party golang libs"

APPDIRNAME=${PWD##*/}
cd ..
SRCDIRFULL=${pwd}
SRCDIR="src/"
cd vendor
VENDOR="../src/"$APPDIRNAME"/vendor"

for d in $VENDOR/*; do
   ln -sv ""$d
done

cd $APPDIRNAME

echo "symlinks created under " $SRCDIRFULL