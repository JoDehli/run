PKG="./lib"
BIN="./bin/dot"

if VERSION=$(git describe --tags); then
  echo "Building dot $VERSION..."
else
  echo "Build failed: unable to set version"
  exit
fi

LDFLAGS="-X github.com/lukecjohnson/dot/lib/utils.CurrentVersion=$VERSION"

if go build -ldflags "$LDFLAGS" -o "$BIN" "$PKG"; then
  echo "Build complete: $BIN"
else
  echo "Build failed: unable to compile go package"
fi
