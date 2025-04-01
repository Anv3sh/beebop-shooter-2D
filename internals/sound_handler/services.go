package soundhandler

import (
	"io"
	"log"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"sync"
	"time"
)

var (
	context     *oto.Context
	contextOnce sync.Once
)

func InitSoundSystem() {
	contextOnce.Do(func() {
		var err error
		context, err = oto.NewContext(46000, 2, 2, 8192) // standard CD-quality audio
		if err != nil {
			log.Fatal("Failed to initialize audio context:", err)
		}
	})
}

func playSoundEffect(soundEffect string) error {
	f, err := soundEffects.Open(soundEffect)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	p := context.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

// func playSoundEffectInLoop(soundEffect string) error {
// 	f, err := soundEffects.Open(soundEffect)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	d, err := mp3.NewDecoder(f)
// 	if err != nil {
// 		return err
// 	}
// 	for{
// 		p := context.NewPlayer()
// 		defer p.Close()

// 			// fmt.Printf("Length: %d[bytes]\n", d.Length())

// 		if _, err := io.Copy(p, d); err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }
var bgmStopChan = make(chan struct{})

func PlayBGM(path string) {
	go func() {
		for {
			select {
			case <-bgmStopChan:
				return
			default:
				f, err := soundEffects.Open(path)
				if err != nil {
					log.Println("bgm open error:", err)
					return
				}
				d, err := mp3.NewDecoder(f)
				if err != nil {
					log.Println("bgm decode error:", err)
					f.Close()
					return
				}

				p := context.NewPlayer()
				if _, err := io.Copy(p, d); err != nil {
					log.Println("bgm stream error:", err)
				}

				p.Close()
				f.Close()

				// slight pause between loops to avoid pop/click
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
}

func StopBGM() {
	bgmStopChan <- struct{}{}
}

func MustPlay(soundEffect string, loop bool){
	InitSoundSystem()
	if loop{
		PlayBGM(soundEffect)
	}else{
		if err:=playSoundEffect(soundEffect); err!=nil{
			log.Fatal(err)
		}
	}
}
