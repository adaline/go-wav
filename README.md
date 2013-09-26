go-wav
======

A Go library for reading and writing WAV files

Example:

```go
package main

import (
  "go-wav"
)
func main() {
  
  filename := "/tmp/test_input.wav"

  header := wav.ReadHeader(filename)
  header.PrintHeader()

	// Going to read 1234 samples
  read_length := 1234

  // We double the buffer size, because stereo has 2 values per sample
  data := make([]int16, read_length * 2)

	// Read the first 1234 samples
  header.ReadData(filename, 0, data)

  // Write a 1234 samples to an output file
  wav.Write("/tmp/test_output.wav", int(header.SampleRate), data)
}
```