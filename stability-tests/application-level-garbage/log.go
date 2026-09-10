package main

import (
	"github.com/nonsense-project/nonsense/v2/infrastructure/logger"
	"github.com/nonsense-project/nonsense/v2/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
