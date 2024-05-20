package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

func read(file string) []byte {

	content, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return content
}
func parseConfig(file string, unpacker interface{}) {

	content := read(file)
	err := yaml.Unmarshal(content, unpacker)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
}

func getConfigDir(dir string) string {
	c := dir[len(dir)-1]
	if os.IsPathSeparator(c) {
		return dir
	}
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		panic("somthing went wrong while get root path")
	}
	moduleRoot := strings.TrimSpace(string(output))

	return moduleRoot + dir + string(os.PathSeparator)
}
