package mymodule

import "io"

type ExecComand func(args string, outWriter io.Writer, inReader io.Reader) (string, error)

type Command struct {
	writer io.Writer
	reader io.Reader
	comm   ExecComand
}

func NewCommand(writer io.Writer, reader io.Reader, comm ExecComand) *Command {
	return &Command{
		writer: writer,
		reader: reader,
		comm:   comm,
	}
}

func (c *Command) Exec(arg string) error {
	result, err := c.comm(arg, c.writer, c.reader)
	if err != nil {
		return err
	}
	_, err = c.writer.Write([]byte(result))
	if err != nil {
		return err
	}
	return nil
}
