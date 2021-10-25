package tabledriventests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/hiepvv/advanced-testing-with-go/calculator"
)

type tableDrivenTestsSuite struct {
	suite.Suite

	terminator *calculator.Terminator
}

func TestTableDrivenTestsSuite(t *testing.T) {
	suite.Run(t, &tableDrivenTestsSuite{})
}

func (s *tableDrivenTestsSuite) TestTerminatorPlus() {
	cases := []struct {
		Name           string
		A, B, Expected float64
	}{
		{
			Name:     "1_plus_1",
			A:        1,
			B:        1,
			Expected: 2,
		},
		{
			Name:     "2_plus_3",
			A:        2,
			B:        3,
			Expected: 5,
		},
		{
			Name:     "3_plus_1",
			A:        3,
			B:        1,
			Expected: 4,
		},
		{
			Name:     "6_plus_3",
			A:        6,
			B:        3,
			Expected: 9,
		},
	}
	for _, tc := range cases {
		s.T().Run(tc.Name, func(t *testing.T) {
			s.Require().Equal(tc.Expected, s.terminator.Plus(tc.A, tc.B))
		})
	}
}
