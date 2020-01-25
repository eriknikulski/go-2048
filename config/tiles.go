package config

import (
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
	"os"
)

var Color = map[string]color.RGBA{
	"0":     {0xcd, 0xc1, 0xb4, 0xff},
	"2":     {0xee, 0xe4, 0xda, 0xff},
	"4":     {0xed, 0xe0, 0xc8, 0xff},
	"8":     {0xf2, 0xb1, 0x79, 0xff},
	"16":    {0xf5, 0x95, 0x63, 0xff},
	"32":    {0xf6, 0x7c, 0x5f, 0xff},
	"64":    {0xf6, 0x5e, 0x3b, 0xff},
	"128":   {0xed, 0xcf, 0x72, 0xff},
	"256":   {0xed, 0xcc, 0x61, 0xff},
	"512":   {0xed, 0xc8, 0x50, 0xff},
	"1024":  {0xed, 0xc5, 0x3f, 0xff},
	"2048":  {0xed, 0xc2, 0x2e, 0xff},
	"4096":  {0xee, 0xe4, 0xda, 0xff},
	"8192":  {0xed, 0xc2, 0x2e, 0xff},
	"16384": {0xf2, 0xb1, 0x79, 0xff},
	"32768": {0xf5, 0x95, 0x63, 0xff},
	"65536": {0xf6, 0x7c, 0x5f, 0xff},
}

var FontColor = map[string]color.RGBA{
	"0":     {0, 0, 0, 0},
	"2":     {0x77, 0x6e, 0x65, 0xff},
	"4":     {0x77, 0x6e, 0x65, 0xff},
	"8":     {0xf9, 0xf6, 0xf2, 0xff},
	"16":    {0xf9, 0xf6, 0xf2, 0xff},
	"32":    {0xf9, 0xf6, 0xf2, 0xff},
	"64":    {0xf9, 0xf6, 0xf2, 0xff},
	"128":   {0xf9, 0xf6, 0xf2, 0xff},
	"256":   {0xf9, 0xf6, 0xf2, 0xff},
	"512":   {0xf9, 0xf6, 0xf2, 0xff},
	"1024":  {0xf9, 0xf6, 0xf2, 0xff},
	"2048":  {0xf9, 0xf6, 0xf2, 0xff},
	"4096":  {0x77, 0x6e, 0x65, 0xff},
	"8192":  {0xf9, 0xf6, 0xf2, 0xff},
	"16384": {0x77, 0x6e, 0x65, 0xff},
	"32768": {0x77, 0x6e, 0x65, 0xff},
	"65536": {0xf9, 0xf6, 0xf2, 0xff},
}

var FontSize = map[string]float64{
	"0":     0,
	"2":     55,
	"4":     55,
	"8":     55,
	"16":    55,
	"32":    55,
	"64":    55,
	"128":   45,
	"256":   45,
	"512":   45,
	"1024":  35,
	"2048":  35,
	"4096":  30,
	"8192":  30,
	"16384": 30,
	"32768": 30,
	"65536": 30,
}

var Atlas = map[string]*text.Atlas{
	"0":     nil,
	"2":     nil,
	"4":     nil,
	"8":     nil,
	"16":    nil,
	"32":    nil,
	"64":    nil,
	"128":   nil,
	"256":   nil,
	"512":   nil,
	"1024":  nil,
	"2048":  nil,
	"4096":  nil,
	"8192":  nil,
	"16384": nil,
	"32768": nil,
	"65536": nil,
}

func loadTTF(path string) (*truetype.Font, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return font, nil
}

func getFontFace(font *truetype.Font, size float64) font.Face {
	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})
}

func init() {
	tileFont, err := loadTTF(FontName)
	if err != nil {
		panic(err)
	}

	for k := range Atlas {
		face := getFontFace(tileFont, FontSize[k])
		Atlas[k] = text.NewAtlas(face, text.ASCII)
	}
}
