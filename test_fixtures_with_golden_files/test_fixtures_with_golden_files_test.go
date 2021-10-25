package testfixtureswithgoldenfiles

import (
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/suite"

	"github.com/hiepvv/advanced-testing-with-go/heroes"
)

type testFixturesWithGoldenFilesSuite struct {
	suite.Suite
	goldie *goldie.Goldie

	marvelDirectory *heroes.Marvel
}

func TestFixturesWithGoldenFilesSuite(t *testing.T) {
	suite.Run(t, &testFixturesWithGoldenFilesSuite{})
}

func (s *testFixturesWithGoldenFilesSuite) SetupSuite() {
	s.goldie = goldie.New(s.T(), goldie.WithFixtureDir("test_fixtures"))
}

func (s *testFixturesWithGoldenFilesSuite) MatchJSON(data interface{}) {
	s.goldie.AssertJson(s.T(), s.T().Name(), data)
}

func (s *testFixturesWithGoldenFilesSuite) TestListPopularHeroes() {
	marvelHeroes, err := s.marvelDirectory.ListPopularHeroes()
	s.Require().NoError(err)
	s.MatchJSON(marvelHeroes)
}
