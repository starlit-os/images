package main

import (
	"context"
	"dagger/bazzite/internal/dagger"
)

// Creates a Bluefin container
func (m *Bazzite) BluefinContainer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
	source_image string,
) *Bazzite {
	return m.From(ctx, source_image).
		WithRpmfusion(ctx).
		WithTerra(ctx).
		WithReposEnabled(ctx, []string{"warpdotdev"}).
		WithDirectory(ctx, "system_files/shared", "/").
		WithDirectory(ctx, "system_files/desktop", "/").
		WithPackages(ctx, []string{
			"coolercontrol",
			"discord",
			"headsetcontrol",
			"liquidctl",
			"lutris",
			"openrgb",
			"podman-docker",
			"warp-terminal",
		}).
		WithOptFix(ctx, "warpdotdev", "warp-terminal", "warp-terminal/warp").
		WithServices(ctx, []string{
			"podman.socket",
			"podman-restart.service",
			"podman-auto-update.timer",
		})
}
