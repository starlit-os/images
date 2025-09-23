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
			"gdm",
			"gnome-bluetooth",
			"gnome-color-manager",
			"gnome-control-center",
			"gnome-disk-utility",
			"gnome-session-wayland-session",
			"gnome-settings-daemon",
			"gnome-shell",
			"nautilus",
			"plymouth",
			"plymouth-system-theme",
			"systemd-container",
			"tuned-ppd",
			"wireguard-tools",
			"wl-clipboard",
			"xdg-desktop-portal-gnome",
			"xdg-user-dirs-gtk",
		})
}
