package main

type Buffer struct {
	// Keeping all these lower case to hide the circular buffer
	// implementation and keep the package interface clean.
	wheel      []byte // Only storing printable ASCII characters.
	writePoint int
	slashes    []int
	standby    bool
}

// For initializing new instances of the Buffer type.
func NewBuffer(size int) *Buffer {
	return &Buffer{
		wheel:      make([]byte, size),
		writePoint: 0, // I suppose you could start anywhere in the buffer.
		slashes:    []int{},
		standby:    true,
	}
}

func (b *Buffer) Write(char byte) {

	// If there are no slashes in the buffer, and the recieved character
	// is not a "/" (signifying the start of a new expansion), then do nothing.
	// The program hopes to spend most of its time in this state.
	if b.standby && char != '/' {
		return
	}

	// Delete slashes out of the front of the slashes array as you overwrite
	// them in the circular buffer. If you overwrite the last remaining
	// slash, enter standby mode and stop adding anything into the buffer
	// until the next encountered "/" character (see above.)
	if len(b.slashes) > 0 && b.writePoint == b.slashes[0] {
		b.slashes = b.slashes[1:]
		if len(b.slashes) == 0 {
			b.standby = true // Enter standby mode.
		}
	}

	// Write the character to the buffer
	b.wheel[b.writePoint] = char

	// Push any slashes to the slashes queue.
	if char == '/' {
		b.standby = false // (Possibly redundantly) exit standby mode.
		b.slashes = append(b.slashes, b.writePoint)
	}

	// Update the write point on the circular buffer wheel.
	b.writePoint = (b.writePoint + 1) % len(b.wheel)

}

func (b *Buffer) DeleteChar() {
	// Don't bother to delete anything if we are in standby mode.
	if b.standby {
		return
	}

	// Simply shift the write point back without actually deleting anything.
	b.writePoint--
	if b.writePoint < 0 {
		b.writePoint = len(b.wheel) - 1
	}

	// When passing over a slash, delete it off the back of the slashes index.
	if len(b.slashes) > 0 && b.slashes[len(b.slashes)-1] == b.writePoint {
		if len(b.slashes) == 1 {
			b.slashes = []int{}
			b.standby = true
		} else {
			b.slashes = b.slashes[:len(b.slashes)-1]
		}
	}
}

func (b *Buffer) DebugRead() string {
	// Combine the two parts of the circular buffer into a single string.
	return string(b.wheel[b.writePoint:]) +
		string(b.wheel[:b.writePoint])
}

func (b *Buffer) Read() []string {
	// Return an array of all possible strings beginning with a slash.

	// starting from the character behind writePoint, and working backwards.
	// First collect the circle buffer into a string ending just
	// before the writePoint.
	readString := string(b.wheel[b.writePoint:]) +
		string(b.wheel[:b.writePoint])

	// Build array strSlashes and then rotate it in exactly the same way.
	strSlashes := make([]int, len(b.slashes))
	for i, slashIndex := range b.slashes {
		strSlashes[i] = (slashIndex - b.writePoint + len(b.wheel)) % len(b.wheel)
	}

	// Build the array of strings going back to and including slash characters
	// of increasing distance from the fixed end of the readable string.
	output := []string{}
	for i := len(strSlashes) - 1; i >= 0; i-- {
		output = append(output, readString[strSlashes[i]:len(readString)])
	}
	return output

}

func (b *Buffer) Clear() {
	// Simply clear the stored slashes array and enable standby.
	// The writePoint can remain where it is, and anything written to the
	//  wheel can now just be written over.
	b.slashes = []int{}
	b.standby = true
}
