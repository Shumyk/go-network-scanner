#!/bin/bash

WEBUI_VERSION="2.4.2"
GO_WEBUI_VERSION=v2.4.2-1.0
MODULE=github.com/webui-dev/go-webui/v2

# Verify GOPATH.
if [ -z "${GOPATH}" ]; then
    echo "Warning: GOPATH is not set."
    GO_PATH="$HOME/go"
    echo -e "Trying to use $GO_PATH instead.\n"
else
    GO_PATH="$GOPATH"
fi

# Verify that module package is installed.
MODULE_PATH="$GO_PATH/pkg/mod/${MODULE}@${GO_WEBUI_VERSION}"
if [ ! -d "$MODULE_PATH" ]; then
    echo "Error: \`$MODULE_PATH\` does not exist in GOPATH."
    echo "Make sure to run \`go get $MODULE@${GO_WEBUI_VERSION}\` first."
    exit 1
fi

CURRENT_DIR=$(pwd)

# Make sure the go modules directory is writable for the current user.
chmod +w "$MODULE_PATH"
cd "$MODULE_PATH" || exit 1
echo "In WebUI directory: $(pwd)"

# Remove the existing WebUI C library.
if [ -d "webui" ]; then
    echo "Removing the existing WebUI C library..."
    rm -fr webui
fi

case "${XGOOS}" in
linux*)
    case "${XGOARCH}" in
    amd64 | x86_64)
        ARCHIVE="webui-linux-gcc-x64.zip"
        ;;
    aarch64 | arm64)
        ARCHIVE="webui-linux-gcc-arm64.zip"
        ;;
    arm*)
        ARCHIVE="webui-linux-gcc-arm.zip"
        ;;
    *)
        echo "The setup script currently does not support ${XGOARCH} architectures on ${XGOOS}."
        exit 1
        ;;
    esac
    ;;
darwin*)
    case "${XGOARCH}" in
    amd64 | x86_64)
        ARCHIVE="webui-macos-clang-x64.zip"
        ;;
    arm64)
        ARCHIVE="webui-macos-clang-arm64.zip"
        ;;
    *)
        echo "The setup script currently does not support ${XGOARCH} architectures on ${XGOOS}."
        exit 1
        ;;
    esac
    ;;
windows*)
    case "${XGOARCH}" in
    amd64 | x86_64)
        ARCHIVE="webui-windows-gcc-x64.zip"
        ;;
    *)
        echo "The setup script currently does not support ${XGOARCH} architectures on ${XGOOS}."
        exit 1
        ;;
    esac
    ;;
*)
    echo "The setup script currently does not support ${XGOOS}."
    exit 1
    ;;
esac

# Download and extract the archive.
echo "Downloading WebUI C library ${WEBUI_VERSION}..."
curl -L "https://github.com/webui-dev/webui/releases/download/${WEBUI_VERSION}/${ARCHIVE}" -o "${ARCHIVE}"

echo "Extracting WebUI C library..."
OUTPUT="webui"
ARCHIVE_DIR="${ARCHIVE%.zip}"
unzip -o "${ARCHIVE}"
mv "${ARCHIVE_DIR}" "${OUTPUT}"

echo ""

# Clean up.
echo "Cleaning up..."
rm -f "${ARCHIVE}"
rm -fr "${OUTPUT:?}/${ARCHIVE_DIR:?}"

echo "Done."
echo ""

cd "${CURRENT_DIR}" || exit 1
