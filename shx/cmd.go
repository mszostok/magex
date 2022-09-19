package shx

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/carolynvs/magex/shx"
	"github.com/mattn/go-shellwords"
	"github.com/samber/lo"

	"go.szostok.io/magex/printer"
)

type Command struct {
	shx.PreparedCommand
}

func Cmdf(format string, a ...interface{}) (*Command, error) {
	rawCmd := fmt.Sprintf(format, a...)
	envs, args, err := shellwords.ParseWithEnvs(rawCmd)
	if err != nil {
		return nil, err
	}

	return &Command{
		PreparedCommand: shx.Command(args[0], args[1:]...).Env(envs...),
	}, nil
}

func CmdCtxf(ctx context.Context, format string, a ...interface{}) (*Command, error) {
	rawCmd := fmt.Sprintf(format, a...)
	envs, args, err := shellwords.ParseWithEnvs(rawCmd)
	if err != nil {
		return nil, err
	}

	c := exec.CommandContext(ctx, args[0], args[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Env = os.Environ()

	return &Command{
		PreparedCommand: shx.PreparedCommand{Cmd: c}.Env(envs...),
	}, nil
}

func MustCmdf(format string, a ...interface{}) *Command {
	return lo.Must1(Cmdf(format, a...))
}

func (c *Command) Run() error {
	printer.Cmd(c.String())
	return c.PreparedCommand.Run()
}

func (c *Command) RunV() error {
	printer.Cmd(c.String())
	return c.PreparedCommand.RunV()
}

func (c *Command) RunE() error {
	printer.Cmd(c.String())
	return c.PreparedCommand.RunE()
}

func (c *Command) RunS() error {
	printer.Cmd(c.String())
	return c.PreparedCommand.RunS()
}

func (c *Command) WithArg(name, val string) *Command {
	c.Args(fmt.Sprintf("%s=%s", name, val))
	return c
}

func (c *Command) In(dir string) *Command {
	c.PreparedCommand = c.PreparedCommand.In(dir)
	return c
}
