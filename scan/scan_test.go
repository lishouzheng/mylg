package scan_test

import (
	"testing"

	"github.com/shouzhengli/mylg/cli"
	"github.com/shouzhengli/mylg/scan"
)

var (
	s      scan.Scan
	cfg, _ = cli.ReadDefaultConfig()
)

func TestIsCIDR(t *testing.T) {
	var err error
	s, err = scan.NewScan("8.8.8.0/24", cfg)
	if err == nil {
		t.Error("NewScan failed")
	}
	if !s.IsCIDR() {
		t.Error("IsCIDR failed")
	}
	s, err = scan.NewScan("8.8.8.8", cfg)
	if err != nil {
		t.Error("NewScan failed")
	}
	if s.IsCIDR() {
		t.Error("IsCIDR failed")
	}
}
