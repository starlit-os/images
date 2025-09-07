package main

import (
	"context"
	"dagger/bazzite/internal/dagger"
)

// Creates a Cayo container
func (m *Bazzite) CayoContainer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
	source_image string,
) *Bazzite {
	return m.From(ctx, source_image).
		WithDnf(ctx, "dnf").
		WithCopr(ctx, "ublue-os/packages").
		WithDirectory(ctx, "system_files/shared", "/").
		WithPackages(ctx, []string{
			"podman-docker",
			"ublue-os-luks",
			"ublue-os-just"
		}).
		WithServices(ctx, []string{
			"podman.socket",
			"podman-restart.service",
			"podman-auto-update.timer",
		})
}
