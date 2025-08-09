#!/bin/bash
set -e

go build -o vugit

sudo mv vugit /usr/local/bin/

