DIST="./dist"

rm -rf "$DIST"
mkdir "$DIST"

if VERSION=$(git describe --tags --abbrev=0); then
  echo "Creating package for dot $VERSION..."
else
  echo "Failed to create package: unable to set version"
  exit
fi

PKG="$DIST/dot-$VERSION-macos-amd64.tar.gz"
BIN="./bin/dot"
README="./README.md"
LICENSE="./LICENSE"

if tar -czf "$PKG" "$BIN" "$README" "$LICENSE"; then
  echo "Package created: $PKG"
else
  echo "Failed to create package: unable to create compressed archive"
fi
