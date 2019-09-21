package script

func build(filename, filepath string) error {
	cmd := exec.Command("go", "build", "-o", filename, filepath)
	return cmd.Start()
}
