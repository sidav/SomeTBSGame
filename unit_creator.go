package main

// TODO: move unit stats to JSON
func createUnit(name string, x, y int, f *faction, alreadyConstructed bool) *pawn {
	var newUnit *pawn
	switch name {
	case "commander":
		newUnit = &pawn{name: "Commander",
			unitInfo:       &unit{appearance: ccell{char: '@'}},
			moveInfo:       &pawnMovementInformation{ticksForMoveSingleCell: 10},
			res:            &pawnResourceInformation{metalIncome: 1, energyIncome: 10, metalStorage: 1000, energyStorage: 1000},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 10, allowedBuildings: []string{"corekbotlab", "solar", "metalmaker", "corevehfactory"}},
		}
	case "coreck":
		newUnit = &pawn{name: "Tech 1 Construction KBot",
			unitInfo:       &unit{appearance: ccell{char: 'k'}},
			moveInfo:       &pawnMovementInformation{ticksForMoveSingleCell: 5},
			res:            &pawnResourceInformation{metalStorage: 25, energyStorage: 50},
			nanolatherInfo: &nanolatherInformation{builderCoeff: 5, allowedBuildings: []string{"corekbotlab", "solar", "metalmaker", "corevehfactory", "quark"}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 25, costM: 650, costE: 1200},
		}
	case "weasel":
		newUnit = &pawn{name: "Weasel",
			moveInfo: &pawnMovementInformation{ticksForMoveSingleCell: 10},
			unitInfo: &unit{appearance: ccell{char: 'w'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
		}
	case "thecan":
		newUnit = &pawn{name: "The Can",
			moveInfo: &pawnMovementInformation{ticksForMoveSingleCell: 10},
			unitInfo: &unit{appearance: ccell{char: 'c'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
		}
	case "ak":
		newUnit = &pawn{name: "A.K.",
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 10},
			unitInfo:                  &unit{appearance: ccell{char: 'a'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
		}
	case "thud":
		newUnit = &pawn{name: "Thud",
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 16},
			unitInfo:                  &unit{appearance: ccell{char: 't'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 12, costM: 350, costE: 650},
		}
	default:
		newUnit = &pawn{name: "UNKNOWN UNIT",
			moveInfo:                  &pawnMovementInformation{ticksForMoveSingleCell: 10},
			unitInfo:                  &unit{appearance: ccell{char: '?'}},
			currentConstructionStatus: &constructionInformation{maxConstructionAmount: 10, costM: 250, costE: 500},
		}
	}
	newUnit.x = x
	newUnit.y = y
	newUnit.faction = f
	if alreadyConstructed {
		newUnit.currentConstructionStatus = nil
	}
	return newUnit
}
