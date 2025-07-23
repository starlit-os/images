#!/bin/bash

set -ouex pipefail

### Install packages

# Enable RPM Fusion repositories
dnf5 config-manager setopt rpmfusion-free.enabled=1 rpmfusion-free-updates.enabled=1 \
    rpmfusion-nonfree.enabled=1 rpmfusion-nonfree-updates.enabled=1

# Enable Ghostty copr
dnf5 copr enable -y scottames/ghostty

dnf5 install -y \
    discord \
    ghostty \
    headsetcontrol \
    openrgb \
    podman-docker

# Disable Ghostty copr
dnf5 copr disable scottames/ghostty

# Disable RPM Fusion repositories
dnf5 config-manager setopt rpmfusion-free.enabled=0 rpmfusion-free-updates.enabled=0 \
    rpmfusion-nonfree.enabled=0 rpmfusion-nonfree-updates.enabled=0

### Enable services
systemctl enable podman.socket
systemctl enable podman-restart.service
systemctl enable podman-auto-update.timer
