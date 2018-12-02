package projectcars2

import (
	"bytes"
	"encoding/binary"
	"os"
	"unsafe"
)

// DefaultName is Project Cars 2 API name as of Shared Mem ver 9
const DefaultName = "%pcars2%"

/*
func (d *SharedMemory) MemoryHandle() (*SharedMemroy, error) {
  var err error
  err = nil
  f, err := os.OpenFile(DefaultName, os.O_RDONLY, 444)
  if err != nil {
    return d, err
  }
  defer f.Close()

  var sharedData *SharedMemory =
}
*/
// ExtractData attempts to open the Project Cars 2 memory location,
// then pulls data written to the file for processing
func (d SharedMemory) ExtractData() (SharedMemory, error) {
	var err error
	b := make([]byte, d.Size())
	f, err := os.OpenFile(DefaultName, os.O_RDONLY, 444)
	if err != nil {
		return d, err
	}
	defer f.Close()

	_, err = f.Read(b)
	if err != nil {
		return d, err
	}
	buff := bytes.NewBuffer(b)
	err = binary.Read(buff, binary.LittleEndian, &d)
	if err != nil {
		return SharedMemory{}, err
	}

	return d, nil
}

// Size returns the size of our shared memory struct
// Does not utilize golang slice types, theoretically should be safe
// and constant
func (d *SharedMemory) Size() int32 {
	return int32(unsafe.Sizeof(*d))
}
