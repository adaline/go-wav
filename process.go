package wav

// Assumes 2 channel interleaved layout
func SplitChannel(data []int16, channel int, output_size int, output []int16){
  var offset int
  for i := 0; i < output_size; i++ {
    offset = i * 2 + channel
    output[i] = data[offset]
  }
  return
}