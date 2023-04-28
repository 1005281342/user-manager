if test -f "bud-run"; then
  echo "File bud-run exists."
else
  echo "File bud-run does not exist."
  rm -rf bud
  export BINDIR=.
  curl -sf https://raw.githubusercontent.com/livebud/bud/main/install.sh | sh
  mv bud bud-run
  chmod +x ./bud-run
fi

npm install
go mod tidy