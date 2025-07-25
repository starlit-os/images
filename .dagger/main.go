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

type ContainerDirectory struct {
	HostPath      string
	ContainerPath string
}

type ContainerLabel struct {
	Key   string
	Value string
}

type RegistryAuth struct {
	Registry string
	Username string
	Password *dagger.Secret
}

type Cache struct {
	Path string
	Name string
}

type Bazzite struct {
	Source      string
	Coprs       []string
	Repos       []string
	Auth        RegistryAuth
	Labels      []ContainerLabel
	Services    []string
	Tags        []string
	Caches      []Cache
	Packages    []string
	Directories []ContainerDirectory
}

func (m *Bazzite) From(ctx context.Context, source string) *Bazzite {
	m.Source = source
	return m
}

func (m *Bazzite) WithServices(ctx context.Context, services []string) *Bazzite {
	m.Services = append(m.Services, services...)
	return m
}

// Enables the specified COPR repository in the container.
func (m *Bazzite) WithCopr(ctx context.Context, copr string) *Bazzite {
	m.Coprs = append(m.Coprs, copr)
	return m
}

// Enables RPM Fusion repositories in the container.
func (m *Bazzite) WithRpmfusion(ctx context.Context) *Bazzite {
	return m.WithReposEnabled(ctx, []string{"rpmfusion-free", "rpmfusion-free-updates", "rpmfusion-nonfree", "rpmfusion-nonfree-updates"})
}

// Enables Terra repositories in the container.
func (m *Bazzite) WithTerra(ctx context.Context) *Bazzite {
	return m.WithReposEnabled(ctx, []string{"terra"})
}

// Enables the specified repositories in the container.
func (m *Bazzite) WithReposEnabled(ctx context.Context, repos []string) *Bazzite {
	m.Repos = append(m.Repos, repos...)
	return m
}

// Adds the specified tags to the Bazzite container.
func (m *Bazzite) WithTags(ctx context.Context, tags []string) *Bazzite {
	m.Tags = append(m.Tags, tags...)
	return m
}

func (m *Bazzite) WithPackages(ctx context.Context, packages []string) *Bazzite {
	m.Packages = append(m.Packages, packages...)
	return m
}

// Adds a label to the Bazzite container
func (m *Bazzite) WithLabel(ctx context.Context, key string, value string) *Bazzite {
	m.Labels = append(m.Labels, ContainerLabel{Key: key, Value: value})
	return m
}

// Adds a directory to the Bazzite container
func (m *Bazzite) WithDirectory(ctx context.Context, hostPath string, containerPath string) *Bazzite {
	m.Directories = append(m.Directories, ContainerDirectory{HostPath: hostPath, ContainerPath: containerPath})
	return m
}

// Sets the registry authentication for publishing the Bazzite container
func (m *Bazzite) WithRegistryAuth(ctx context.Context, registry string, username string, password *dagger.Secret) *Bazzite {
	m.Auth = RegistryAuth{
		Registry: registry,
		Username: username,
		Password: password,
	}
	return m
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
	// +optional
	version string,
	// +optional
	hash string,
) *Bazzite {

	if version == "" {
		version = "stable"
	}

	if hash != "" {
		hash = "@sha256:" + hash
	}

	return m.From(ctx, "ghcr.io/ublue-os/bazzite-gnome:"+version+hash).
		WithRpmfusion(ctx).
		WithTerra(ctx).
		WithDirectory(ctx, "system_files", "/").
		WithCopr(ctx, "scottames/ghostty").
		WithPackages(ctx, []string{"discord", "ghostty", "headsetcontrol", "openrgb", "podman-docker", "liquidctl", "coolercontrol"}).
		WithServices(ctx, []string{"podman.socket", "podman-restart.service", "podman-auto-update.timer"})
}
