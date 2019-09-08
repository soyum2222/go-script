package template

var Import_template = `
package main
import (
	"github.com/natefinch/npipe"
	"encoding/gob"
	"bytes"
	"flag"
	"time"
	"os"
	"fmt"
	"io/ioutil"
)`

var Script_template = `

type In struct {
	Body interface{}
}

func main() {

	fmt.Println("script start")
	flag.Parse()
	args:= flag.Args()
	if args==nil||len(args)<1{
		panic("bad run")
	}
	conn,err:=npipe.DialTimeout(args[0],5*time.Second)
	if err != nil {
		panic(err)
	}

	b,err:=ioutil.ReadAll(os.Stdin)
	if err!=nil{
		panic(err)
	}

	in:=In{}
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err = decoder.Decode(&in)

	out := main_script(in.Body)

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err = encoder.Encode(out)
	ret_byte := result.Bytes()

	conn.Write(ret_byte)
}
`
