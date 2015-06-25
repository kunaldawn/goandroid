package logging

import (
	"log"
	"os"
)

const (
	NONE          = ""
	ACTIVE        = "A"
	VERBOSE       = "V"
	VERBOSE_MORE  = "VV"
	VERBOSE_EXTRA = "VVV"
	ENV           = "DEBUG_LOG"
)

func Log(format string, args ...interface{}) {
	env := os.Getenv(ENV)
	if env == NONE {
		return
	}
	if (env == ACTIVE) || (env == VERBOSE) || (env == VERBOSE_MORE) || (env == VERBOSE_EXTRA) {
		log.Printf("[goandroid] : "+format, args...)
	}
}

func LogV(format string, args ...interface{}) {
	env := os.Getenv(ENV)
	if env == NONE {
		return
	}
	if (env == VERBOSE) || (env == VERBOSE_MORE) || (env == VERBOSE_EXTRA) {
		log.Printf("[goandroid] : "+format, args...)
	}
}

func LogVV(format string, args ...interface{}) {
	env := os.Getenv(ENV)
	if env == NONE {
		return
	}
	if (env == VERBOSE_MORE) || (env == VERBOSE_EXTRA) {
		log.Printf("[goandroid] : "+format, args...)
	}
}

func LogVVV(format string, args ...interface{}) {
	env := os.Getenv(ENV)
	if env == NONE {
		return
	}
	if env == VERBOSE_EXTRA {
		log.Printf("[goandroid] : "+format, args...)
	}
}
