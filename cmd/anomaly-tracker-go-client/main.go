package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	initGui()
}

func initGui() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}

	return g.SetViewOnTop(name)
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("info", 0, 0, maxX/2-1, maxY/6-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Anomaly Tracker"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, " Welcome to the anomaly tracker.")
	}
	if v, err := g.SetView("stats", maxX/2, 0, (maxX/5)*4-1, maxY/6-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Stats"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, " Anomalies: 5")
		fmt.Fprintln(v, " Ping: 67ms")
	}
	if v, err := g.SetView("list", 0, maxY/6, (maxX/5)*4-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Anomalies"
		v.Editable = false
		v.Wrap = true

		if _, err = setCurrentViewOnTop(g, "list"); err != nil {
			return err
		}
	}
	if v, err := g.SetView("send", (maxX/5)*4, 0, maxX-1, maxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Send Anomaly"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, " Id     |")
		fmt.Fprintln(v, " System |")
		fmt.Fprintln(v, " Type   |")
		fmt.Fprintln(v, " Name   |")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	log.Println("Shutting down.")
	return gocui.ErrQuit
}
