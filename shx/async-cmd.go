package shx

import (
	"fmt"
	"os/exec"

	"github.com/mattn/go-shellwords"
	"github.com/samber/lo"

	"go.szostok.io/magex/printer"
)

type AsyncCommand struct {
	cmd *exec.Cmd
}

func MustAsyncCmdf(format string, a ...interface{}) *AsyncCommand {
	return lo.Must(AsyncCmdf(format, a...))
}

func AsyncCmdf(format string, a ...interface{}) (*AsyncCommand, error) {
	rawCMD := fmt.Sprintf(format, a...)
	envs, args, err := shellwords.ParseWithEnvs(rawCMD)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = append(cmd.Env, envs...)
	return &AsyncCommand{
		cmd: cmd,
	}, nil
}

func (c AsyncCommand) Start() error {
	printer.Cmd(c.cmd.String())
	return c.cmd.Start()
}

func (c AsyncCommand) Stop() error {
	return c.cmd.Process.Kill()
}

func (c AsyncCommand) MustStart() {
	lo.Must0(c.Start())
}

func (c AsyncCommand) MustStop() {
	lo.Must0(c.Stop())
}
