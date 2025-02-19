//go:build linux || windows

// https://go.googlesource.com/proposal/+/master/design/draft-gobuild.md
// $ go mod init oddstream.games/gosol
// $ go mod tidy

// the package defining a command (an excutable Go program) always has the name main
// this is a signal to go build that it must invoke the linker to make an executable file
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	// load png decoder in main package
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	sol "oddstream.games/gosol/sol"
	"oddstream.games/gosol/ui"
)

func main() {

	log.SetFlags(0)

	// pearl from the mudbank: don't have any flags that will overwrite ThePreferences
	flag.BoolVar(&sol.DebugMode, "debug", false, "turn debug graphics on")
	flag.BoolVar(&sol.NoGameLoad, "noload", false, "do not load saved game when starting")
	flag.BoolVar(&sol.NoGameSave, "nosave", false, "do not save game before exit")
	flag.BoolVar(&sol.NoCardLerp, "nolerp", false, "do not animate card movements")
	flag.BoolVar(&sol.NoCardFlip, "noflip", false, "do not animate card flips")
	flag.BoolVar(&sol.NoShuffle, "noshuf", false, "do not shuffle cards")
	flag.BoolVar(&sol.NoScrunch, "noscrunch", false, "do not scrunch cards")
	flag.BoolVar(&ui.GenerateIcons, "generateicons", false, "generate icon files")

	flag.Parse()

	if sol.DebugMode {
		for i, a := range os.Args {
			println(i, a)
		}
	}

	ebiten.SetWindowResizable(true) //ebiten panics if a window to maximize is not resizable
	if ebiten.IsWindowMaximized() || ebiten.IsWindowMinimized() {
		// GNOME (maybe) annoyingly keeps maximizing the window
		ebiten.RestoreWindow()
	}

	ebiten.SetWindowIcon(sol.WindowIcons())
	{
		var title string = sol.ThePreferences.Title
		if sol.DebugMode {
			title = fmt.Sprintf("%s (%s/%s)", title, runtime.GOOS, runtime.GOARCH)
		}
		ebiten.SetWindowTitle(title)
	}

	ebiten.SetScreenClearedEveryFrame(true)

	game, err := sol.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if !sol.NoGameLoad {
		if undoStack := sol.LoadUndoStack(); undoStack != nil {
			sol.TheBaize.SetUndoStack(undoStack)
		}
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	// we come here if the user closed the window with the x button
	// println("main exit")

	if !sol.NoGameSave {
		sol.TheBaize.Save()
	}

	sol.ThePreferences.Save()
}
