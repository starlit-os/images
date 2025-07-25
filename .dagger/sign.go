package main

import (
	"fmt"
	"strings"
)

func (m *Bazzite) sign(addresses []string) (string, error) {
	digest, err := m.extractDigest(addresses)
	if err != nil {
		return "", err
	}

	return digest, nil
}

// ExtractHashFromFirstAddress extracts the registry/owner/image@sha256:hash from the first published address
func (m *Bazzite) extractDigest(addresses []string) (string, error) {
	if len(addresses) == 0 {
		return "", fmt.Errorf("no addresses provided")
	}

	addr := addresses[0]
	// Find the @ symbol that precedes the hash
	atIndex := strings.LastIndex(addr, "@sha256:")
	if atIndex == -1 {
		return "", fmt.Errorf("no SHA256 hash found in address: %s", addr)
	}

	// Extract registry/owner/image part (everything before :tag)
	colonIndex := strings.LastIndex(addr[:atIndex], ":")
	if colonIndex == -1 {
		return "", fmt.Errorf("no tag separator found in address: %s", addr)
	}

	registryOwnerImage := addr[:colonIndex]
	hashPart := addr[atIndex:] // includes "@sha256:"

	return registryOwnerImage + hashPart, nil
}
