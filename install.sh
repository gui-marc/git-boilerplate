#!/usr/bin/env sh

set -eu
printf '\n'

# Determine the latest release version
VERSION=$(curl -s https://api.github.com/repos/gui-marc/git-boilerplate/releases/latest | grep "tag_name" | awk '{print substr($2, 2, length($2)-3)}')

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

echo https://github.com/gui-marc/git-boilerplate/releases/download/${VERSION}/git-boilerplate-${OS}-${ARCH}

rm -rf ~/.git-boilerplate
mkdir ~/.git-boilerplate

# Download and install the binary file
curl -L -o ~/.git-boilerplate/git-boilerplate https://github.com/gui-marc/git-boilerplate/releases/download/${VERSION}/git-boilerplate-${OS}-${ARCH}

chmod +x ~/.git-boilerplate/git-boilerplate

echo "git-boilerplate has been installed to ~/.git-boilerplate"

# Write the path to the .bashrc file
echo "PATH=$PATH:$HOME/.git-boilerplate" >> ~/.bashrc