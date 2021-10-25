package heroes

type Marvel struct{}

func (m *Marvel) ListPopularHeroes() ([]*Hero, error) {
	return []*Hero{
		{
			MadeUpName: "Ant Man",
			RealName:   "Scott Lang",
		},
		{
			MadeUpName: "Black Panther",
			RealName:   "T'Challa",
		},
		{
			MadeUpName: "Captain America",
			RealName:   "Steve Rogers",
		},
		{
			MadeUpName: "Doctor Strange",
			RealName:   "Stephen Vincent Strange",
		},
	}, nil
}
