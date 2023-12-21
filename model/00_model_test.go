package model

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

func TestCodeGen(t *testing.T) {
	gen := []gdb.GeneratorModelForStructParam{
		competitionModel(),
		matchModel(),
	}

	gdb.GeneratorModelFromStruct(gen...)
}

func competitionModel() gdb.GeneratorModelForStructParam {
	return gdb.GeneratorModelForStructParam{
		Src: &Competition{},
		SpecifiationTable: gdb.SpecifiationTable{
			TableName: "competition",
		},
		Tag:      "db",
		FileName: "competition",
	}
}

func matchModel() gdb.GeneratorModelForStructParam {
	return gdb.GeneratorModelForStructParam{
		Src: &Match{},
		SpecifiationTable: gdb.SpecifiationTable{
			TableName: "match",
		},
		Tag:      "db",
		FileName: "match",
	}
}
