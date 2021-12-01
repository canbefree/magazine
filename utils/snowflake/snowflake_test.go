package snowflake_test

import (
	"net"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/prashantv/gostub"

	"github.com/canbefree/magazine/utils/snowflake"
)

var _ = Describe("Snowflake", func() {
	var ip1 *net.IPNet
	var ctrl *gostub.Stubs
	BeforeEach(func() {
		ip1 = &net.IPNet{
			IP:   []byte{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\xff', '\xff', '\x0a', '\x00', '\x0a', '\x01'},
			Mask: []byte{'\xff', '\x00', '\x00', '\x00'},
		}
		ctrl = gostub.Stub(&snowflake.InterfaceAddrs, func() ([]net.Addr, error) {
			return []net.Addr{
				ip1,
			}, nil
		})
	})
	It("test get ip node int", func() {
		Expect(snowflake.NewSnowflakeIP().GetNodeInt()).To(Equal(int64(101)))
	})
	AfterEach(func() {
		ctrl.Reset()
	})
})
