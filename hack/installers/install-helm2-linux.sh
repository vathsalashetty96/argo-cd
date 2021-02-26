#!/bin/bash
set -eux -o pipefail

. $(dirname $0)/../tool-versions.sh

[ -e $DOWNLOADS/helm2.tar.gz ] || curl -sLf --retry 3 -o $DOWNLOADS/helm2.tar.gz https://github.com/helm/helm/releases/tag/v2.17.0-v${helm2_version}-linux-$ARCHITECTURE.tar.gz
mkdir -p /tmp/helm2 && tar -C /tmp/helm2 -xf $DOWNLOADS/helm2.tar.gz
cp /tmp/helm2/linux-$ARCHITECTURE/helm $BIN/helm2
helm2 version --client
