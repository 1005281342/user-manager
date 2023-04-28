export BINDIR=.
export BINARY=bud-run
curl -sf https://raw.githubusercontent.com/livebud/bud/main/install.sh | sh
chmod +x ./bud-run
npm install
go mod tidy