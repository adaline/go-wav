package wav

import (
  bin "encoding/binary"
  "os"
  "log"
)

func ReadHeader( fn string ) (wav Wav) {
  file, err := os.OpenFile(fn, os.O_RDONLY, 0)
  if err != nil {
    log.Fatal( "Error opening\n" )
  }

  // Read the wave header
  bin.Read( file, bin.BigEndian, &wav.ChunkID )
  bin.Read( file, bin.LittleEndian, &wav.ChunkSize )
  bin.Read( file, bin.BigEndian, &wav.Format )

  // Read the format header
  bin.Read( file, bin.BigEndian, &wav.Subchunk1ID )
  bin.Read( file, bin.LittleEndian, &wav.Subchunk1Size )
  bin.Read( file, bin.LittleEndian, &wav.AudioFormat )
  bin.Read( file, bin.LittleEndian, &wav.NumChannels )
  bin.Read( file, bin.LittleEndian, &wav.SampleRate )
  bin.Read( file, bin.LittleEndian, &wav.ByteRate )
  bin.Read( file, bin.LittleEndian, &wav.BlockAlign )
  bin.Read( file, bin.LittleEndian, &wav.BitsPerSample )

  if wav.Subchunk1Size > 16 {
    BytesToRead := wav.Subchunk1Size - 16
    var ReadByte byte
    for i := 0; i < int(BytesToRead); i++ {
      bin.Read(file, bin.LittleEndian, &ReadByte )
    }
  }


  bin.Read( file, bin.BigEndian, &wav.Subchunk2ID )
  bin.Read( file, bin.LittleEndian, &wav.Subchunk2Size )


  if (wav.ChunkID != [4]byte{'R', 'I','F', 'F'}) ||
  (wav.Format != [4]byte{'W','A','V','E'}) ||
  (wav.Subchunk1ID != [4]byte{'f','m','t',' '}) ||
  (wav.Subchunk2ID != [4]byte{'d','a','t','a'}) {
    log.Fatal("Wave file format incorrect\n")
  }

  wav.NumSamples = uint32(wav.Subchunk2Size / uint32(wav.NumChannels) / uint32(wav.BitsPerSample/8))
  
  file.Close()

  return
}