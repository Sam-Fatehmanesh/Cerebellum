package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo"
	"go.arsenm.dev/logger"
	"go.arsenm.dev/logger/log"
)

func init() {
	log.Logger = logger.NewPretty(os.Stdout)
}

func main() {
	fmt.Println("start")

	// Initializing malgo context
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		log.Error("Error initializing malgo context").Err(err).Send()
	}

	// Defering malgo context uninitialization to end of main
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	// malgo audio configuration
	deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = uint32(1) // 1 here is temp, replace with required sample rate
	deviceConfig.Alsa.NoMMap = 1

	// Buffer for the incoming audio data
	var pCapturedSamples []byte
	// Counter for frames of audio
	var capturedSampleCount uint32

	// Size of chunk in bytes
	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))

	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {
		// Audio bytes count
		sampleCount := framecount * deviceConfig.Capture.Channels * sizeInBytes
		// Total audio bytes count
		newCapturedSampleCount := capturedSampleCount + sampleCount
		// Updates audio buffer
		pCapturedSamples = append(pCapturedSamples, pSample...)
		// Updates audio buffer bytes count
		capturedSampleCount = newCapturedSampleCount
	}

}
