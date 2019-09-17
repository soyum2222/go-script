# go-script
Go dynamic pseudo parser

## What done
A Go pseudo script executor ,can be embedded in normal code.

## Example
##### main.go
```cassandraql
func main() {

	for {
		script := scripte.NewScript("example/script/script")

		script.Arge("go-script")
		decoder, err := script.Run()
		if err != nil {
			panic(err)
		}

		var r int
		decoder.Decode(&r)

		fmt.Println("go-script return :", r)

		time.Sleep(5 * time.Second)

	}

}
```

##### script
```cassandraql
func main_script(in interface{})interface{}{

    fmt.Println("in put arge is ", in)

    fmt.Println("for test")

    fmt.Println("the script return value is ",123)

    return 123

}
```

You can modify `script` file in runtime , get new modified results

