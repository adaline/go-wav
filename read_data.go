package wav

import (
  bin "encoding/binary"
  "os"
  "log"
  "bytes"
)

func (wav Wav) ReadData(FileName string, start int, data []int16){
  offset := int64((28 + wav.Subchunk1Size) + uint32(start * int(wav.NumChannels) * int(wav.BitsPerSample/8)))

  file, open_err := os.OpenFile(FileName, os.O_RDONLY, 0)
  if open_err != nil {
    log.Fatal( "Error opening file for reading\n" )
  }
  buffer := make([]byte, len(data) * 4)
  file.ReadAt(buffer, offset)

  reader := bytes.NewReader(buffer)

  read_err := bin.Read(reader, bin.LittleEndian, data )
  if read_err != nil {
    log.Fatal("binary.Read failed:", read_err)
  }
  file.Close()

  return
}