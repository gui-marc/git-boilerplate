#!/bin/bash

# Determine the latest release version
VERSION=$(curl -s https://api.github.com/repos/gui_marc/git-boilerplate/releases/latest | grep "tag_name" | awk '{print substr($2, 2, length($2)-3)}')

# Determine the operating system and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "${OS}" == "darwin" ]; then
    OS="macos"
elif [ "${OS}" == "linux" ]; then
    if [ "${ARCH}" == "x86_64" ]; then
        ARCH="amd64"
    elif [ "${ARCH}" == "aarch64" ]; then
        ARCH="arm64"
    fi
fi

# Download and install the binary file
curl -L -o /usr/local/bin/git-boilerplate https://github.com/gui_marc/git-boilerplate/releases/download/${VERSION}/git-boilerplate_${OS}_${ARCH}

chmod +x /usr/local/bin/git-boilerplate

echo "git-boilerplate has been installed to /usr/local/bin"
