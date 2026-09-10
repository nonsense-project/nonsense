package prefixmanager

import (
	"github.com/nonsense-project/nonsense/v2/infrastructure/logger"
	"github.com/nonsense-project/nonsense/v2/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
