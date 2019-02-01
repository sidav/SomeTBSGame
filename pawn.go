package main

import "fmt"

type pawn struct {
	// pawn is a building or a unit.
	name                      string
	unitInfo                  *unit
	buildingInfo              *building
	faction                   *faction
	x, y                      int
	order                     *order
	res                       *pawnResourceInformation
	nanolatherInfo            *nanolatherInformation
	currentConstructionStatus *constructionInformation
	moveInfo                  *pawnMovementInformation
	weapons                   []*pawnWeaponInformation
	nextTurnToAct             int
	isCommander               bool
	// armor info:
	hitpoints, maxHitpoints int
	isLight, isHeavy        bool // these are not mutually excluding lol. Trust me, I'm a programmer
}

func (p *pawn) hasWeapons() bool {
	return len(p.weapons) > 0
}

func (p *pawn) getMaxRadiusToFire() int {
	max := 0
	for _, weap := range p.weapons {
		if weap.attackRadius > max {
			max = weap.attackRadius
		}
	}
	return max
}

func (p *pawn) isUnit() bool {
	return p.unitInfo != nil
}

func (p *pawn) isBuilding() bool {
	return p.buildingInfo != nil
}

func (p *pawn) getCoords() (int, int) {
	return p.x, p.y
}

func (p *pawn) setOrder(o *order) {
	p.order = o
	log.appendMessage(fmt.Sprintf("Order for %d, %d received!", o.x, o.y))
}

func (p *pawn) isOccupyingCoords(x, y int) bool {
	if p.isBuilding() {
		return areCoordsInRect(x, y, p.x, p.y, p.buildingInfo.w, p.buildingInfo.h)
	} else {
		return x == p.x && y == p.y
	}
}

func (p *pawn) getCenter() (int, int) {
	if p.isUnit() {
		return p.x, p.y
	} else {
		return p.x + p.buildingInfo.w/2, p.y + p.buildingInfo.h/2
	}
}

//func (p *pawn) executeOrders(m *gameMap)	{
//	if p.isBuilding() {
//		return
//	} else {
//		p.executeOrders(m)
//	}
//}

func (p *pawn) getCurrentOrderDescription() string {
	if p.currentConstructionStatus != nil {
		constr := p.currentConstructionStatus
		return fmt.Sprintf("UNDER CONSTRUCTION: %d%%", constr.getCompletionPercent())
	}
	if p.order == nil {
		return "STANDBY"
	}
	switch p.order.orderType {
	case order_hold:
		return "STANDBY"
	case order_move:
		return "MOVING"
	case order_attack:
		return "ASSAULTING"
	case order_build:
		return fmt.Sprintf("NANOLATHING (%d%% ready)",
			p.order.buildingToConstruct.currentConstructionStatus.getCompletionPercent())
	case order_construct:
		if len(p.order.constructingQueue) > 0 {
			return fmt.Sprintf("CONSTRUCTING: %s (%d%% ready)", p.order.constructingQueue[0].name,
				p.order.constructingQueue[0].currentConstructionStatus.getCompletionPercent())
		} else {
			return "FINISHING CONSTRUCTION"
		}
	default:
		return "DOING SOMETHING"
	}
}
