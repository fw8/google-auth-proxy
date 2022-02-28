# Google auth proxy

Based on <https://github.com/LautaroJayat/go_auth_proxy_example>

and <https://gist.github.com/salrashid123/23b8a666371cd5ccd3f8bcd5ca4a0105>

This repo is meant to be a playground for learning "Go". Don't expect anything production-ready or very sophisticated.
I just want to explore "go".

## Build

```bash
docker build -t google-auth-proxy . --build-arg app_env=production
```

## Start

```bash
docker run -it -p 8080:8080 -e HOST=some-cloud-run-service -e SCHEME=https -e GOOGLE_APPLICATION_CREDENTIALS=/creds.json -v /path/to/gsa.json:/gsa.json  google-auth-proxy
```
