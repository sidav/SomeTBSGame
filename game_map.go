package main

const (
	mapW = 70
	mapH = 20
)

type gameMap struct {
	tileMap [mapW][mapH] *tile
	factions []*faction
	pawns []*pawn
}

func (g *gameMap) addPawn(p *pawn) {
	g.pawns = append(g.pawns, p)
}

func (g *gameMap) addBuilding(b *pawn, asAlreadyConstructed bool) {
	if asAlreadyConstructed {
		b.currentConstructionStatus = nil
		b.buildingInfo.hasBeenPlaced = true
	}

	if b.nanolatherInfo != nil && len(b.nanolatherInfo.allowedUnits) > 0 { // sets default rally point for build units.
		b.nanolatherInfo.defaultOrderForUnitBuilt = &order{orderType: order_move, x: b.x + b.buildingInfo.w / 2, y: b.y + b.buildingInfo.h + 1}
	}

	g.addPawn(b)
}

func (g *gameMap) removePawn(p *pawn) {
	for i := 0; i < len(g.pawns); i++ {
		if p == g.pawns[i] {
			g.pawns = append(g.pawns[:i], g.pawns[i+1:]...) // ow it's fucking... magic!
		}
	}
}

func (g *gameMap) getPawnAtCoordinates(x, y int) *pawn {
	for _, b := range g.pawns {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

//func (g *gameMap) getUnitAtCoordinates(x, y int) *pawn { // TODO: remove (as a duplicate of getPawnAtCoordinates)
//	for _, b := range g.pawns {
//		if b.isOccupyingCoords(x, y) {
//			return b
//		}
//	}
//	return nil
//}

func (g *gameMap) getPawnsInRadiusFrom(x, y, radius int) []*pawn {
	var arr []*pawn
	for _, p := range g.pawns {
		px, py := p.getCenter()
		if getSqDistanceBetween(x, y, px, py) <= radius*radius {
			arr = append(arr, p)
		}
	}
	return arr
}

func (g *gameMap) getPawnsInRect(x, y, w, h int) []*pawn {
	var arr []*pawn
	for _, p := range g.pawns {
		px, py := p.getCenter()
		if (px >= x) && (px < x+w) && (py >= y) && (py < y+h) { // TODO: pawns bigger than one cell
			arr = append(arr, p)
		}
	}
	return arr
}

func (g *gameMap) getEnemyPawnsInRadiusFrom(x, y, radius int, f *faction) []*pawn {
	var arr []*pawn
	for _, p := range g.pawns {
		px, py := p.getCenter()
		if p.faction != f && getSqDistanceBetween(x, y, px, py) <= radius*radius {
			arr = append(arr, p)
		}
	}
	return arr
}

func (g *gameMap) getBuildingAtCoordinates(x, y int) *pawn {
	for _, b := range g.pawns {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

func (g *gameMap) getNumberOfMetalDepositsInRect(x, y, w, h int) int {
	total := 0
	for i:=0; i<w;i++ {
		for j:=0;j<h;j++{
			if areCoordsValid(x+i, y+j) {
				total += g.tileMap[x+i][y+j].metalAmount
			}
		}
	}
	return total
}

func (g *gameMap) getNumberOfMetalDepositsUnderBuilding(b *pawn) int {
	return g.getNumberOfMetalDepositsInRect(b.x, b.y, b.buildingInfo.w, b.buildingInfo.h)
}

func (g *gameMap) canBuildingBeBuiltAt(b *pawn, cx, cy int) bool {
	b.x = cx - b.buildingInfo.w/2
	b.y = cy - b.buildingInfo.h/2
	if b.buildingInfo.canBeBuiltOnMetalOnly && g.getNumberOfMetalDepositsUnderBuilding(b) == 0 {
		return false
	}
	if len(g.getPawnsInRect(b.x, b.y, b.buildingInfo.w, b.buildingInfo.h)) > 0 {
		return false
	}
	return true
}

func (g *gameMap) init() {
	g.pawns = make([]*pawn, 0)
	g.factions = make([]*faction, 0)
	for i:=0; i < mapW; i++ {
		for j:=0; j < mapH; j++ {
			g.tileMap[i][j] = &tile{appearance: &ccell{char: '.', r: 64, g: 128, b: 64, color: 3}, isPassable: true}
		}
	}

	// place metal deposits
	g.tileMap[2][2] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}
	g.tileMap[3][2] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}
	g.tileMap[2][3] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}
	g.tileMap[13][2] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}
	g.tileMap[14][3] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}
	g.tileMap[13][1] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}
	g.tileMap[12][15] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}
	g.tileMap[11][15] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}
	g.tileMap[12][16] = &tile{appearance: &ccell{char: ';', r: 64, g: 64, b: 128, color: 8}, metalAmount: 1, isPassable: true}


	g.factions = append(g.factions, createFaction("The Core Corporation", 0, true))
	g.addPawn(createUnit("protocommander", 3, 9, g.factions[0], true))
	g.factions[0].cursor.x = 3
	g.factions[0].cursor.y = 9
	// g.addBuilding(createBuilding("metalmaker", 5, 1, g.factions[0]), true)
	// g.addUnit(createUnit("weasel", 3, 6, g.factions[0]))
	// g.addUnit(createUnit("thecan", 3, 4, g.factions[0]))
	// g.addBuilding(createBuilding("corekbotlab", 5, 1, g.factions[0]), true)
	// g.addBuilding(createBuilding("corevehfactory", 5, 5, g.factions[0]), true)

	g.factions = append(g.factions, createFaction("The rogue Arm AI", 1, false))
	// g.addPawn(createUnit("armcommander", mapW-10, 5, g.factions[1], true))
	g.addBuilding(createBuilding("armhq", mapW-5, 9, g.factions[1]), true )
	// g.addPawn(createUnit("ak", mapW-1, 4, g.factions[1], true))
	g.addBuilding(createBuilding("lturret", mapW-10, 1, g.factions[1]), true )
	g.addBuilding(createBuilding("lturret", mapW-10, 4, g.factions[1]), true)
	g.addBuilding(createBuilding("guardian", mapW-7, 3, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 8, g.factions[1]), true )
	g.addBuilding(createBuilding("lturret", mapW-10, 12, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 16, g.factions[1]), true )
	g.addBuilding(createBuilding("guardian", mapW-7, 14, g.factions[1]), true)
	g.addBuilding(createBuilding("lturret", mapW-10, 19, g.factions[1]), true )

	for _, f := range g.factions {
		f.recalculateFactionEconomy(g)
	}
}
