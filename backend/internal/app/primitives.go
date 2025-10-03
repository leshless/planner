package app

import (
	"planner/backend/internal/util/interrupter"
	"planner/backend/internal/util/stdio"

	xos "github.com/hack-pad/hackpadfs/os"

	xfs "github.com/hack-pad/hackpadfs"

	"github.com/benbjohnson/clock"
)

type Primitives struct {
	fs          xfs.FS
	clock       clock.Clock
	interrupter interrupter.Interrupter
	stdIO       stdio.StdIO
}

func NewPrimitives(
	fs xfs.FS,
	clock clock.Clock,
	interrupter interrupter.Interrupter,
	stdIO stdio.StdIO,
) Primitives {
	return Primitives{
		fs:          fs,
		clock:       clock,
		interrupter: interrupter,
		stdIO:       stdIO,
	}
}

func NewDefaultPrimitives() Primitives {
	return Primitives{
		fs:          xos.NewFS(),
		clock:       clock.New(),
		interrupter: interrupter.NewInterrupter(),
		stdIO:       stdio.NewStdIO(),
	}
}
