package main

import (
	"SomeTBSGame/routines"
	"strconv"
)

const (
	AI_WRITE_DEBUG_TO_LOG = true
)

var (
	AI_CONTROL_PERIOD = 100
	AI_MIN_PERIOD           = 10
	AI_PERIOD_DECREMENT     = 5
)

func ai_write(text string) {
	if AI_WRITE_DEBUG_TO_LOG {
		log.appendMessage("AI: "+text)
	}
}

func ai_controlFaction(f *faction) {
	if CURRENT_TURN/10%AI_CONTROL_PERIOD != 0 {
		return
	}
	if AI_CONTROL_PERIOD -AI_PERIOD_DECREMENT >= AI_MIN_PERIOD{
		AI_CONTROL_PERIOD -= AI_PERIOD_DECREMENT
		ai_write("PERIOD changed to " + strconv.Itoa(AI_CONTROL_PERIOD))
	}
	ai_write("assuming direct control over " + f.name)
	for _, p := range CURRENT_MAP.pawns {
		if p.faction == f {
			ai_controlPawn(p)
		}
	}
}

func ai_controlPawn(p *pawn) {
	if p.order != nil {
		return
	}
	if p.canMove() && p.hasWeapons() {
		enemyCommander := ai_getEnemyCommander(p.faction)
		if enemyCommander != nil {
			p.order = &order{orderType: order_attack_move, x: enemyCommander.x, y: enemyCommander.y}
			return
		}
	}
	if p.canConstructUnits() {
		productionVariants := &p.nanolatherInfo.allowedUnits
		pawnToProduce := ai_decideProduction(productionVariants, p.faction)
		p.order = &order{orderType: order_construct, constructingQueue: []*pawn{pawnToProduce}}
	}
}

func ai_getEnemyCommander(f *faction) *pawn {
	for _, p := range CURRENT_MAP.pawns {
		if p.faction != f && p.isCommander {
			return p
		}
	}
	return nil
}

func ai_decideProduction(variants *[]string, f *faction) *pawn {
	listOfCombatUnits := make([]*pawn, 0)
	for _, variant := range *variants {
		pawnUnderConsideration := createUnit(variant, 0, 0, f, false)
		if pawnUnderConsideration.canMove() && pawnUnderConsideration.hasWeapons() {
			listOfCombatUnits = append(listOfCombatUnits, pawnUnderConsideration)
		}
	}
	if len(listOfCombatUnits) > 0 {
		pwn :=  listOfCombatUnits[routines.Random(len(listOfCombatUnits))]
		ai_write("producing " + pwn.name)
		return pwn
	}
	ai_write("nothing to produce.")
	return nil
}