LIB="./lib"
BIN="./bin/run"

if VERSION=$(git describe --tags --abbrev=0); then
  echo "Building run $VERSION..."
else
  echo "Build failed: unable to set version"
  exit
fi

LDFLAGS="-X github.com/lukecjohnson/run/lib/utils.CurrentVersion=$VERSION"

if go build -ldflags "$LDFLAGS" -o "$BIN" "$LIB"; then
  echo "Build complete: $BIN"
else
  echo "Build failed: unable to compile go package"
fi
