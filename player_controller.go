package main

import (
	cw "TCellConsoleWrapper/tcell_wrapper"
)

var PLR_LOOP = true

func plr_control(f *faction, m *gameMap) {
	PLR_LOOP = true
	snapCursorToPawn(f, m)
	for PLR_LOOP {
		if plr_selectPawn(f, m) {
			// plr_selectOrder(f, m)
			plr_giveDefaultOrderToUnit(f, m)
		}
	}
}

func plr_selectPawn(f *faction, m *gameMap) bool { // true if pawn was selected
	f.cursor.currentCursorMode = CURSOR_SELECT
	for {
		r_renderScreenForFaction(f, m)
		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "SPACE", " ":
			PLR_LOOP = false // end turn
			return false
		case "ENTER", "RETURN":
			u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
			if u == nil {
				// log.appendMessage("SELECTED NIL")
				return false
			}
			if u.faction.factionNumber != f.factionNumber {
				log.appendMessage("Enemy units can't be selected, Commander.")
				return false
			}
			return true
		case "ESCAPE":
			GAME_IS_RUNNING = false
			PLR_LOOP = false
			return false
		default:
			plr_moveCursor(m, f, keyPressed)
		}
	}
}

func plr_selectOrder(f *faction, m *gameMap) {

}

func plr_giveDefaultOrderToUnit(f *faction, m *gameMap) {
	u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
	log.appendMessage(u.name + " is awaiting orders.")

	for {
		cx, cy := f.cursor.getCoords()
		f.cursor.currentCursorMode = CURSOR_MOVE
		r_renderScreenForFaction(f, m)


		keyPressed := cw.ReadKey()
		switch keyPressed {
		case "ENTER", "RETURN":
			issueDefaultOrderToUnit(u, m, cx, cy)
			return
		case "b": // Temporary dohuya!
			u.order = &order{orderType: order_build, x: cx, y: cy, targetBuilding: createBuilding("corekbotlab", cx, cy, f)}
			return
		case "ESCAPE":
			return
		default:
			plr_moveCursor(m, f, keyPressed)
		}
	}
}

func plr_moveCursor(g *gameMap, f *faction, keyPressed string) {
	vx, vy := plr_keyToDirection(keyPressed)
	cx, cy := f.cursor.getCoords()
	if areCoordsValid(cx+vx, cy+vy) {
		f.cursor.moveByVector(vx, vy)
	}

	snapB := f.cursor.snappedPawn
	if snapB != nil { // unsnap cursor
		for snapB.isOccupyingCoords(f.cursor.x, f.cursor.y) {
			if areCoordsValid(f.cursor.x+vx, f.cursor.y+vy) {
				f.cursor.moveByVector(vx, vy)
			} else {
				break
			}
		}
		f.cursor.snappedPawn = nil
	}
	snapCursorToPawn(f, g)
}

func snapCursorToPawn(f *faction, g *gameMap) {
	b := g.getPawnAtCoordinates(f.cursor.x, f.cursor.y)
	if b == nil {
		f.cursor.snappedPawn = nil
	} else {
		f.cursor.x, f.cursor.y = b.getCenter()
		f.cursor.snappedPawn = b
	}
}

func plr_keyToDirection(keyPressed string) (int, int) {
	switch keyPressed {
	case "2":
		return 0, 1
	case "8":
		return 0, -1
	case "4":
		return -1, 0
	case "6":
		return 1, 0
	case "7":
		return -1, -1
	case "9":
		return 1, -1
	case "1":
		return -1, 1
	case "3":
		return 1, 1
	default:
		return 0, 0
	}
}
