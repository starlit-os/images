# Bazzite (Starlit Edition)

Personal tweaked version of [Bazzite](https://bazzite.gg)

## Additions

### Packages

Preinstalls some packages for convenience.

* [Ghostty](https://ghosty.org)
* [OpenRGB](https://openrgb.org) for lighting control
* [Discord](https://discord.com) to avoid Discord RPC issues with flatpak
* [HeadsetControl](https://github.com/Sapd/HeadsetControl) for battery info from my wireless headset
* [CoolerControl](https://docs.coolercontrol.org) for fan and AIO control

### Other stuff

* Preloads iptable_nat kernel module to let me work with [Dagger](https://dagger.io/)

## Installation

Simply install Bazzite in your preferred way and rebase to this image
```bash
rpm-ostree rebase ostree-image-signed:docker://ghcr.io/starlit-os/bazzite:latest
```

## Verification

The image is signed with cosign keyless signing. The signature can be verified using 

```bash
cosign verify ghcr.io/starlit-os/bazzite --certificate-oidc-issuer=https://token.actions.githubusercontent.com --certificate-identity https://github.com/starlit-os/bazzite/.github/workflows/build.yml@refs/heads/main
```
