package main

import (
	"Twipex_project/image_generation"

	"github.com/carlescere/scheduler"
)

func SetSchedule() {
	jobam0 := func() { image_generation.PostImage("am0") }
	jobam3 := func() { image_generation.PostImage("am3") }
	jobam6 := func() { image_generation.PostImage("am6") }
	jobam9 := func() { image_generation.PostImage("am9") }
	jobpm0 := func() { image_generation.PostImage("pm0") }
	jobpm3 := func() { image_generation.PostImage("pm3") }
	jobpm6 := func() { image_generation.PostImage("pm6") }
	jobpm9 := func() { image_generation.PostImage("pm9") }
	jobweek := func() { image_generation.PostImage("week") }

	scheduler.Every().Day().At("00:00:05").Run(jobam0)
	scheduler.Every().Day().At("03:00:05").Run(jobam3)
	scheduler.Every().Day().At("06:00:05").Run(jobam6)
	scheduler.Every().Day().At("09:00:05").Run(jobam9)
	scheduler.Every().Day().At("12:00:05").Run(jobpm0)
	scheduler.Every().Day().At("15:00:05").Run(jobpm3)
	scheduler.Every().Day().At("18:00:05").Run(jobpm6)
	scheduler.Every().Day().At("21:00:05").Run(jobpm9)
	scheduler.Every().Monday().At("00:01:05").Run(jobweek)

}
