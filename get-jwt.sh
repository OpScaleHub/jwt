#!/bin/bash
RELEASE_TAG=$(http https://api.github.com/repos/OpScaleHub/jwt/releases/latest | jq -r '.name')
wget -qO- https://github.com/OpScaleHub/jwt/releases/download/${RELEASE_TAG}/jwt.tgz | tar -xvz
install jwt /usr/local/bin/jwt
rm -f jwt