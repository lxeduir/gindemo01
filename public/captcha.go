package public

import (
	"math/rand"
	"strconv"
	"time"
)

func Captcha() string {
	rot := ""
	rand.Seed(time.Now().Unix())
	r := rand.Int()
	rot = rot + strconv.Itoa(r%100000)
	return rot
}
