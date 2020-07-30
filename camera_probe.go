package webcam

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
)

func SearchVideoDevices() ([]string, error) {

	files, error := filepath.Glob("/dev/video*")

	if error != nil {
		return nil, error
	}

	channel := make(chan VideoDevice)
	go probeDevices(files, channel)

	cameraDevices := make(map[string]string)

	for cameraDevice := range channel {
		capStr := fmt.Sprintf("%v", cameraDevice.Capabilities())
		cameraDevices[capStr] = cameraDevice.Name()

		if err := cameraDevice.Close(); err != nil {
			log.Printf("Cannot be closed %s: %v", cameraDevice.Name(), err)
		}
	}

	names := convertToSlice(cameraDevices)
	return names, nil
}

func probeDevices(files []string, channel chan VideoDevice) {
	wg := sync.WaitGroup{}
	wg.Add(len(files))

	for _, file := range files {
		go probeDevice(file, channel, &wg)
	}

	wg.Wait()
	close(channel)
}

func probeDevice(file string, ch chan VideoDevice, wg *sync.WaitGroup) {

	device, error := OpenVideoDevice(file)

	if error != nil {
		log.Printf("Device %s is not a camera\n", file)
		if err := device.Close(); err != nil {
			log.Printf(" and cannot be gracefully closed: %v\n", err)
		}
		wg.Done()
		return
	}

	ch <- device

	wg.Done()
}

func convertToSlice(capToName map[string]string) []string {
	slice := make([]string, 0, len(capToName))
	for _, v := range capToName {
		slice = append(slice, v)
	}
	return slice
}
