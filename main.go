package main

import (
	"fmt"
	"strings"

	"log"
	"os"
	"os/signal"
	"time"

	"github.com/atotto/clipboard"
	"github.com/micmonay/keybd_event"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
)

const (
	VK_LSHIFT   = 0xA0
	VK_RSHIFT   = 0xA1
	VK_BACK     = 0x08
	VK_LCONTROL = 0xA2
	VK_RCONTROL = 0xA3
)

var (
	leftShiftPressed  bool
	rightShiftPressed bool
	leftCtrlPressed   bool
	rightCtrlPressed  bool
)

func main() {
	fmt.Println("Running. Press Ctrl+C to exit")

	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	buf := NewBuffer(55)
	// buf.Write('/')
	// buf.Write('h')
	// buf.Write('e')
	// buf.Write('l')
	// buf.Write('l')

	getChar := getErgodoxChar
	// Buffer size is depends on your need. The 100 is placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("start capturing keyboard input")

	for {
		select {
		case <-time.After(25 * time.Minute):
			fmt.Println("Received timeout signal")
			return nil
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return nil
		case k := <-keyboardChan:
			// DEBUG LINE
			// fmt.Printf("Received %v %v\n", k.Message, k.VKCode)
			// Update shift key states
			if k.VKCode == VK_LSHIFT {
				leftShiftPressed = k.Message == types.WM_KEYDOWN
			} else if k.VKCode == VK_RSHIFT {
				rightShiftPressed = k.Message == types.WM_KEYDOWN

			} else if k.VKCode == VK_LCONTROL {
				leftCtrlPressed = k.Message == types.WM_KEYDOWN
			} else if k.VKCode == VK_RCONTROL {
				rightCtrlPressed = k.Message == types.WM_KEYDOWN

			} else if k.Message == types.WM_KEYDOWN {
				// fmt.Printf("Event: %+v\n", k)
				if k.VKCode == VK_BACK {
					if leftCtrlPressed || rightCtrlPressed {
						// fmt.Printf("Recieved word deletion macro." +
						//	"Clearing buffer \n")
						buf.Clear()
					} else {
						// fmt.Printf("Recieved backspace. \n")
						buf.DeleteChar()
						// if !buf.standby {
						// 	fmt.Println(buf.DebugRead())
						// }
					}					

				} else {
					shiftPressed := leftShiftPressed || rightShiftPressed
					charByte := getChar(uint(k.VKCode), shiftPressed)
					if charByte != 0 {
						if charByte == '*' {
							foundExpansion := false
							// fmt.Println("Asterisk key pressed")
							expandedTexts := buf.Read()

							for _, text := range expandedTexts {
								trimmedText := strings.Trim(text, "/") // Remove slashes from the text
								if expansion, ok := ExpansionsMap[trimmedText]; ok {
									// You can perform the expansion here
									fmt.Printf("Expanded %s --> %s\n",
										trimmedText,
										string(expansion))
							
									// Step 1: Simulate backspaces
									// +1 for the "/" in the front
									backspaces := len(trimmedText) + 2
									for i := 0; i < backspaces; i++ {
										kb, err := keybd_event.NewKeyBonding()
										if err != nil {
											fmt.Println("Error creating keybd_event instance:", err)
											return err
										}
										kb.SetKeys(keybd_event.VK_BACK)
										err = kb.Launching()
										if err != nil {
											fmt.Println("Error pressing backspace key:", err)
											return err
										}
										time.Sleep(5 * time.Millisecond) // Add a small delay between backspaces
										err = kb.Release()
										if err != nil {
											fmt.Println("Error releasing backspace key:", err)
											return err
										}
									}
							
									// Step 2: Copy the rune array to the clipboard in a printable format
									err := clipboard.WriteAll(string(expansion))
									if err != nil {
										fmt.Println("Error setting clipboard:", err)
										return err
									}
							
									// Step 3: Paste the value into the active window
									kb, err := keybd_event.NewKeyBonding()
									if err != nil {
										fmt.Println("Error creating keybd_event instance:", err)
										return err
									}
									kb.SetKeys(keybd_event.VK_V)
									kb.HasCTRL(true)
									err = kb.Launching()
									if err != nil {
										fmt.Println("Error pressing 'Ctrl+V' key:", err)
										return err
									}
									time.Sleep(50 * time.Millisecond) // Add a small delay before releasing the key
									err = kb.Release()
									if err != nil {
										fmt.Println("Error releasing 'Ctrl+V' key:", err)
										return err
									}
							
									buf.Clear()
									foundExpansion = true
									break
								}
							}
							if !foundExpansion {
								buf.Write(charByte)

							    // Debug logs
						        // fmt.Printf("Recieved character: %c\n", charByte)
						        // if !buf.standby {
						        // 	fmt.Println(buf.DebugRead())
						        // }
							}
						} else {
							buf.Write(charByte)

							// Debug logs
						    // fmt.Printf("Recieved character: %c\n", charByte)
						    // if !buf.standby {
						    // 	fmt.Println(buf.DebugRead())
						    // }
						}
					}
				}	
			}
		}
	}

	// not reached
	return nil
}
