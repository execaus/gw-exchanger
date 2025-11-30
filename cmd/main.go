package cmd

import "go.uber.org/zap"

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}
