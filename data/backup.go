package data

import (
	"fmt"

	"github.com/robfig/cron/v3"
)



func CronTest() {
c := cron.New()
c.AddFunc("* * * * *", func() {fmt.Printf("Bitch waddup")})

c.Start()
}
