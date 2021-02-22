# go-webcam

Simple API that allows managing webcameras on Linux system. It takes advantage of Video For Linux version 2 API (aka v4l2).

The API offers following operations:

* Finding all available webcams
* Querying capabilities of a webcam
* Querying supported formats of pictures
* Querying supported frame sizes of picures
* Taking a snapshot for given frame size
* Streaming series of snapshots 

### Installing module

```bash
go get -u github.com/jalasoft/go-webcam
```

### API description

1. The first step is to obtain an instance of webcam.Webcam and to defer closing it:
   ```go
    webcam, err := webcam.OpenWebcam("/dev/video0")
	
	if err != nil {
	    //handle
	}

	defer func() {
		if err := webcam.Close(); err != nil {
			//handle
		} 
	}()
   ```

2. Usual usage of this API is to either take a snapshot or to stream snapshots. Both cases requires an instance of webcam.DiscreteFrameSize. To have one, there are two ways:
   * call __webcam.Webcam.QueryFormats()__ and then based on selected *webcam.PixelFormat* invoke __webcam.Webcam.QueryFrameSizes()__
   ```go
       //cam is an instance of webcam.Webcam
       formats, err := cam.QueryFormats()

       if err != nil {
          //handle
       }

	   //select desired pixel format, the most usuall and 
	   //wide spread is V4L2_PIX_FMT_MJPEG
       fs, err := cam.QueryFrameSizes(formats[0])

        if err != nil {
            //handle
        }

        for _, d := range fs.Discrete() {
			//iterate over all supported discrete frame sizes
			//and select the desired one
        }
   ```
   * use "automatic mode" that tries to select pixel format, width and height automatically, still allowing to set some of the paramters.
	```go
	//cam is an instance of webcam.Webcam
	frmSize, err := cam.DiscreteFrameSize()
		.PixelFormatName("V4L2_PIX_FMT_MJPEG")
		.PixelFormat(<existing instance webcam.PixelFormat>)
		.Width(640)
		.Height(480)
		.Select()
	```
	You can omit any (or both) of methods __Width()__, __Height()__, use either __PixelFormatName()__ or __PixelFormat()__, or none of them and the most appriproate pixel format will be choosen automatically. 



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