package main

import (
	"context"
	"dagger/bazzite/internal/dagger"
)

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

type OptFix struct {
	Directory    string
	Binary       string
	BinaryTarget string
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
	OptFixes    []OptFix
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

func (m *Bazzite) WithOptFix(ctx context.Context, directory string, binary string, binary_target string) *Bazzite {
	m.OptFixes = append(m.OptFixes, OptFix{
		Directory:    directory,
		Binary:       binary,
		BinaryTarget: binary_target,
	})
	return m
}
