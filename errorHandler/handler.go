package errorHandler

import "log"

func GetError(err error, msg string) {
	if err != nil {
		log.Panicf("%s, %s", msg, err)
	}
}
