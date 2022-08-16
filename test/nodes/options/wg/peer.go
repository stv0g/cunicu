package options

import (
	"net"
	"time"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"riasc.eu/wice/pkg/crypto"
	"riasc.eu/wice/test/nodes"
)

// PublicKey specifies the public key of this peer.  PublicKey is a
// mandatory field for all PeerConfigs.
type PublicKey wgtypes.Key

func (pk PublicKey) Apply(p *nodes.WireGuardPeer) {
	p.PublicKey = wgtypes.Key(pk)
}

// PresharedKey specifies a peer's pre-shared key configuration, if not nil.
//
// A non-nil, zero-value Key will clear the pre-shared key.
type PresharedKey wgtypes.Key

func (pk PresharedKey) Apply(p *nodes.WireGuardPeer) {
	pkp := wgtypes.Key(pk)
	p.PresharedKey = &pkp
}

// Endpoint specifies the endpoint of this peer entry, if not nil.
type Endpoint net.UDPAddr

func (ep Endpoint) Apply(p *nodes.WireGuardPeer) {
	epp := net.UDPAddr(ep)
	p.Endpoint = &epp
}

// PersistentKeepaliveInterval specifies the persistent keepalive interval
// for this peer, if not nil.
//
// A non-nil value of 0 will clear the persistent keepalive interval.
type PersistentKeepaliveInterval time.Duration

func (pki PersistentKeepaliveInterval) Apply(p *nodes.WireGuardPeer) {
	pkip := time.Duration(pki)
	p.PersistentKeepaliveInterval = &pkip
}

// AllowedIPs specifies a list of allowed IP addresses in CIDR notation
// for this peer.
type AllowedIP net.IPNet

func (aip AllowedIP) Apply(p *nodes.WireGuardPeer) {
	p.AllowedIPs = append(p.AllowedIPs, net.IPNet(aip))
}

func AllowedIPv4(a, b, c, d byte, m int) AllowedIP {
	return AllowedIP{
		IP:   net.IPv4(a, b, c, d),
		Mask: net.CIDRMask(m, 32),
	}
}

func AllowedIPStr(str string) AllowedIP {
	ip, n, _ := net.ParseCIDR(str)

	return AllowedIP{
		IP:   ip,
		Mask: n.Mask,
	}
}

func Peer(pk crypto.Key, opts ...nodes.WireGuardPeerOption) *nodes.WireGuardPeer {
	p := &nodes.WireGuardPeer{
		PeerConfig: wgtypes.PeerConfig{
			PublicKey: wgtypes.Key(pk),
		},
	}

	for _, o := range opts {
		switch opt := o.(type) {
		case nodes.WireGuardPeerOption:
			opt.Apply(p)
		}
	}

	return p
}

func PeerFromNames(agentName, intfName string, opts ...nodes.WireGuardPeerOption) *nodes.WireGuardPeer {
	sk := crypto.PrivateKeyFromStrings(agentName, intfName)
	return Peer(sk.PublicKey(), opts...)
}
