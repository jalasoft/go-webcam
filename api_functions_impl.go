package webcam

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	//"github.com/jalasoft/go-webcam/v4l2"
)

func openVideoDevice(path string) (VideoDevice, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)

	log.Printf("Opening device %s\n", path)

	if err != nil {
		return nil, err
	}

	log.Println("Reading capability")

	dev := &device{file}

	caps, err := dev.QueryCapabilities()

	if err != nil {
		return nil, err
	}

	//TODO close when not supported
	if !caps.HasCapability(CAP_VIDEO_CAPTURE) {
		return nil, errors.New(fmt.Sprintf("Device %s is not a video capturing device.", caps.Card()))
	}

	if !caps.HasCapability(CAP_STREAMING) {
		return nil, errors.New(fmt.Sprintf("Device %s is not able to stream frames.", caps.Card()))
	}

	log.Printf("Device %s is a video device", file.Name())
	return dev, nil
}

func searchVideoDevices() ([]string, error) {

	files, error := filepath.Glob("/dev/video*")

	if error != nil {
		return nil, error
	}

	channel := make(chan VideoDevice)
	go probeDevices(files, channel)

	valid := []string{}

	for device := range channel {
		valid = append(valid, device.File().Name())
	}

	return valid, nil
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

	device, error := openVideoDevice(file)

	defer wg.Done()

	if error != nil {
		log.Printf("Device %s is not a camera\n", file)
		return
	}

	defer func() {
		if err := device.Close(); err != nil {
			log.Printf("Cannot close device '%s'", device.File().Name())
		}
	}()

	ch <- device
}

func convertToSlice(capToName map[string]string) []string {
	slice := make([]string, 0, len(capToName))
	for _, v := range capToName {
		slice = append(slice, v)
	}
	return slice
}
