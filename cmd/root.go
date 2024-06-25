package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/webshare/internal/handler"
	"github.com/sinameshkini/webshare/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var (
	port     int
	path     string
	username string
	password string
	rootCmd  = &cobra.Command{
		Use:   "webshare",
		Short: "WebShare is a simple and efficient tool for sharing files over HTTP/HTTPS with basic authentication.",
		Long: `
	WebShare is a user-friendly application designed for quick and secure file sharing over HTTP/HTTPS. 
	It allows you to easily set up a web service that provides access to specified directories for both 
	reading and writing data. With cross-platform compatibility and support for basic authentication, WebShare 
	ensures that your files can be accessed securely from any device on your local network or over the internet.
	Whether you need to share files for collaboration or distribute data efficiently,
	WebShare offers a straightforward solution with minimal setup.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Validate directory path
			info, err := os.Stat(path)
			if err != nil {
				logrus.Fatalln(err.Error())
			}
			if !info.IsDir() {
				logrus.Fatalf("Error: %s is not a directory\n", path)
			}

			e := echo.New()
			e.HideBanner = true

			handler.Init(e, path, port, username, password)

			address := fmt.Sprintf(":%d", port)
			e.Logger.Fatal(e.Start(address))
		},
	}
)

func init() {
	banner := figure.NewFigure("WebShare", "", true).String()
	fmt.Println(banner)
	fmt.Printf("version: %s \n", utils.GetGitVersion())
	fmt.Println("Available on https://github.com/sinameshkini/webshare")
	rootCmd.Flags().IntVarP(&port, "port", "p", 4242, "Specifies the port number on which the web service will run.")
	rootCmd.Flags().StringVarP(&path, "dir", "d", ".", "Specifies the directory path to be shared.")
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Set the username for basic authentication.")
	rootCmd.Flags().StringVarP(&password, "password", "P", "", "Set the password for basic authentication.")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
