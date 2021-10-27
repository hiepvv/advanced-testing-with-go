package networkinghelper

import (
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
)

type networkingHelper struct {
	suite.Suite
}

func TestNetworkingHelperSuite(t *testing.T) {
	suite.Run(t, &networkingHelper{})
}

func (s *networkingHelper) InitConnHelper() (client, server net.Conn) {
	s.T().Helper()

	ln, err := net.Listen("tcp", "127.0.0.1:80")
	s.Require().NoError(err)

	go func() {
		defer func() {
			_ = ln.Close()
		}()
		server, err = ln.Accept()
		s.Require().NoError(err)
	}()
	client, err = net.Dial(ln.Addr().Network(), ln.Addr().String())

	return
}
