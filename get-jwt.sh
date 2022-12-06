#!/bin/bash
########################################
### installer script [Gnu/Linux]     ###
########################################

```bash
curl --output jwt -L "https://github.com/OpScaleHub/jwt/releases/download/$(curl --silent https://api.github.com/repos/OpScaleHub/jwt/releases/latest | jq -r '.name')/jwt"
install jwt /usr/local/bin/jwt
rm -f jwt
```