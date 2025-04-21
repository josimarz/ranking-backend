package mock

import "github.com/josimarz/ranking-backend/internal/domain/entity"

var (
	Rank entity.Rank = entity.Rank{
		Id:     "1ac85e34-cb6f-40c9-97bb-16267877bb13",
		Name:   "Video Game Consoles",
		Public: true,
	}
	Attrs []entity.Attribute = []entity.Attribute{{
		Id:     "be44503b-1fac-4d5a-aae0-0239159bdc4a",
		Name:   "Controls",
		Desc:   "Evaluate the quality and accessibility of controls",
		Order:  1,
		RankId: "1ac85e34-cb6f-40c9-97bb-16267877bb13",
	}, {
		Id:     "53e1515d-7fed-4d94-8b36-4cd49b2f11be",
		Name:   "Graphics",
		Desc:   "Evaluate the graphics capacity of the console",
		Order:  2,
		RankId: "1ac85e34-cb6f-40c9-97bb-16267877bb13",
	}, {
		Id:     "b2ac5f2c-a65c-4eb8-a0e1-a66a6bea4aac",
		Name:   "Sound",
		Desc:   "Evaluate the sound capacity of the console",
		Order:  3,
		RankId: "1ac85e34-cb6f-40c9-97bb-16267877bb13",
	}}
	Entries []entity.Entry = []entity.Entry{{
		Id:       "d10961ca-e9ed-4d3b-b086-f756a3118894",
		Name:     "Neo Geo CD",
		ImageURL: "https://videogame.com/neo-geo-cd.png",
		Scores: entity.Scores{
			"Controls": 90,
			"Graphics": 97,
			"Sound":    97,
		},
		RankId: "1ac85e34-cb6f-40c9-97bb-16267877bb13",
	}, {
		Id:       "e006f3be-88a4-4891-8c8e-f1de6d6b5324",
		Name:     "Nintendo Entertainment System",
		ImageURL: "https://videogame.com/nes.png",
		Scores: entity.Scores{
			"Controls": 70,
			"Graphics": 72,
			"Sound":    70,
		},
		RankId: "1ac85e34-cb6f-40c9-97bb-16267877bb13",
	}, {
		Id:       "da2b4fc6-f933-4214-b742-4f199aec2481",
		Name:     "Sega Master System",
		ImageURL: "https://videogame.com/sms.png",
		Scores: entity.Scores{
			"Controls": 73,
			"Graphics": 78,
			"Sound":    76,
		},
		RankId: "1ac85e34-cb6f-40c9-97bb-16267877bb13",
	}, {
		Id:       "25658fa3-6721-42ae-8e25-7ba9c8f1cd85",
		Name:     "Sega Mega Drive",
		ImageURL: "https://videogame.com/smd.png",
		Scores: entity.Scores{
			"Controls": 80,
			"Graphics": 84,
			"Sound":    83,
		},
		RankId: "1ac85e34-cb6f-40c9-97bb-16267877bb13",
	}, {
		Id:       "959c559e-db6a-4c4a-9164-f3eab305e076",
		Name:     "Super Nintendo Entertainment System",
		ImageURL: "https://videogame.com/snes.png",
		Scores: entity.Scores{
			"Controls": 84,
			"Graphics": 89,
			"Sound":    87,
		},
		RankId: "1ac85e34-cb6f-40c9-97bb-16267877bb13",
	}}
)
