#!/bin/bash
set -eux -o pipefail
. $(dirname $0)/../tool-versions.sh
[ -e $DOWNLOADS/swagger ] || curl -sLf --retry 3 -o $DOWNLOADS/swagger https://github.com/vathsalashetty25/go-swagger/raw/master/deploybuild/swagger_linux_ppc64le

cp $DOWNLOADS/swagger $BIN
ls -l $BIN/swagger
chmod +x $BIN/swagger
swagger version
