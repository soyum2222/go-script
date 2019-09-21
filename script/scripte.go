package script

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"github.com/soyum2222/go-script/template"

	"io"
	"os"
	"path"
)

type Script struct {
	script_path string
	arge        interface{} //only have one
}

func (script *Script) Run() (*gob.Decoder, error) {
	return runScripte(script.script_path, script.arge)
}

func (script *Script) Arge(arge interface{}) {
	script.arge = arge
}

func NewScript(script_path string) *Script {
	return &Script{
		script_path: script_path,
	}
}

func md5sum(src []byte) string {

	m := md5.New()
	m.Write(src)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func compile(code []byte) string {

	tmpdir := os.TempDir()
	tmpdir += "/go-script"

	file_hash_code := md5sum(code)

	code_filepath := tmpdir + "/code/" + file_hash_code + ".go"
	err := os.MkdirAll(path.Dir(code_filepath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(code_filepath)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(code)
	if err != nil {
		panic(err)
	}

	build_filepath := tmpdir + "/exec/" + file_hash_code

	err = build(code_filepath, build_filepath)

	if err != nil {
		panic(fmt.Sprintf("build fail %s", err))
	}

	return build_filepath

}

func coding(code []byte) []byte {

	codebuf := bytes.Buffer{}

	codebuf.WriteString(template.Import_template)
	codebuf.WriteByte('\n')
	codebuf.Write(code)
	codebuf.WriteByte('\n')
	codebuf.WriteString(template.Script_template)

	return codebuf.Bytes()

}

func run(path string, arg interface{}) ([]byte, error) {

	runner := NewRunner(path)

	err := runner.Arg(arg)
	if err != nil {
		return nil, err
	}
	err = runner.Run()
	if err != nil {
		return nil, err
	}
	ret, err := runner.Ret()
	return ret, err

}

func runScripte(path string, arg interface{}) (*gob.Decoder, error) {

	//code, err := ioutil.ReadFile(path)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)

	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}
	code := buf.Bytes()

	code = coding(code)

	exec_path := compile(code)

	ret, err := run(exec_path, arg)
	if err != nil {
		return nil, err
	}

	return gob.NewDecoder(bytes.NewReader(ret)), nil

}
