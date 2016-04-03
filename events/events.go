package events

import "fmt"

type Button struct {
	eventListeners map[string][]chan string
}

func InitEventsFactory()  {
	btn := MakeButton()

	handlerTurnOnVolume := make(chan string)
	handlerTurnOffVolume := make(chan string)

	btn.AddEventListener("volume", handlerTurnOnVolume)
	btn.AddEventListener("volume", handlerTurnOffVolume)
	btn.AddEventListener("volumeOff", handlerTurnOffVolume)

	go func() {
		for {
			msg := <- handlerTurnOnVolume
			fmt.Println("Handler turned on: " + msg)
		}
	}()

	go func() {
		for {
			msg := <- handlerTurnOffVolume
			fmt.Println("Handler turned off: " + msg)
		}
	}()

	btn.PushTheButton("volume", "Volume turned on!")
	btn.RemoveEventListener("volume", handlerTurnOnVolume)
	btn.PushTheButton("volume", "Volume turned off!")
	btn.RemoveEventListener("volume", handlerTurnOffVolume)

	btn.PushTheButton("volumeOff", "VolumeOFF is pushed on")

	fmt.Scanln()
}

func MakeButton() *Button {
	button := new(Button)
	button.eventListeners = make(map[string][]chan string)

	return button
}

func (button *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := button.eventListeners[event]; present {
		// because button.eventListeners[event] it's slice
		button.eventListeners[event] = append(button.eventListeners[event], responseChannel)
	} else {
		button.eventListeners[event] = []chan string{responseChannel}
	}
}

func (button *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := button.eventListeners[event]; present {
		for index, _ := range button.eventListeners[event] {
			if button.eventListeners[event][index] == listenerChannel {
				button.eventListeners[event] = append(button.eventListeners[event][:index],
					button.eventListeners[event][index+1:]...)
				break
			}
		}
	}
}

func (button *Button) PushTheButton(event, response string) {
	if _, present := button.eventListeners[event]; present {
		for _, handler := range button.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}