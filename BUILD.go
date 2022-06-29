package minicov

import (
	"dbt-rules/RULES/cc"
)

var Lib = cc.Library{
	Out: out("minicov.a"),
	Srcs: ins(
		"minicov/c/InstrProfiling.c",
		"minicov/c/InstrProfilingBuffer.c",
		"minicov/c/InstrProfilingInternal.c",
		"minicov/c/InstrProfilingMerge.c",
		"minicov/c/InstrProfilingPlatformLinux.c",
		"minicov/c/InstrProfilingVersionVar.c",
		"minicov/c/InstrProfilingWriter.c",
	),
	AlwaysLink: true,
}
