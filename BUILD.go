package minicov

import (
	"utest/ulib"
	"utest/RULES/testlib"
)

var Lib = testlib.TestLib{
	Out:           out("minicov.a"),
	Srcs:          ins(
		"minicov/c/InstrProfiling.c",
		"minicov/c/InstrProfilingBuffer.c",
		"minicov/c/InstrProfilingInternal.c",
		"minicov/c/InstrProfilingMerge.c",
		"minicov/c/InstrProfilingPlatformLinux.c",
		"minicov/c/InstrProfilingVersionVar.c",
		"minicov/c/InstrProfilingWriter.c",
	),
	Deps:          []interface{}{ulib.ULib},
	AlwaysLink:    true,
}
