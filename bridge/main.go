package main

import "fmt"

// --- Implementor ---
type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	GetVolume() int
	SetVolume(volume int)
}

// --- ConcreteImplementor TV ---
type TV struct {
	enabled bool
	volume  int
}

func (t *TV) IsEnabled() bool {
	return t.enabled
}

func (t *TV) Enable() {
	t.enabled = true
	fmt.Println("TV is turned ON")
}

func (t *TV) Disable() {
	t.enabled = false
	fmt.Println("TV is turned OFF")
}

func (t *TV) GetVolume() int {
	return t.volume
}

func (t *TV) SetVolume(volume int) {
	t.volume = volume
	fmt.Printf("TV volume set to %d\n", volume)
}

// --- ConcreteImplementor Radio ---
type Radio struct {
	enabled bool
	volume  int
}

func (r *Radio) IsEnabled() bool {
	return r.enabled
}

func (r *Radio) Enable() {
	r.enabled = true
	fmt.Println("Radio is turned ON")
}

func (r *Radio) Disable() {
	r.enabled = false
	fmt.Println("Radio is turned OFF")
}

func (r *Radio) GetVolume() int {
	return r.volume
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
	fmt.Printf("Radio volume set to %d\n", volume)
}

// --- Abstraction ---
type Remote struct {
	device Device
}

func (r *Remote) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *Remote) VolumeDown() {
	volume := r.device.GetVolume()
	r.device.SetVolume(volume - 10)
}

func (r *Remote) VolumeUp() {
	volume := r.device.GetVolume()
	r.device.SetVolume(volume + 10)
}

// === Refined Abstraction ===
type AdvancedRemote struct {
	Remote
}

func (a *AdvancedRemote) Mute() {
	a.device.SetVolume(0)
	fmt.Println("Device is muted")
}

// === Main ===
func main() {
	tv := &TV{}
	radio := &Radio{}

	fmt.Println("--- TV Remote ---")
	tvRemote := &AdvancedRemote{Remote{device: tv}}
	tvRemote.TogglePower()
	tvRemote.VolumeUp()
	tvRemote.Mute()
	tvRemote.TogglePower()

	fmt.Println("\n--- Radio Remote ---")
	radioRemote := &AdvancedRemote{Remote{device: radio}}
	radioRemote.TogglePower()
	radioRemote.VolumeUp()
	radioRemote.VolumeUp()
	radioRemote.Mute()
}
