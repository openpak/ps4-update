# Deploy

Not yet deployed.

## Run

```sh
cp example.env .env
set -a; . ./.env; set +a
GOWORK=off go run .
```

Quadlet: `deploy/openpak-ps4-update.container`; image `ghcr.io/openpak/ps4-update`; listens on `:8540`.

## Environment

| Variable | Default | |
| --- | --- | --- |
| `PS4_UPDATE_ADDR` | `:8540` | Listen address |
| `PS4_UPDATE_VERSION` | `00.000.000` | Advertised system version |
| `PS4_UPDATE_LABEL` | `0.00` | Advertised label |
| `PS4_UPDATE_SDK_VERSION` | the version | Advertised SDK version |
| `PS4_UPDATE_IMAGE_URL` | empty | PUP URL; empty omits the image blocks |

## DNS

OpenPak DNS must point every `<d|f|h><region>01.ps4.update.playstation.net` (regions au br cn eu hk jp kr mx
ru sa tw uk us) at this service on port 80, for example through Traefik to `:8540`.
