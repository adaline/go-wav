package wav

import (
  bin "encoding/binary"
  "os"
  "log"
)

func Write(fn string, sample_rate int, data []int16){
  file, err := os.OpenFile(fn, os.O_WRONLY | os.O_CREATE, 0666)
  if err != nil {
    log.Fatal( "Error opening file for writing: %v\n", err )
  }

  var header Wav

  header.ChunkID = [4]byte{'R', 'I','F', 'F'}
  header.Format = [4]byte{'W','A','V','E'} 
  header.Subchunk1ID = [4]byte{'f','m','t',' '}
  header.Subchunk1Size = 16
  header.AudioFormat = 1 // 1 == PCM
  header.NumChannels = 2 // 2 == Stereo
  header.SampleRate = uint32(sample_rate)
  header.BitsPerSample = 16 // 16bit integer samples
  header.ByteRate = uint32( int(header.SampleRate) * int(header.NumChannels) * int(header.BitsPerSample/8))
  header.BlockAlign = uint16(int(header.NumChannels) * int(header.BitsPerSample/8))
  header.Subchunk2ID = [4]byte{'d','a','t','a'}

  header.Subchunk2Size = uint32(len(data) * int(header.BitsPerSample/8))
  header.ChunkSize = 4 + (8 + header.Subchunk1Size) + (8 + header.Subchunk2Size)
  
  bin.Write( file, bin.BigEndian, &header.ChunkID )
  bin.Write( file, bin.LittleEndian, &header.ChunkSize )
  bin.Write( file, bin.BigEndian, &header.Format )

  bin.Write( file, bin.BigEndian, &header.Subchunk1ID )
  bin.Write( file, bin.LittleEndian, &header.Subchunk1Size )
  bin.Write( file, bin.LittleEndian, &header.AudioFormat )
  bin.Write( file, bin.LittleEndian, &header.NumChannels )
  bin.Write( file, bin.LittleEndian, &header.SampleRate )
  bin.Write( file, bin.LittleEndian, &header.ByteRate )
  bin.Write( file, bin.LittleEndian, &header.BlockAlign )
  bin.Write( file, bin.LittleEndian, &header.BitsPerSample )


  bin.Write( file, bin.BigEndian, &header.Subchunk2ID )
  bin.Write( file, bin.LittleEndian, &header.Subchunk2Size )

  bin.Write( file, bin.LittleEndian, data )

  file.Close()

  return
}