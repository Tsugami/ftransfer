package storage_provider_client

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type StorageProviderClientSFTP struct {
	client *sftp.Client
}

type StorageProviderClientConnection struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewStorageProviderClientSFTP(connection *StorageProviderClientConnection) (*StorageProviderClientSFTP, error) {
	config := &ssh.ClientConfig{
		User: connection.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(connection.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := net.JoinHostPort(connection.Host, strconv.Itoa(connection.Port))
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	client, err := sftp.NewClient(sshClient)
	if err != nil {
		sshClient.Close()
		return nil, fmt.Errorf("failed to create SFTP client: %w", err)
	}

	return &StorageProviderClientSFTP{client: client}, nil
}

func (s *StorageProviderClientSFTP) ListFiles(ctx context.Context, path string) ([]string, error) {
	files, err := s.client.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

func (s *StorageProviderClientSFTP) ReadFile(ctx context.Context, path string) (io.ReadCloser, error) {
	file, err := s.client.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return file, nil
}

func (s *StorageProviderClientSFTP) WriteFile(ctx context.Context, path string, data io.Reader) error {
	file, err := s.client.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (s *StorageProviderClientSFTP) MoveFile(ctx context.Context, sourcePath string, destinationPath string) error {
	err := s.client.Rename(sourcePath, destinationPath)
	if err != nil {
		return fmt.Errorf("failed to move file: %w", err)
	}
	return nil
}
