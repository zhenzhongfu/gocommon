package captcha

import (
	"image/color"
	"io/ioutil"

	"github.com/afocus/captcha"
)

var cap *captcha.Captcha

func GetInstance() *captcha.Captcha {
	return cap
}

func Setup() {
	cap = captcha.New()
	/*
		if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
		}
	*/
	//We can load font not only from localfile, but also from any []byte slice
	fontContenrs, err := ioutil.ReadFile("/usr/share/fonts/comic.ttf")
	if err != nil {
		panic(err.Error())
	}
	err = cap.AddFontFromBytes(fontContenrs)
	if err != nil {
		panic(err.Error())
	}

	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetBkgColor(color.RGBA{255, 255, 255, 255})
	cap.SetFrontColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
}
