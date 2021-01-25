package pkg

import (
	"bufio"
	"context"
	"io"
	"os"
	"syscall"
)

type Pipe interface {
	io.Reader
	io.Writer
	io.Closer
	Listen(del byte) chan []byte
}

type pipe struct {
	*os.File
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewPipe(name string) (Pipe, error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := syscall.Mkfifo(name, 0666)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeNamedPipe)
	if err != nil {
		return nil, err
	}
	p := &pipe{File: file}
	p.ctx, p.cancelFunc = context.WithCancel(context.Background())
	return p, err
}

func (p *pipe) Close() error {
	p.cancelFunc()
	return p.File.Close()
}

func (p *pipe) Listen(del byte) chan []byte {
	out := make(chan []byte)
	go func() {
		br := bufio.NewReader(p.File)
		var err error
		var bf []byte
		for err == nil {
			select {
			case <-p.ctx.Done():
				close(out)
				return
			default:
			}

			bf, err = br.ReadBytes(del)
			if err != nil && err != io.EOF {
				return
			}
			out <- bf
		}
	}()
	return out
}
