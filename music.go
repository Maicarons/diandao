package main

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

func (g *GPlayer) BGMStateChange() {
	if g.bgmPlayer.IsPlaying() {
		g.bgmPlayer.Pause()
	} else {
		g.bgmPlayer.Play()
	}
}
func (g *GPlayer) BGMBegain() error {
	const sampleRate = 44100
	a  := audio.NewContext(44100)
	g.AudioContext = a
	{
		f, err := ResourceFS.Open("bgm.ogg")
		if err != nil {
			return err
		}
		defer f.Close()
		decoded, err := vorbis.DecodeWithSampleRate(sampleRate, f)
		if err != nil {
			return err
		}
		loop := audio.NewInfiniteLoop(decoded, decoded.Length())
		p, err := g.AudioContext.NewPlayer(loop)
		if err != nil {
			return err
		}
		g.bgmPlayer = p
		g.bgmPlayer.SetVolume(0.8)
		g.bgmPlayer.Play()
	}
	return nil
}
