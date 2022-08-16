//go:build linux

package test_test

import (
	"riasc.eu/wice/pkg/wg"
	"riasc.eu/wice/test/nodes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	g "github.com/stv0g/gont/pkg"
	gopt "github.com/stv0g/gont/pkg/options"
)

var _ = Context("single", Pending, func() {
	var n Network

	BeforeEach(func() {
		n.Init()

		n.AgentOptions = append(n.AgentOptions,
			gopt.EmptyDir(wg.ConfigPath),
			gopt.EmptyDir(wg.SocketPath),
		)
	})

	AfterEach(func() {
		n.Close()
	})

	JustBeforeEach(func() {
		By("Initializing core network")

		nw, err := g.NewNetwork(n.Name, n.NetworkOptions...)
		Expect(err).To(Succeed(), "Failed to create network: %s", err)

		By("Initializing agent node")

		n1, err := nodes.NewAgent(nw, "n1", n.AgentOptions...)
		Expect(err).To(Succeed(), "Failed to create agent node: %s", err)

		By("Starting network")

		n.Network = nw
		n.AgentNodes = nodes.AgentList{n1}
	})

	Context("create", func() {
		Context("kernel", func() {

		})

		Context("userspace", func() {

		})
	})

	Context("watcher", Ordered, func() {
		It("detects a new interface", func() {

		})

		It("detects a change of the interface", func() {

		})

		It("detects a new peer", func() {

		})

		It("detects a change of the peer", func() {

		})

		It("detects the removal of the peer", func() {

		})

		It("detects the removal of the interface", func() {

		})
	})

	Context("host-sync", func() {

	})

	Context("route-sync", func() {

	})

	Context("config-sync", func() {

	})
})
