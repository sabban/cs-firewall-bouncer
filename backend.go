package main

import (
	"fmt"
	"runtime"

	"github.com/crowdsecurity/crowdsec/pkg/models"
	log "github.com/sirupsen/logrus"
)

type backend interface {
	Init() error
	ShutDown() error
	Add(*models.Decision) error
	Delete(*models.Decision) error
	Commit() error
}

type backendCTX struct {
	firewall backend
}

func (b *backendCTX) Init() error {
	err := b.firewall.Init()
	if err != nil {
		return err
	}
	return nil
}

func (b *backendCTX) Commit() error {
	err := b.firewall.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (b *backendCTX) ShutDown() error {
	err := b.firewall.ShutDown()
	if err != nil {
		return err
	}
	return nil
}

func (b *backendCTX) Add(decision *models.Decision) error {
	if err := b.firewall.Add(decision); err != nil {
		return err
	}
	return nil
}

func (b *backendCTX) Delete(decision *models.Decision) error {
	if err := b.firewall.Delete(decision); err != nil {
		return err
	}
	return nil
}

func isPFSupported(runtimeOS string) bool {
	var supported bool

	switch runtimeOS {
	case "openbsd", "freebsd":
		supported = true
	default:
		supported = false
	}

	return supported
}

func newBackend(config *bouncerConfig) (*backendCTX, error) {
	var err error

	b := &backendCTX{}
	log.Printf("backend type : %s", config.Mode)
	if config.DisableIPV6 {
		log.Println("IPV6 is disabled")
	}
	switch config.Mode {
	case "iptables", "ipset":
		if runtime.GOOS != "linux" {
			return nil, fmt.Errorf("iptables and ipset is linux only")
		}
		b.firewall, err = newIPTables(config)
		if err != nil {
			return nil, err
		}
	case "nftables":
		if runtime.GOOS != "linux" {
			return nil, fmt.Errorf("nftables is linux only")
		}
		b.firewall, err = newNFTables(config)
		if err != nil {
			return nil, err
		}
	case "pf":
		if !isPFSupported(runtime.GOOS) {
			return nil, fmt.Errorf("pf mode is supported only for openbsd and freebsd")
		}
		b.firewall, err = newPF(config)
		if err != nil {
			return nil, err
		}
	default:
		return b, fmt.Errorf("firewall '%s' is not supported", config.Mode)
	}
	return b, nil
}
