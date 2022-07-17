package socket

import (
	"net"

	"github.com/pion/ice/v2"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"riasc.eu/wice/internal/wg"
	"riasc.eu/wice/pkg/core"
	"riasc.eu/wice/pkg/feat/disc/ep"
	"riasc.eu/wice/pkg/pb"
	"riasc.eu/wice/pkg/signaling"
)

func (s *Server) OnInterfaceAdded(i *core.Interface) {
	s.events.C <- &pb.Event{
		Type:      pb.Event_INTERFACE_ADDED,
		Interface: i.Name(),
	}
}

func (s *Server) OnInterfaceRemoved(i *core.Interface) {
	s.events.C <- &pb.Event{
		Type:      pb.Event_INTERFACE_REMOVED,
		Interface: i.Name(),
	}
}

func (s *Server) OnInterfaceModified(i *core.Interface, old *wg.Device, mod core.InterfaceModifier) {
	s.events.C <- &pb.Event{
		Type:      pb.Event_INTERFACE_MODIFIED,
		Interface: i.Name(),
		Event: &pb.Event_InterfaceModified{
			InterfaceModified: &pb.InterfaceModifiedEvent{
				Modified: uint32(mod),
			},
		},
	}
}

func (s *Server) OnPeerAdded(p *core.Peer) {
	s.events.C <- &pb.Event{
		Type:      pb.Event_PEER_ADDED,
		Interface: p.Interface.Name(),
		Peer:      p.PublicKey().Bytes(),
	}
}

func (s *Server) OnPeerRemoved(p *core.Peer) {
	s.events.C <- &pb.Event{
		Type:      pb.Event_PEER_REMOVED,
		Interface: p.Interface.Name(),
		Peer:      p.PublicKey().Bytes(),
	}
}

func (s *Server) OnPeerModified(p *core.Peer, old *wgtypes.Peer, mod core.PeerModifier, ipsAdded, ipsRemoved []net.IPNet) {
	s.events.C <- &pb.Event{
		Type:      pb.Event_PEER_MODIFIED,
		Interface: p.Interface.Name(),
		Peer:      p.PublicKey().Bytes(),

		Event: &pb.Event_PeerModified{
			PeerModified: &pb.PeerModifiedEvent{
				Modified: uint32(mod),
			},
		},
	}
}

func (s *Server) OnConnectionStateChange(p *ep.Peer, cs ice.ConnectionState) {
	s.events.C <- &pb.Event{
		Type: pb.Event_PEER_CONNECTION_STATE_CHANGED,

		Interface: p.Interface.Name(),
		Peer:      p.PublicKey().Bytes(),

		Event: &pb.Event_PeerConnectionStateChange{
			PeerConnectionStateChange: &pb.PeerConnectionStateChangeEvent{
				NewState: pb.NewConnectionState(cs),
			},
		},
	}
}

func (s *Server) OnBackendReady(b signaling.Backend) {
	s.events.C <- &pb.Event{
		Type: pb.Event_BACKEND_READY,

		Event: &pb.Event_BackendReady{
			BackendReady: &pb.BackendReadyEvent{
				Type: b.Type(),
			},
		},
	}
}
