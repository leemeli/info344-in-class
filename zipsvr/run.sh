#!/usr/bin/env bash
# chmod +x run.sh
docker rm -f zipsvr

docker run -d \
-p 443:443 \
--name zipsvr \
-v /Users/iguest/go/src/github.com/leemeli/info344-in-class/zipsvr/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
leemeli/zipsvr