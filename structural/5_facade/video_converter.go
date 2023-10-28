package facade

import "fmt"

// Complex video conversion subsystem
type VideoFile struct{}

func (f *VideoFile) Load(fileName string) error {
	fmt.Println("Loading video file...", fileName)
	return nil
}

func (f *VideoFile) CheckFormat(fileName string) error {
	fmt.Println("Checking video file format...", fileName)
	return nil
}

type Codec struct{}

func (c *Codec) Decode(format string) error {
	fmt.Println("Decoding video...", format)
	return nil
}

func (c *Codec) ReduceFrameRate(fps int) error {
	fmt.Println("Reducing framerate...", fps)
	return nil
}

func (c *Codec) AdjustBitrate(br int) error {
	fmt.Println("Adjusting bitrate...", br)
	return nil
}

func (c *Codec) Encode(format string) error {
	fmt.Println("Encoding video...", format)
	return nil
}

type BitrateCalculator struct{}

func (b *BitrateCalculator) Calculate(resolution string) (int, error) {
	fmt.Println("Calculating bitrate for resolution...", resolution)
	return 3000, nil
}

// Facade VideoConverter.Convert(...) hides the complexity of encoding videos
type VideoConverter struct {
	file  *VideoFile
	codec *Codec
	calc  *BitrateCalculator
}

func NewVideoConverter() *VideoConverter {
	vc := &VideoConverter{}
	// init dependencies of `vc`
	// ...
	return vc
}

func (c *VideoConverter) Convert(fileName, resolution, format string) error {
	if err := c.file.Load(fileName); err != nil {
		return fmt.Errorf("load video file fail: %v", err)
	}
	if err := c.file.CheckFormat(fileName); err != nil {
		return fmt.Errorf("check video format fail: %v", err)
	}
	if err := c.codec.Decode(fileName); err != nil {
		return fmt.Errorf("decode video data fail: %v", err)
	}
	if err := c.codec.ReduceFrameRate(30); err != nil {
		return fmt.Errorf("reduce frame rate fail: %v", err)
	}

	bitrate, err := c.calc.Calculate(resolution)
	if err != nil {
		return fmt.Errorf("calculate bit rate fail: %v", err)
	}
	if err := c.codec.AdjustBitrate(bitrate); err != nil {
		return fmt.Errorf("adjust bitrate fail: %v", err)
	}

	if err := c.codec.Encode(format); err != nil {
		return fmt.Errorf("encode video fail: %v", err)
	}
	return nil
}
