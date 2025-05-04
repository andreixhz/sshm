package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"github.com/andreixhz/sshm/models"
)

var dataFile = filepath.Join(os.Getenv("HOME"), ".sshmanager_hosts.json")

func LoadHosts() ([]models.Host, error) {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return []models.Host{}, nil
	}
	var hosts []models.Host
	err = json.Unmarshal(file, &hosts)
	return hosts, err
}

func SaveHosts(hosts []models.Host) error {
	data, err := json.MarshalIndent(hosts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}
