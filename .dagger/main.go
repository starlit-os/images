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
		WithDirectory("/", source.Directory("system_files"))

	// Mount caches
	for _, cache := range m.Caches {
		container = container.
			WithMountedCache(cache.Path, dag.CacheVolume(cache.Name))
	}

	container = container.
		WithExec([]string{"mkdir", "-p", "/var/opt"}).
		WithExec([]string{"ln", "-s", "/var/opt", "/opt"})

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

	for _, opt := range m.OptFixes {
		container = container.
			WithExec([]string{"mv", "/var/opt/" + opt.Directory, "/lib/" + opt.Directory}).
			WithExec([]string{"rm", "/usr/bin/" + opt.Binary}).
			WithExec([]string{"ln", "-s", "/opt/" + opt.Directory + "/" + opt.BinaryTarget, "/usr/bin/" + opt.Binary})
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
		WithExec([]string{
			"sh",
			"-c",
			"echo 'import? \"/usr/share/ublue-os/just/99-lily.just\"' | tee -a /usr/share/ublue-os/justfile",
		}).
		WithExec([]string{"bootc", "container", "lint"})

	return container
}
