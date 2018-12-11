package main

func issueDefaultOrderToUnit(p *pawn, m *gameMap, x, y int) {
	cx, cy := p.getCenter()
	if x == cx && y == cy {
		p.reportOrderCompletion(p.getCurrentOrderDescription() + " order untouched")
		return
	}
	target := m.getPawnAtCoordinates(x, y)
	if target != nil {
		if target.isBuilding() && target.currentConstructionStatus.isCompleted() == false {
			p.setOrder(&order{orderType: order_build, buildingToConstruct: target})
			log.appendMessage(p.name + ": Helps nanolathing")
			return
		}
	}
	if p.canMove() {
		p.setOrder(&order{orderType: order_move, x: x, y: y})
	}  else if p.canConstructUnits() {
		p.nanolatherInfo.defaultOrderForUnitBuilt = &order{orderType: order_move, x: x, y: y}
		p.reportOrderCompletion("rally point set")
	}
}
