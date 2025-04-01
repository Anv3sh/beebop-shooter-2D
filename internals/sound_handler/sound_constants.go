package soundhandler

import (
	"embed"
)

//go:embed sound_effects/*
var soundEffects embed.FS


// sound effects
const(
	LASER_HIT_SOUND = "sound_effects/hit1.mp3"
	BGM = "sound_effects/bgm.mp3"
)