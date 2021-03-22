package request

type CaptchaVerify struct {
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
