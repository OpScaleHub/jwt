#!/bin/bash
curl -L "https://github.com/OpScaleHub/jwt/releases/download/$(curl --silent https://api.github.com/repos/OpScaleHub/jwt/releases/latest | jq -r '.name')/jwt.tgz" | tar -xvz
install jwt /usr/local/bin/jwt
rm -f jwt