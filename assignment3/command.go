package main

import "fmt"

type MusicPlayer struct {
	isPlaying bool
	song      string
}

func (m *MusicPlayer) play(song string) {
	m.song = song
	m.isPlaying = true
	fmt.Printf("Playing %s\n", song)
}

func (m *MusicPlayer) pause() {
	if m.isPlaying {
		fmt.Printf("Pausing %s\n", m.song)
		m.isPlaying = false
	}
}

type Command interface {
	execute()
}

type PlayCommand struct {
	player *MusicPlayer
	song   string
}

func (pc *PlayCommand) execute() {
	pc.player.play(pc.song)
}

type PauseCommand struct {
	player *MusicPlayer
}

func (pc *PauseCommand) execute() {
	pc.player.pause()
}

type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) pressButton() {
	rc.command.execute()
}

func main() {
	player := &MusicPlayer{}
	playCommand := &PlayCommand{player, "Song Name"}
	pauseCommand := &PauseCommand{player}

	remote := &RemoteControl{}

	remote.command = playCommand
	remote.pressButton()

	remote.command = pauseCommand
	remote.pressButton()
}
