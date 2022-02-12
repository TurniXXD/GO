package functions

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func SaveOnFourHoursPeriod() {
	for time.Now().Format("04") != "00" {
		time.Sleep(1 * time.Minute)
	}
	gocron.Every(4).Hours().Do(fmt.Println(time.Now()))
	<-gocron.Start()
}
