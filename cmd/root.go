package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	port    int
	path    string
	rootCmd = &cobra.Command{
		Use:   "webshare",
		Short: "simple share file",
		Run: func(cmd *cobra.Command, args []string) {
			// Validate directory path
			info, err := os.Stat(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if !info.IsDir() {
				fmt.Fprintf(os.Stderr, "Error: %s is not a directory\n", path)
				os.Exit(1)
			}

			e := echo.New()
			e.HideBanner = true

			e.Static("/", path)
			e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
				Root:   path,
				Browse: true,
			}))
			e.GET("/status", func(c echo.Context) error {
				return c.String(http.StatusOK, "status: running")
			})
			address := fmt.Sprintf(":%d", port)
			fmt.Printf("enter this address in other host in local network: http://%s:%d \n", getIP(), port)
			e.Logger.Fatal(e.Start(address))
		},
	}
)

func init() {
	banner := figure.NewFigure("WebShare", "", true).String()
	fmt.Println(banner)
	fmt.Printf("version: %s \n", getGitVersion())
	fmt.Println("available on https://github.com/sinameshkini/webshare")
	rootCmd.Flags().IntVarP(&port, "port", "p", 4242, "Port number")
	rootCmd.Flags().StringVarP(&path, "dir", "d", ".", "Directory path")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func getIP() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	return
}

func getGitVersion() string {
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
