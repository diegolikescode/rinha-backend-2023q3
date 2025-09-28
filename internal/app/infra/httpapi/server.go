package httpapi

import (
	"github.com/gofiber/fiber/v3/log"
	"honnef.co/go/tools/config"
)

func startProfiling(config *config.Config) {
	log.Infof("Starting CPU and MEM profiling on %s and %s", config.Profiling.CPU, config.Profiling.Mem)
}
