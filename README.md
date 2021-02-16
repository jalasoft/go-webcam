# go-webcam

Simple API that allows managing webcameras on Linux system. It takes advantage of Video For Linux version 2 API (aka v4l2).

The API offers following operations:

* Finding all available webcams
* Querying capabilities of a webcam
* Querying supported formats of pictures
* Querying supported frame sizes of picures
* Taking a snapshot for given frame size

### Installing module

```bash
mkdir my-module
cd my-module
go mod init my-module
go get github.com/jalasoft/go-webcam
```

### Example of probing all available video devices

```go
devices, err := webcam.FindWebcams()

if err != nil {
    log.Fatal(err)
}

for _, device := range devices {
    log.Printf("%v\n", device)
}
```

### Example of taking snapshot

```go
//open webcam
dev, err := webcam.OpenWebcam("/dev/video0")

if err != nil {
	log.Fatal(err)
}

defer dev.Close()

//query all supported formats
formats, err := dev.QueryFormats()

if err != nil {
    log.Fatal(err)
}

for _, f := range formats {
	log.Printf("format: %v\n", f)
}

//query all supported framesizes for given format
fs, err := dev.QueryFrameSizes(formats[0])

if err != nil {
	log.Fatal(err)
}

//the API now supports querying discrete and stepwise formats
for _, d := range fs.Discrete() {
	log.Printf("%v\n", d)
}

for _, s := range fs.Stepwise() {
    log.Printf("%v\n", s)
}

//take a picture with given resolution, the API now supports
//just discrete frame sizes, which is most often used
s, err := dev.TakeSnapshot(&webcam.DiscreteFrameSize{PixelFormat: formats[0], Width: 640, Height: 480})
if err != nil {
	log.Fatal(err)
}

//Data() method of the snapshot now provides bytes of the picture
ioutil.WriteFile("/home/me/picture.jpg", s.Data(), 0644)
```