package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	open "github.com/fermyon/otel-plugin/cmd/open"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otel",
	Short: "A plugin that makes using Spin with OTel easy.",
	Long:  "A plugin that makes using Spin with OTel easy by automatically standing up dependencies, sourcing environment variables, and linking to dashboards.",
}

var otelConfigDirName = "otel-config"
var otelConfigPath string

// setOtelConfigPath allows for someone to run the otel plugin directly from source or via the Spin plugin installation
func setOtelConfigPath() error {
	executablePath, err := os.Executable()
	if err != nil {
		return err
	}

	otelConfigPath = path.Join(path.Dir(executablePath), otelConfigDirName)

	if _, err := os.Stat(otelConfigPath); os.IsNotExist(err) {
		return fmt.Errorf("The directory in which the plugin binary is executed is missing necessary files, so please make sure the plugin was installed using \"spin plugins install otel\"")
	}

	return nil
}

// checkDocker checks whether Docker is installed and the Docker daemon is running
func checkDocker() error {
	cmd := exec.Command("docker", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Docker appears not to be installed, so please visit their install page and try again once installed: https://www.docker.com/products/docker-desktop")
	}

	cmd = exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("The Docker daemon appears not to be running. The command to start Docker depends on your operating system. For instructions, check the correct page under https://docs.docker.com/engine/install")
	}

	return nil
}

func Execute() {
	if err := setOtelConfigPath(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("Error finding the \"otel-config\" directory: %w", err))
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(cleanUpCmd)
	rootCmd.AddCommand(setUpCmd)
	rootCmd.AddCommand(openCmd)
	openCmd.AddCommand(open.GrafanaCmd)
	openCmd.AddCommand(open.JaegerCmd)
	openCmd.AddCommand(open.PrometheusCmd)
	openCmd.AddCommand(open.AspireCmd)
}
