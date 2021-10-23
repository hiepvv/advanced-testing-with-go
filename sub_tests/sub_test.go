package sub_tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/hiepvv/advanced-testing-with-go/calculator"
)

type subTestsSuite struct {
	suite.Suite

	terminator *calculator.Terminator
}

func TestSubTestsSuite(t *testing.T) {
	suite.Run(t, &subTestsSuite{})
}

func (s *subTestsSuite) TestTerminatorFunction() {
	s.T().Run("1_plus_1", func(t *testing.T) {
		s.Require().Equal(float64(2), s.terminator.Plus(1, 1))
	})
	s.T().Run("2_multiply_3", func(t *testing.T) {
		s.Require().Equal(float64(6), s.terminator.Multiply(2, 3))
	})
	s.T().Run("3_minus_1", func(t *testing.T) {
		s.Require().Equal(float64(2), s.terminator.Minus(3, 1))
	})
	s.T().Run("6_divide_3", func(t *testing.T) {
		s.Require().Equal(float64(2), s.terminator.Divide(6, 3))
	})
}
