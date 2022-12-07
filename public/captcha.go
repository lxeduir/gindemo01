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
	r = r % 1000000
	if r < 100000 {
		r += 100000
	}
	rot = rot + strconv.Itoa(r)
	return rot
}
