package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Deploy the agent to a Kubernetes cluster",
		Run:   deployAgent,
	}

	var rootCmd = &cobra.Command{Use: "observability-cli"}
	rootCmd.AddCommand(deployCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func deployAgent(cmd *cobra.Command, args []string) {
	// Detect available kubeconfig files
	kubeconfigs, err := detectKubeconfigs()
	if err != nil {
		fmt.Println("Error detecting kubeconfig files:", err)
		return
	}

	if len(kubeconfigs) == 0 {
		fmt.Println("No kubeconfig files found.")
		return
	}

	// Allow the user to choose a kubeconfig file
	selectedKubeconfig := chooseKubeconfig(kubeconfigs)
	if selectedKubeconfig == "" {
		fmt.Println("No kubeconfig selected.")
		return
	}

	// Allow the user to choose a context
	selectedContext := chooseContext(selectedKubeconfig)
	if selectedContext == "" {
		fmt.Println("No context selected.")
		return
	}

	// Deploy the agent to the chosen cluster using the selected kubeconfig and context
	deploy(selectedKubeconfig, selectedContext)
}

func detectKubeconfigs() ([]string, error) {
	
	// Logic to detect kubeconfig files
	// Return a list of file paths
}

func chooseKubeconfig(kubeconfigs []string) string {
	fmt.Print("hello")
	// Logic to allow the user to choose a kubeconfig file
	// Return the selected file path
}

func chooseContext(kubeconfig string) string {
	fmt.Print("hello")
	// Logic to list available contexts in the chosen kubeconfig
	// and allow the user to choose a context
	// Return the selected context name
}

func deploy(kubeconfig, context string) {
	// Logic to deploy the agent to the chosen cluster
	cmd := exec.Command("kubectl", "apply", "-f", "k8s-agent-deployment.yml", "--kubeconfig", kubeconfig, "--context", context)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error deploying agent: %v\n", err)
		return
	}
	fmt.Printf("Agent deployed successfully:\n%s\n", output)
}
