#!/usr/bin/env bash
MAX_ICON=100

type convert >/dev/null 2>&1 || { sudo apt install imagemagick; }

# generate icons
for i in $(seq 0 $MAX_ICON)
do
  convert -size 32x32 xc:transparent -font Palatino-Bold   -pointsize 30 -fill '#b6d0fa' -draw "text 0,28 '$i'"  images/$i.ico
  convert -size 32x32 xc:transparent -font Palatino-Bold   -pointsize 30 -fill '#b6d0fa' -draw "text 0,28 '$i'"  images/$i.png
done

# check go installed
if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

# check 2goarray installed
if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    go get github.com/noahhai/2goarray
    if [ $? -ne 0 ]; then
        echo Failure executing go get github.com/noahhai/2goarray
        exit
    fi
fi


# generate go binary data of icons - linux/darwin
OUTPUT=iconunix.go
echo Generating Go Binary Data file: Unix/Darwin
echo "//+build linux darwin" > $OUTPUT
echo >> $OUTPUT
for i in $(seq 0 $MAX_ICON)
do
  if [ $i -eq 0 ]; then
    cat "images/$i.png" | $GOPATH/bin/2goarray Data$i main >> $OUTPUT
  else
    cat "images/$i.png" | $GOPATH/bin/2goarray Data$i >> $OUTPUT
  fi
done
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi
echo Finished Unix/Dawin

# generate go binary data of icons - Windows
OUTPUT=iconwin.go
echo Generating Go Binary Data file: Windows
echo "//+build windows" > $OUTPUT
echo >> $OUTPUT
for i in $(seq 0 $MAX_ICON)
do
    if [ $i -eq 0 ]; then
    cat "images/$i.ico" | $GOPATH/bin/2goarray Data$i main >> $OUTPUT
  else
    cat "images/$i.ico" | $GOPATH/bin/2goarray Data$i >> $OUTPUT
  fi
done
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi
echo Finished Windows
