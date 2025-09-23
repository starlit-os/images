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
			"fwupd",
			"gdm",
			"gnome-{bluetooth,color-manager,control-center,disk-utility,session-wayland-session,settings-daemon,shell}",
			"nautilus",
			"plymouth{,-system-theme}",
			"systemd-{resolved,container,oomd}",
			"tuned-ppd",
			"wireguard-tools",
			"wl-clipboard",
			"xdg-desktop-portal-gnome",
			"xdg-user-dirs-gtk",
		})
}
