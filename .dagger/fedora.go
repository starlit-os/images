package main

import (
	"context"
	"dagger/bazzite/internal/dagger"
)

// Creates a Fedora bootc container
func (m *Bazzite) FedoraContainer(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
	source_image string,
) *Bazzite {
	return m.From(ctx, source_image).
		WithJust(ctx, false).
		WithPackages(ctx, []string{
			"gnome-control-center",
			"gnome-session",
			"gnome-shell",
			"nautilus",
		})
}
