package main

var ergodoxAsciiTable [100]byte

func init() {
	// See table details:
	// TODO: << Add hyperlink. >>
	characters := "" +
		// indecies 0 - 47
		// Append a buffer of three, in case more characters need to be added,
		// and to allow clean addition of 50 when the shift key is pressed.
		" 0123456789abcdefghijklmnopqrstuvwxyz;=,-./`[\\]'  " +
		// indecies 50 - 97 (Shifted versions of the above)
		" )!@#$%^&*(ABCDEFGHIJKLMNOPQRSTUVWXYZ:+<_>?~{|}\"  "

	// Convert to a byte slice because \ has strange behavior and is converted
	// to a unicode point.
	byteCharacters := []byte(characters)

	for i, char := range byteCharacters {
		ergodoxAsciiTable[i] = byte(char)
	}
}

func getErgodoxChar(vkCode uint, isShiftPressed bool) byte {
	// ... Your existing custom ErgoDox character mapping logic ...
	var shiftAdd uint = 0
	if isShiftPressed {
		shiftAdd = 50
	}

	// Decision tree with max depth 6 or so. Better than a switch statement.
	// but this have a constant factor of a single array lookup if we expanded
	// the array size from 100 to about 256. Which we will probably do.
	// This is just an exercise in memory efficiency.
	if vkCode >= 65 {
		if vkCode <= 90 {
			return ergodoxAsciiTable[vkCode-54+shiftAdd]
		} else {
			if vkCode >= 186 && vkCode <= 222 {
				if vkCode <= 192 {
					return ergodoxAsciiTable[vkCode-149+shiftAdd]
				} else {
					return ergodoxAsciiTable[vkCode-175+shiftAdd]
				}
			}
		}
	} else {
		if vkCode == 32 {
			return ergodoxAsciiTable[0] // don't bother with shiftAdd
		} else {
			if vkCode >= 48 && vkCode <= 57 {
				return ergodoxAsciiTable[vkCode-47+shiftAdd]
			}
		}
	}

	return 0
}
