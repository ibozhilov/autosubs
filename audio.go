package autosubs

import (
	"fmt"
	"os/exec"
)

func VideoToAudio(videoPath, audioPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vn", audioPath)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("VideoToAudio")
	return nil
}

func WavToFLAC(wavPath, flacPath string) error {
	cmd := exec.Command("ffmpeg", "-i", wavPath, "-ab", "16k", flacPath)
	err := cmd.Run()
	if err != nil {
		return nil
	}
	fmt.Println("WavToFlac")
	return nil
}

func ToFlac(videoPath, flacPath string) error {
	fmt.Println("in")
	err := VideoToAudio(videoPath, "/home/ivan/Desktop/Temp.wav")
	if err == nil {
		err = WavToFLAC("/home/ivan/Desktop/Temp.wav", flacPath)
		if err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
