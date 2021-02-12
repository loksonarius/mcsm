package server

import (
	"github.com/loksonarius/mcsm/internal/logging"
)

var Log logging.Logger

func init() {
	Log = logging.NewLogger("")
}
