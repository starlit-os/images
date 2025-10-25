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
		WithReposEnabled(ctx, []string{"docker-ce-stable"}).
		WithDirectory(ctx, "system_files/bazzite", "/").
		WithCopr(ctx, "ublue-os/packages").
		WithDirectory(ctx, "system_files/shared", "/").
		WithPackages(ctx, []string{
			"docker-ce",
			"docker-ce-cli",
			"docker-compose-plugin",
			"docker-model-plugin",
			"ublue-os-luks",
		}).
		WithServices(ctx, []string{
			"docker.service",
			"docker.socket",
			"podman.socket",
			"podman-restart.service",
			"podman-auto-update.timer",
		})
}
