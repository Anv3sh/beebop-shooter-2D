package internals

import (
	"io"
	"embed"
	"github.com/hajimehoshi/oto"
	"log"
	"github.com/hajimehoshi/go-mp3"
)

//go:embed sounds/*
var soundsEffects embed.FS

func playSoundEffect(soundEffect string) error {
	f, err := soundsEffects.Open(soundEffect)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

        // fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

func MustPlay(soundEffect string){
	if err:=playSoundEffect(LASER_HIT_SOUND); err!=nil{
		log.Fatal(err)
	}
}
