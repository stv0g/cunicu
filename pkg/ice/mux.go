package ice

import (
	"net"

	"github.com/pion/ice/v2"
	"github.com/pion/zapion"
	"go.uber.org/zap"
)

type PacketListener interface{}

func NewMultiUDPMuxWithListen(listen func(ip net.IP) (net.PacketConn, error), interfaceFilter func(string) bool, ipFilter func(net.IP) bool, networkTypes []NetworkType, includeLoopback bool) (*ice.MultiUDPMuxDefault, error) {
	ips, err := localInterfaces(interfaceFilter, ipFilter, networkTypes, includeLoopback)
	if err != nil {
		return nil, err
	}

	conns := make([]net.PacketConn, 0, len(ips))
	for _, ip := range ips {
		conn, listenErr := listen(ip)
		if listenErr != nil {
			err = listenErr
			break
		}

		conns = append(conns, conn)
	}

	if err != nil {
		for _, conn := range conns {
			_ = conn.Close()
		}
		return nil, err
	}

	lf := zapion.ZapFactory{
		BaseLogger: zap.L().Named("ice"),
	}

	muxes := make([]ice.UDPMux, 0, len(conns))
	for _, conn := range conns {
		mux := ice.NewUDPMuxDefault(ice.UDPMuxParams{
			Logger:  lf.NewLogger("udpmux"),
			UDPConn: conn,
		})
		muxes = append(muxes, mux)
	}

	mux := ice.NewMultiUDPMuxDefault(muxes...)

	return mux, nil
}

func localInterfaces(interfaceFilter func(string) bool, ipFilter func(net.IP) bool, networkTypes []NetworkType, includeLoopback bool) ([]net.IP, error) { //nolint:gocognit
	ips := []net.IP{}
	ifaces, err := net.Interfaces()
	if err != nil {
		return ips, err
	}

	var IPv4Requested, IPv6Requested bool
	for _, typ := range networkTypes {
		if typ.IsIPv4() {
			IPv4Requested = true
		}

		if typ.IsIPv6() {
			IPv6Requested = true
		}
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if (iface.Flags&net.FlagLoopback != 0) && !includeLoopback {
			continue // loopback interface
		}

		if interfaceFilter != nil && !interfaceFilter(iface.Name) {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch addr := addr.(type) {
			case *net.IPNet:
				ip = addr.IP
			case *net.IPAddr:
				ip = addr.IP
			}
			if ip == nil || (ip.IsLoopback() && !includeLoopback) {
				continue
			}

			if ipv4 := ip.To4(); ipv4 == nil {
				if !IPv6Requested {
					continue
				} else if !isSupportedIPv6(ip) {
					continue
				}
			} else if !IPv4Requested {
				continue
			}

			if ipFilter != nil && !ipFilter(ip) {
				continue
			}

			ips = append(ips, ip)
		}
	}
	return ips, nil
}

// The conditions of invalidation written below are defined in
// https://tools.ietf.org/html/rfc8445#section-5.1.1.1
func isSupportedIPv6(ip net.IP) bool {
	if len(ip) != net.IPv6len ||
		isZeros(ip[0:12]) || // !(IPv4-compatible IPv6)
		ip[0] == 0xfe && ip[1]&0xc0 == 0xc0 || // !(IPv6 site-local unicast)
		ip.IsLinkLocalUnicast() ||
		ip.IsLinkLocalMulticast() {
		return false
	}
	return true
}

func isZeros(ip net.IP) bool {
	for i := 0; i < len(ip); i++ {
		if ip[i] != 0 {
			return false
		}
	}
	return true
}