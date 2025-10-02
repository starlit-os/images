package main

import (
	"context"
	"dagger/bazzite/internal/dagger"
)

// Creates a Bazzite container
func (m *Bazzite) BazziteContainer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
	source_image string,
) *Bazzite {
	return m.From(ctx, source_image).
		WithRpmfusion(ctx).
		WithTerra(ctx).
		WithReposEnabled(ctx, []string{"docker-ce-stable"}).
		WithReposEnabled(ctx, []string{"warpdotdev"}).
		WithDirectory(ctx, "system_files/shared", "/").
		WithDirectory(ctx, "system_files/desktop", "/").
		WithDirectory(ctx, "system_files/bazzite", "/").
		WithCopr(ctx, "scottames/ghostty").
		WithCopr(ctx, "che/nerd-fonts").
		WithPackages(ctx, []string{
			"coolercontrol",
			"discord",
			"ghostty",
			"headsetcontrol",
			"liquidctl",
			"nerd-fonts",
			"openrgb",
			"docker-ce",
			"docker-ce-cli",
			"docker-compose-plugin",
			"docker-model-plugin",
			"warp-terminal",
		}).
		WithOptFix(ctx, "warpdotdev", "warp-terminal", "warp-terminal/warp").
		WithServices(ctx, []string{
			"docker.service",
			"docker.socket",
			"podman.socket",
			"podman-restart.service",
			"podman-auto-update.timer",
			"tailscaled.service",
		})
}
