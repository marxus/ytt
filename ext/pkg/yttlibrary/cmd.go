package yttlibrary

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"carvel.dev/ytt/pkg/template/core"
	"github.com/k14s/starlark-go/starlark"
	"github.com/k14s/starlark-go/starlarkstruct"
)

var (
	CmdAPI = starlark.StringDict{
		"cmd": &starlarkstruct.Module{
			Name: "cmd",
			Members: starlark.StringDict{
				"run": starlark.NewBuiltin("cmd.run", core.ErrWrapper(cmdModule{}.Run)),
			},
		},
	}

	cmdRuns = map[string]starlark.String{}
)

type cmdModule struct{}

func (cmdModule) Run(thread *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	if args.Len() == 0 {
		return starlark.None, fmt.Errorf("expected at least one argument")
	}

	var vals []string
	for i := 0; i < args.Len(); i++ {
		val, err := core.NewStarlarkValue(args.Index(i)).AsString()
		if err != nil {
			return starlark.None, err
		}
		vals = append(vals, val)
	}

	key := strings.Join(vals, "")
	if val, ok := cmdRuns[key]; ok {
		return val, nil
	}

	cmd := exec.Command(vals[0], vals[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if cmd.Run() != nil {
		return nil, fmt.Errorf("\n%s", stderr.String())
	}

	val := starlark.String(stdout.String())
	cmdRuns[key] = val
	return val, nil
}
