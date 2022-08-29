package pdisc

import (
	"fmt"
	"net"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"riasc.eu/wice/pkg/crypto"
)

func (pd *PeerDescription) Config() wgtypes.PeerConfig {
	var pk crypto.Key
	if pd.PublicKeyNew != nil {
		pk, _ = crypto.ParseKeyBytes(pd.PublicKeyNew)
	} else {
		pk, _ = crypto.ParseKeyBytes(pd.PublicKey)
	}

	allowedIPs := []net.IPNet{}
	for _, allowedIP := range pd.AllowedIps {
		_, ipnet, err := net.ParseCIDR(allowedIP)
		if err != nil {
			panic(fmt.Errorf("failed to parse WireGuard AllowedIP: %w", err))
		}

		allowedIPs = append(allowedIPs, *ipnet)
	}

	return wgtypes.PeerConfig{
		ReplaceAllowedIPs: true,
		PublicKey:         wgtypes.Key(pk),
		AllowedIPs:        allowedIPs,
	}
}
