package webcam

import (
	"log"
	"fmt"
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
		capStr := fmt.Sprintf("%v", cameraDevice.Capability())
		cameraDevices[capStr] = cameraDevice.Name()

		if err := cameraDevice.Close(); err != nil {
			log.Printf("Cannot be closed %s: %v", cameraDevice.Name(), err)
		}
	}

	result := make([]string, len(cameraDevices))
	log.Printf("******DETECTED CAMERA DEVICES********")
	for k,v := range cameraDevices {
		log.Printf("%s: %s", v, k)
		result = append(result, v)
	}
	log.Printf("*************************************")

	return result, nil
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

	//cap := device.Capability()
	//log.Printf("%s, %s, %s.", cap.Driver(), cap.Card(), cap.BusInfo())

	ch <- device

	wg.Done()
}
