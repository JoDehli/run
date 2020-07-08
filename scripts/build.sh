if VERSION=$(git describe --tags --abbrev=0); then
  echo "Building run $VERSION..."
else
  echo "Build failed: unable to set version"
  exit
fi

LDFLAGS="-X main.currentVersion=$VERSION"
BIN="./bin/run"

if go build -ldflags "$LDFLAGS" -o "$BIN"; then
  echo "Build complete: $BIN"
else
  echo "Build failed: unable to compile go package"
fi
