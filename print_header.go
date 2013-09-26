package wav

import (
  "fmt"
)

func (wav Wav) PrintHeader(){
  fmt.Printf( "\n" )
  fmt.Printf( "ChunkSize: %d\n", wav.ChunkSize )
  fmt.Printf( "\n" )
  fmt.Printf( "Subchunk1Size: %d\n", wav.Subchunk1Size )
  fmt.Printf( "AudioFormat: %d\n", wav.AudioFormat )
  fmt.Printf( "NumChannels: %d\n", wav.NumChannels )
  fmt.Printf( "SampleRate: %d\n", wav.SampleRate )
  fmt.Printf( "ByteRate: %d\n", wav.ByteRate )
  fmt.Printf( "BlockAlign: %d\n", wav.BlockAlign )
  fmt.Printf( "BitsPerSample: %d\n", wav.BitsPerSample )
  fmt.Printf( "\n" )
  fmt.Printf( "Subchunk2Size: %d\n", wav.Subchunk2Size )
  fmt.Printf( "NumSamples: %d\n", wav.NumSamples )
  fmt.Printf( "\n\n" )
}