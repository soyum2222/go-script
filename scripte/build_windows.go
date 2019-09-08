package scripte

import (
	"os"
	"os/exec"
)

func build(filename, filepath string) error {
	filepath += ".exe"
	cmd := exec.Command("go", "build", "-o", filepath, filename)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
