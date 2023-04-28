export BINDIR=.
curl -sf https://raw.githubusercontent.com/livebud/bud/main/install.sh | sh
mv bud bud-run
chmod +x ./bud-run
npm install
go mod tidy