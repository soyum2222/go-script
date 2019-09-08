package scripte

import (
	"bytes"
	"encoding/gob"
	"github.com/natefinch/npipe"
	"io/ioutil"
	"os"
	"os/exec"
)

type In struct {
	Body interface{}
}

const pipe_url = `\\.\pipe\`

type WinRun struct {
	exec_path string
	pipe_url  string
	args      []string
	stdin     *bytes.Buffer
}

func (w *WinRun) Run() error {

	cmd := exec.Command(w.exec_path, w.args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = w.stdin

	return cmd.Start()
}

func (w *WinRun) Arg(i interface{}) error {
	in := In{
		Body: i,
	}
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(in)
	if err != nil {
		return err
	}

	arg_byte := result.Bytes()
	if w.stdin == nil {
		w.stdin = bytes.NewBuffer(arg_byte)
	} else {
		w.stdin.Write(arg_byte)
	}

	return err
}

func (w *WinRun) Ret() ([]byte, error) {

	listen, err := npipe.Listen(w.pipe_url)
	if err != nil {
		return nil, err
	}

	defer listen.Close()

	conn, err := listen.Accept()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	ret, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}
	//
	//var v interface{}
	//decoder := gob.NewDecoder(bytes.NewReader(ret))
	//err = decoder.Decode(&v)
	return ret, nil

}

func NewRunner(path string) Runner {

	return &WinRun{
		exec_path: path,
		pipe_url:  pipe_url + "test",
		stdin:     bytes.NewBuffer(nil),
		args:      append([]string{}, pipe_url+"test"),
	}
}
