package core

import (
	"github.com/go-logr/logr"
)

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.V(2).Info("Debug: level 2")
	logger.V(3).Info("Debug: level 3")
	logger.V(4).Info("Debug: level 4")
	logger.V(5).Info("Debug: level 5")
}
