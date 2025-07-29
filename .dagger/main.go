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
	"fmt"
)

func New() *Bazzite {
	return &Bazzite{
		Coprs: []string{},
		Repos: []string{},
		Auth: RegistryAuth{
			Registry: "",
			Username: "",
			Password: nil,
		},
		Labels:      []ContainerLabel{},
		Services:    []string{},
		Tags:        []string{},
		Directories: []ContainerDirectory{},
		Caches: []Cache{
			{
				Path: "/var/roothome",
				Name: "var-roothome",
			},
			{
				Path: "/var/cache",
				Name: "var-cache",
			},
			{
				Path: "/var/log",
				Name: "var-log",
			},
			{
				Path: "/var/tmp",
				Name: "var-tmp",
			},
			{
				Path: "/var/lib/dnf",
				Name: "var-lib-dnf",
			},
		},
	}
}

// Publishes the Bazzite container to a registry
func (m *Bazzite) Publish(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
	// Registry
	registry string,
	// Image
	image string,
) ([]string, error) {
	container := m.Build(ctx, source).
		WithRegistryAuth(m.Auth.Registry, m.Auth.Username, m.Auth.Password)

	addr := []string{}

	for _, tag := range m.Tags {
		a, err := container.Publish(ctx, fmt.Sprintf("%s/%s:%s", registry, image, tag))
		if err != nil {
			return addr, err
		}
		addr = append(addr, a)
	}

	return addr, nil
}

func (m *Bazzite) Build(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Container {
	container := dag.Container().
		From(m.Source).
		WithDirectory("/", source.Directory("system_files")).
		WithExec([]string{
			"echo",
			"'import? \"/usr/share/ublue-os/just/70-lily.just\"'",
			">>",
			"/usr/share/ublue-os/justfile"})

	// Mount caches
	for _, cache := range m.Caches {
		container = container.
			WithMountedCache(cache.Path, dag.CacheVolume(cache.Name))
	}

	// Enable repositories
	for _, repo := range m.Repos {
		container = container.
			WithExec([]string{"dnf5", "config-manager", "setopt", fmt.Sprintf("%s.enabled=1", repo)})
	}

	// Enable Copr repositories
	for _, copr := range m.Coprs {
		container = container.
			WithExec([]string{"dnf5", "-y", "copr", "enable", copr})
	}

	// Install packages
	container = container.
		WithExec(append([]string{"dnf5", "-y", "install"}, m.Packages...))

	// Disable repositories
	for _, repo := range m.Repos {
		container = container.
			WithExec([]string{"dnf5", "config-manager", "setopt", fmt.Sprintf("%s.enabled=0", repo)})
	}

	// Disable Copr repositories
	for _, copr := range m.Coprs {
		container = container.
			WithExec([]string{"dnf5", "-y", "copr", "disable", copr})
	}

	container = container.
		WithExec(append([]string{"systemctl", "enable"}, m.Services...)).
		WithExec([]string{"ostree", "container", "commit"})

	// Unmount caches
	for _, cache := range m.Caches {
		container = container.
			WithoutMount(cache.Path)
	}

	container = container.
		WithExec([]string{"bootc", "container", "lint"})

	return container
}

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
		WithDirectory(ctx, "system_files", "/").
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
			"podman-docker",
		}).
		WithServices(ctx, []string{
			"podman.socket",
			"podman-restart.service",
			"podman-auto-update.timer",
		})
}
