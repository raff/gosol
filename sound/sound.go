package sound

import (
	"bytes"
	_ "embed" // go:embed only allowed in Go files that import "embed"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"oddstream.games/gosol/util"
)

//go:embed assets/cardFan1.wav
var cardFan1Bytes []byte

//go:embed assets/cardFan2.wav
var cardFan2Bytes []byte

//go:embed assets/cardOpenPackage1.wav
var cardOpenPackage1Bytes []byte

//go:embed assets/cardOpenPackage2.wav
var cardOpenPackage2Bytes []byte

//go:embed assets/cardPlace1.wav
var cardPlace1Bytes []byte

//go:embed assets/cardPlace2.wav
var cardPlace2Bytes []byte

//go:embed assets/cardPlace3.wav
var cardPlace3Bytes []byte

//go:embed assets/cardPlace4.wav
var cardPlace4Bytes []byte

//go:embed assets/cardShove1.wav
var cardShove1Bytes []byte

//go:embed assets/cardShove2.wav
var cardShove2Bytes []byte

//go:embed assets/cardShove3.wav
var cardShove3Bytes []byte

//go:embed assets/cardShove4.wav
var cardShove4Bytes []byte

//go:embed assets/cardShuffle.wav
var cardShuffleBytes []byte

//go:embed assets/cardSlide1.wav
var cardSlide1Bytes []byte

//go:embed assets/cardSlide2.wav
var cardSlide2Bytes []byte

//go:embed assets/cardSlide3.wav
var cardSlide3Bytes []byte

//go:embed assets/cardSlide4.wav
var cardSlide4Bytes []byte

//go:embed assets/cardSlide5.wav
var cardSlide5Bytes []byte

//go:embed assets/cardSlide6.wav
var cardSlide6Bytes []byte

//go:embed assets/cardSlide7.wav
var cardSlide7Bytes []byte

//go:embed assets/cardSlide8.wav
var cardSlide8Bytes []byte

//go:embed assets/cardTakeOutPackage1.wav
var cardTakeOutPackage1Bytes []byte

//go:embed assets/cardTakeOutPackage2.wav
var cardTakeOutPackage2Bytes []byte

//go:embed assets/complete.wav
var completeBytes []byte

// https://freesound.org/people/AlienXXX/sounds/249895/
//go:embed assets/249895__alienxxx__blip2.wav
var blipBytes []byte

var audioContext *audio.Context

var soundMap map[string]*audio.Player

var Volume float64

func decode(name string, wavBytes []byte) {
	if len(wavBytes) == 0 {
		log.Panic("empty wav file ", name)
	}
	d, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(wavBytes))
	if err != nil {
		log.Panic(err)
	}
	audioPlayer, err := audioContext.NewPlayer(d)
	if err != nil {
		log.Panic(err)
	}
	soundMap[name] = audioPlayer
}

func init() {
	defer util.Duration(time.Now(), "sound.init")

	audioContext = audio.NewContext(44100)
	soundMap = make(map[string]*audio.Player)

	decode("Fan1", cardFan1Bytes)
	decode("Fan2", cardFan2Bytes)
	decode("OpenPackage1", cardOpenPackage1Bytes)
	decode("OpenPackage2", cardOpenPackage2Bytes)
	decode("Place1", cardPlace1Bytes)
	decode("Place2", cardPlace2Bytes)
	decode("Place3", cardPlace3Bytes)
	decode("Place4", cardPlace4Bytes)
	decode("Shove1", cardShove1Bytes)
	decode("Shove2", cardShove2Bytes)
	decode("Shove3", cardShove3Bytes)
	decode("Shove4", cardShove4Bytes)
	decode("Shuffle", cardShuffleBytes)
	decode("Slide1", cardSlide1Bytes)
	decode("Slide2", cardSlide2Bytes)
	decode("Slide3", cardSlide3Bytes)
	decode("Slide4", cardSlide4Bytes)
	decode("Slide5", cardSlide5Bytes)
	decode("Slide6", cardSlide6Bytes)
	decode("Slide7", cardSlide7Bytes)
	decode("Slide8", cardSlide8Bytes)
	decode("TakeOutPackage1", cardTakeOutPackage1Bytes)
	decode("TakeOutPackage2", cardTakeOutPackage2Bytes)
	decode("Complete", completeBytes)
	decode("Blip", blipBytes)
}

var soundRandomizer = map[string]int{
	"Fan":            2,
	"OpenPackage":    2,
	"Place":          4,
	"Shove":          4,
	"Shuffle":        0,
	"Slide":          8,
	"TakeOutPackage": 2,
	"Complete":       0,
	"Blip":           0,
}

func SetVolume(vol float64) {
	Volume = vol
}

func Play(name string) {
	if Volume == 0.0 {
		return
	}
	n := soundRandomizer[name]
	// if !ok {
	// println(name, " not found in sound randomizer")
	// caller may have passed a full/specific name, eg Slide1, so try to play that
	// }
	var fullName string
	if n == 0 {
		fullName = name
	} else {
		// rand.Intn(8) produces 0..7 inclusive
		fullName = fmt.Sprintf("%s%d", name, rand.Intn(n)+1)
	}
	audioPlayer, ok := soundMap[fullName]
	if !ok {
		log.Panic(fullName, " not found in sound map")
	}
	if !audioPlayer.IsPlaying() {
		audioPlayer.Rewind()
		audioPlayer.SetVolume(Volume)
		audioPlayer.Play()
	}
}
