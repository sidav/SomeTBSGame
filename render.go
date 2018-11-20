package main

import (
	cw "TCellConsoleWrapper/tcell_wrapper"
	"fmt"
)

const (
	CONSOLE_W = 80
	CONSOLE_H = 25
	VIEWPORT_W = 40
	VIEWPORT_H = 20
)

func r_setFgColorByCcell(c *ccell) {
	cw.SetFgColor(c.color)
	// cw.SetFgColorRGB(c.r, c.g, c.b)
}

func r_renderScreenForFaction(f *faction, g*gameMap) {
	r_renderMapAroundCursor(g, f.cx, f.cy)
	renderFactionStats(f)
	cw.Flush_console()
}

func r_renderMapAroundCursor(g *gameMap, cx, cy int) {
	cw.Clear_console()
	vx := cx - VIEWPORT_W / 2
	vy := cy - VIEWPORT_H / 2
	renderMapInViewport(g, vx, vy)
	renderUnitsInViewport(g, vx, vy)
	renderLog(false)
}

func renderMapInViewport(g *gameMap, vx, vy int) {
	for x := vx; x < vx+VIEWPORT_W; x++ {
		for y := vy; y < vy+VIEWPORT_H; y++ {
			if areCoordsValid(x, y) {
				tileApp := g.tileMap[x][y].appearance
				r_setFgColorByCcell(tileApp)
				cw.PutChar(tileApp.char, x-vx, y-vy)
			}
		}
	}
}

func renderUnitsInViewport(g *gameMap, vx, vy int) {
	for _, u := range g.units {
		tileApp := u.appearance
		// r, g, b := getFactionRGB(u.faction.factionNumber)
		// cw.SetFgColorRGB(r, g, b)
		cw.SetFgColor(getFactionColor(u.faction.factionNumber))
		cw.PutChar(tileApp.char, u.x-vx, u.y-vy)
	}

}

func renderSelectCursor() {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	// cw.SetFgColorRGB(128, 128, 128)
	cw.SetFgColor(cw.WHITE)

	cw.PutChar('[', x-1, y)
	cw.PutChar(']', x+1, y)

	// outcommented for non-SDL console
	//cw.PutChar(16*13+10, x-1, y-1)
	//cw.PutChar(16*11+15, x+1, y-1)
	//cw.PutChar(16*12, x-1, y+1)
	//cw.PutChar(16*13+9, x+1, y+1)
	cw.Flush_console()
}

func renderMoveCursor() {
	x := VIEWPORT_W / 2
	y := VIEWPORT_H / 2
	// cw.SetFgColorRGB(128, 255, 128)
	cw.SetFgColor(cw.GREEN)

	cw.PutChar('>', x-1, y)
	cw.PutChar('<', x+1, y)

	//cw.PutChar('\\', x-1, y-1)
	//cw.PutChar('/', x+1, y-1)
	//cw.PutChar('/', x-1, y+1)
	//cw.PutChar('\\', x+1, y+1)

	cw.Flush_console()
}

func renderFactionStats(f *faction) {
	statsx := VIEWPORT_W+1
	
	// fr, fg, fb := getFactionRGB(f.factionNumber)
	// cw.SetFgColorRGB(fr, fg, fb)
	cw.SetFgColor(getFactionColor(f.factionNumber))
	cw.PutString(fmt.Sprintf("%s: turn %d", f.name, CURRENT_TURN), statsx, 0)
	
	metal, maxmetal := f.currentMetal, f.maxMetal
	cw.SetFgColor(cw.DARK_CYAN)
	renderStatusbar("METAL", metal, maxmetal, statsx, 1, CONSOLE_W-VIEWPORT_W-3, cw.DARK_CYAN)
	
	energy, maxenergy := f.currentEnergy, f.maxEnergy
	cw.SetFgColor(cw.DARK_YELLOW)
	renderStatusbar("ENERGY", energy, maxenergy, statsx, 2, CONSOLE_W-VIEWPORT_W-3, cw.DARK_YELLOW)
}

func renderStatusbar(name string, curvalue, maxvalue, x, y, width, barColor int) {
	barTitle := name
	cw.PutString(barTitle, x, y)
	barWidth := width - len(name)
	filledCells := barWidth * curvalue / maxvalue
	barStartX := x + len(barTitle) + 1
	for i := 0; i < barWidth; i++ {
		if i < filledCells {
			cw.SetFgColor(barColor)
			cw.PutChar('=', i+barStartX, y)
		} else {
			cw.SetFgColor(cw.DARK_BLUE)
			cw.PutChar('-', i+barStartX, y)
		}
	}
}

func renderLog(flush bool) {
	cw.SetFgColor(cw.WHITE)
	for i := 0; i < LOG_HEIGHT; i++ {
		cw.PutString(log.last_msgs[i].getText(), 0, VIEWPORT_H+i)
	}
	if flush {
		cw.Flush_console()
	}
}