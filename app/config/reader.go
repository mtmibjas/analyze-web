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
func parseConfig(file string, unpacker any) {
	content := read(file)
	err := yaml.Unmarshal(content, unpacker)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
}

func parseLogger(file string, unpacker any) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = yaml.Unmarshal(data, unpacker)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return
	}
}

func getDirPath(dir string) string {

	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		panic("somthing went wrong while get root path")
	}
	moduleRoot := strings.TrimSpace(string(output))

	c := dir[len(dir)-1]
	if os.IsPathSeparator(c) {
		return moduleRoot + dir
	}
	return moduleRoot + dir + string(os.PathSeparator)
}
