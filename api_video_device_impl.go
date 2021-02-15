package webcam

import (
	"log"
	"os"
)

type device struct {
	file *os.File
}

func (d *device) File() *os.File {
	return d.file
}

func (d *device) Close() error {
	log.Printf("Closing video device.\n")
	return d.file.Close()
}
