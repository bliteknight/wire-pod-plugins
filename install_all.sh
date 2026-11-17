POD_DIR=~/wire-pod
PLUGIN_DIR="$POD_DIR/chipper/plugins"
REPO_URL="https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/refs/heads/main"

echo "##################################"
echo "###  Installing all plugins"
echo "##################################"

# use the same Go toolchain wire-pod is built with (sets $GO, reads GO_VERSION)
source "$POD_DIR/goenv.sh"
if [[ ! -x "$GO" ]]; then
    getGo   # downloads the right Go if it isn't installed yet
fi
echo "Building plugins with: $("$GO" version)"

# plugins must be built with the same Go version and cgo setting as the wire-pod binary
export CGO_ENABLED=1

for name in greetings simonsays whatdate; do
    echo "--- $name"
    wget -q -O "$PLUGIN_DIR/$name.go" "$REPO_URL/$name/$name.go"
    "$GO" build -buildmode=plugin -o "$PLUGIN_DIR/$name.so" "$PLUGIN_DIR/$name.go"
done

sudo systemctl restart wire-pod
