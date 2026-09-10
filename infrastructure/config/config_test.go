package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/subnetworks"
	"github.com/nonsense-project/nonsense/v2/domain/dagconfig"

	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"
)

func TestCreateDefaultConfigFile(t *testing.T) {
	// find out where the sample config lives
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("Failed finding config file path")
	}
	sampleConfigFile := filepath.Join(filepath.Dir(path), "sample-nonsensed.conf")

	// Setup a temporary directory
	tmpDir, err := ioutil.TempDir("", "nonsensed")
	if err != nil {
		t.Fatalf("Failed creating a temporary directory: %v", err)
	}
	testpath := filepath.Join(tmpDir, "test.conf")

	// copy config file to location of nonsensed binary
	data, err := ioutil.ReadFile(sampleConfigFile)
	if err != nil {
		t.Fatalf("Failed reading sample config file: %v", err)
	}
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatalf("Failed obtaining app path: %v", err)
	}
	tmpConfigFile := filepath.Join(appPath, "sample-nonsensed.conf")
	err = ioutil.WriteFile(tmpConfigFile, data, 0644)
	if err != nil {
		t.Fatalf("Failed copying sample config file: %v", err)
	}

	// Clean-up
	defer func() {
		os.Remove(testpath)
		os.Remove(tmpConfigFile)
		os.Remove(tmpDir)
	}()

	err = createDefaultConfigFile(testpath)
	if err != nil {
		t.Fatalf("Failed to create a default config file: %v", err)
	}
}

func TestApplyDefaultMainnetPeers(t *testing.T) {
	tests := []struct {
		name         string
		params       *dagconfig.Params
		addPeers     []string
		connectPeers []string
		expected     []string
	}{
		{name: "fresh mainnet", params: &dagconfig.MainnetParams, expected: mainnetBootstrapPeers},
		{name: "explicit addpeer", params: &dagconfig.MainnetParams, addPeers: []string{"127.0.0.1:1"}, expected: []string{"127.0.0.1:1"}},
		{name: "explicit connect", params: &dagconfig.MainnetParams, connectPeers: []string{"127.0.0.1:2"}},
		{name: "testnet", params: &dagconfig.TestnetParams},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := DefaultConfig()
			cfg.NetworkFlags.ActiveNetParams = test.params
			cfg.AddPeers = append([]string(nil), test.addPeers...)
			cfg.ConnectPeers = append([]string(nil), test.connectPeers...)
			applyDefaultMainnetPeers(cfg)

			if strings.Join(cfg.AddPeers, ",") != strings.Join(test.expected, ",") {
				t.Fatalf("unexpected addpeer list: got %v, expected %v", cfg.AddPeers, test.expected)
			}
			if strings.Join(cfg.ConnectPeers, ",") != strings.Join(test.connectPeers, ",") {
				t.Fatalf("connect list was changed: got %v, expected %v", cfg.ConnectPeers, test.connectPeers)
			}
		})
	}
}

// TestConstants makes sure that all constants hard-coded into the help text were not modified.
func TestConstants(t *testing.T) {
	zero := externalapi.DomainSubnetworkID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if subnetworks.SubnetworkIDNative != zero {
		t.Errorf("subnetworks.SubnetworkIDNative value was changed from 0, therefore you probably need to update the help text for SubnetworkID")
	}
	one := externalapi.DomainSubnetworkID{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if subnetworks.SubnetworkIDCoinbase != one {
		t.Errorf("subnetworks.SubnetworkIDCoinbase value was changed from 1, therefore you probably need to update the help text for SubnetworkID")
	}
	two := externalapi.DomainSubnetworkID{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if subnetworks.SubnetworkIDRegistry != two {
		t.Errorf("subnetworks.SubnetworkIDRegistry value was changed from 2, therefore you probably need to update the help text for SubnetworkID")
	}
}
