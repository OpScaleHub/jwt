# jwt
Portable JWT token validation 

- support for TOKEN read as cli argument
- support for TOKEN read from stdin


## How To Install
```bash
curl --silent https://raw.githubusercontent.com/OpScaleHub/jwt/main/get-jwt.sh | sudo bash
```

## Usage

```bash
 $ jwt ${TOKEN}
      [#] HEADER:ALGORITHM & TOKEN TYPE
alg:   [ RS256 ]
kid:   [ 9f0106c270df0da4d406c6520b76c9157887b493 ]
      [#] PAYLOAD:DATA
iss:   [ https://idp.opscale.ir/dex ]
sub:   [ 45499github ]
aud:   [ kubernetes ]
exp:   [ 2022-05-11 17:59:11 +0200 CEST ]
iat:   [ 2022-05-10 17:59:11 +0200 CEST ]
nonce: [ dHz0utfydrof6XYXJtXJi8X6xSx5Rzpw1eSlDkjEIW0 ]
      [#] VERIFY SIGNATURE
#ToDo
(⎈ |opscale:argocd) ~/Projects/jwt (main *)
 $ echo ${TOKEN} | jwt 
      [#] HEADER:ALGORITHM & TOKEN TYPE
alg:   [ RS256 ]
kid:   [ 9f0106c270df0da4d406c6520b76c9157887b493 ]
      [#] PAYLOAD:DATA
iss:   [ https://idp.opscale.ir/dex ]
sub:   [ 45499github ]
aud:   [ kubernetes ]
exp:   [ 2022-05-11 17:59:11 +0200 CEST ]
iat:   [ 2022-05-10 17:59:11 +0200 CEST ]
nonce: [ dHz0utfydrof6XYXJtXJi8X6xSx5Rzpw1eSlDkjEIW0 ]
      [#] VERIFY SIGNATURE
#ToDo
```