package cmd

import (
	"fmt"
	"os"

	"ai-orchestrator/internal/docker"
	"ai-orchestrator/internal/llm"

	"github.com/spf13/cobra"
)

var prompt string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the AI orchestrator",
	Run: func(cmd *cobra.Command, args []string) {
		
		if prompt == "" {
			fmt.Println(" Please provide a prompt using --prompt or -p")
			return
		}

		
		apiKey := os.Getenv("GROQ_API_KEY")
		if apiKey == "" {
			fmt.Println(" Set your GROQ_API_KEY environment variable")
			return
		}

		
		err := os.MkdirAll("data", os.ModePerm)
		if err != nil {
			fmt.Println(" Failed to create data folder:", err)
			return
		}
		err = os.WriteFile("data/input.txt", []byte(prompt), 0644)
		if err != nil {
			fmt.Println(" Failed to write prompt to input.txt:", err)
			return
		}


		fmt.Println(" Sending prompt to LLM:", prompt)
		tasks, err := llm.CallLLM(prompt, apiKey)
		if err != nil {
			fmt.Println("LLM error:", err)
			return
		}

		
		fmt.Println(" Tasks returned by LLM:", tasks)

		
		for _, task := range tasks {
			fmt.Println(" Running task in container:", task)
			if err := docker.RunDockerTask(task); err != nil {
				fmt.Println(" Error running task:", task, err)
				return
			}
		}

		
		fmt.Println(" All tasks completed.")
	},
}

func init() {
	
	runCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Prompt to process with the AI Orchestrator")

	
	rootCmd.AddCommand(runCmd)
}
