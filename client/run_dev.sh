set -a allexport
source ./config/dev.env
set +a allexport

go run cmd/app/main.go