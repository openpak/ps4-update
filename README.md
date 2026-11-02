# PS4 Update List Server

OpenPak's PlayStation 4 system-update list server.

| | |
| --- | --- |
| Network | PlayStation 4 |
| Role | Serves `ps4-updatelist.xml` so consoles stay on their current firmware |
| Depends on | Nothing (Go standard library only) |
| Used by | PlayStation 4 consoles via OpenPak DNS |
| Image | `ghcr.io/openpak/ps4-update` |
| Licence | AGPL-3.0-only |

## What it serves

| Endpoint | What |
| --- | --- |
| `GET /update/ps4/list/{region}/ps4-updatelist.xml` | Update list for any region code, in Sony's format |
| `GET /healthz` | Health check |

## Run

```sh
cp example.env .env
set -a; . ./.env; set +a
GOWORK=off go run .
```

Tests: `GOWORK=off go test ./...`

## Docs

- [Protocol](docs/protocol.md)
- [Deploy](docs/deploy.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
