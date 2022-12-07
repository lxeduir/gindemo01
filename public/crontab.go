package public

import (
	"github.com/robfig/cron/v3"
	"time"
)

func Online() {
	crontab := cron.New(cron.WithSeconds())
	_, err := crontab.AddFunc("0 0 */5 * * ?", func() {
		var toemail []string
		toemail = append(toemail, "2508339002@qq.com")
		Email(toemail, "定时任务"+time.Now().Format("2006-01-02 15:04:05"))
	},
	)
	if err != nil {
		return
	}
	crontab.Start()
}
