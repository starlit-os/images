// A generated module for Bazzite functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/bazzite/internal/dagger"
)

type Bazzite struct{}

func MountCaches(container *dagger.Container) *dagger.Container {
	return container.
		WithMountedCache("/var/roothome", dag.CacheVolume("var-roothome")).
		WithMountedCache("/var/cache", dag.CacheVolume("var-cache")).
		WithMountedCache("/var/log", dag.CacheVolume("var-log")).
		WithMountedCache("/var/lib/dnf", dag.CacheVolume("var-lib-dnf"))
}

func UnmountCaches(container *dagger.Container) *dagger.Container {
	return container.
		WithoutMount("/var/roothome").
		WithoutMount("/var/cache").
		WithoutMount("/var/log").
		WithoutMount("/var/lib/dnf")
}

func EnableServices(container *dagger.Container, services []string) *dagger.Container {
	return container.
		WithExec(append([]string{"systemctl", "enable"}, services...))
}

func Commit(container *dagger.Container) *dagger.Container {
	return container.
		WithExec([]string{"ostree", "container", "commit"})
}

func Lint(container *dagger.Container) *dagger.Container {
	return container.
		WithExec([]string{"bootc", "container", "lint"})
}

func EnableCopr(container *dagger.Container, copr string) *dagger.Container {
	return container.
		WithExec([]string{"dnf5", "copr", "enable", "-y", copr})
}

func DisableCopr(container *dagger.Container, copr string) *dagger.Container {
	return container.
		WithExec([]string{"dnf5", "copr", "disable", "-y", copr})
}

func EnableRpmFusion(container *dagger.Container) *dagger.Container {
	return container.
		WithExec([]string{"dnf5", "config-manager", "setopt", "rpmfusion-free.enabled=1", "rpmfusion-free-updates.enabled=1", "rpmfusion-nonfree.enabled=1", "rpmfusion-nonfree-updates.enabled=1"})
}

func DisableRpmFusion(container *dagger.Container) *dagger.Container {
	return container.
		WithExec([]string{"dnf5", "config-manager", "setopt", "rpmfusion-free.enabled=0", "rpmfusion-free-updates.enabled=0", "rpmfusion-nonfree.enabled=0", "rpmfusion-nonfree-updates.enabled=0"})
}

func InstallPackages(container *dagger.Container, packages []string) *dagger.Container {
	return container.
		WithExec(append([]string{"dnf5", "install", "-y"}, packages...))
}

func EnableTerra(container *dagger.Container) *dagger.Container {
	return container.
		WithExec([]string{"dnf5", "config-manager", "setopt", "terra.enabled=1"})
}

func DisableTerra(container *dagger.Container) *dagger.Container {
	return container.
		WithExec([]string{"dnf5", "config-manager", "setopt", "terra.enabled=0"})
}

// Returns a container that echoes whatever string argument is provided
func (m *Bazzite) BazziteContainer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Container {
	container := dag.Container().
		From("ghcr.io/ublue-os/bazzite-gnome:stable@sha256:23c02860f424869463e363a6e96948918c1036b6e6474b0420186cfb1fd31c68")
	container = MountCaches(container)
	container = EnableRpmFusion(container)
	container = EnableCopr(container, "scottames/ghostty")
	container = EnableTerra(container)
	container = InstallPackages(container, []string{"discord", "ghostty", "headsetcontrol", "openrgb", "podman-docker", "liquidctl", "coolercontrol"})
	container = DisableTerra(container)
	container = DisableCopr(container, "scottames/ghostty")
	container = DisableRpmFusion(container)
    container = EnableServices(container, []string{"podman.socket", "podman-restart.service", "podman-auto-update.timer"})
	container = Commit(container)
	container = UnmountCaches(container)
	return Lint(container)
}
