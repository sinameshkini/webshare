package utils

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func GetIP() (ip string) {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	for _, address := range adders {
		if aspnet, ok := address.(*net.IPNet); ok && !aspnet.IP.IsLoopback() && aspnet.IP.To4() != nil {
			return aspnet.IP.String()
		}
	}

	return
}

func GetGitVersion() string {
	// Get the most recent tag and the current branch
	tagCmd := exec.Command("git", "describe", "--tags")
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")

	tagOutput, err := tagCmd.Output()
	if err != nil {
		fmt.Println("Error getting Git tag:", err)
		return ""
	}

	branchOutput, err := branchCmd.Output()
	if err != nil {
		fmt.Println("Error getting Git branch:", err)
		return ""
	}

	tag := strings.TrimSpace(string(tagOutput))
	branch := strings.TrimSpace(string(branchOutput))

	// Construct the version string
	version := fmt.Sprintf("%s-%s", tag, branch)

	return version
}
