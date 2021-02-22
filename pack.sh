#!/usr/bin/env bash

go build -o alfred-strlen
zip alfred-strlen.alfredworkflow ./alfred-strlen ./info.plist ./icon.png
rm ./alfred-strlen
