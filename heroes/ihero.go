package heroes

type Hero struct {
	MadeUpName string `json:"made_up_name"`
	RealName   string `json:"real_name"`
}

type IHero interface {
	ListPopularHeroes() ([]*Hero, error)
}
