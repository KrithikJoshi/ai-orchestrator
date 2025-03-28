
package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunDockerTask(task string) error {
	
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current dir: %v", err)
	}
	dataPath := filepath.Join(cwd, "data")

	fmt.Println(" Mounting host data folder:", dataPath)

	cmd := exec.Command("docker", "run", "--rm", "-v", dataPath+":/data", task)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
