// // package cmd

// // import (
// // 	"fmt"
// // 	"github.com/spf13/cobra"
// // )

// // // runCmd represents the run command
// // var runCmd = &cobra.Command{
// // 	Use:   "run",
// // 	Short: "Run the AI orchestrator with a prompt",
// // 	Long:  `This command sends a prompt to the LLM and orchestrates container tasks.`,
// // 	Run: func(cmd *cobra.Command, args []string) {
// // 		prompt, _ := cmd.Flags().GetString("prompt")
// // 		if prompt == "" {
// // 			fmt.Println("Please provide a prompt using --prompt")
// // 			return
// // 		}
// // 		fmt.Println("Received prompt:", prompt)

// // 		// Add your orchestrator logic here!
// // 		// Call LLM → get tasks → run docker containers → print result
// // 	},
// // }

// // func init() {
// // 	rootCmd.AddCommand(runCmd)
// // 	runCmd.Flags().StringP("prompt", "p", "", "Prompt to process")
// // }
// package cmd

// import (
// 	"fmt"
// 	"os"

// 	"github.com/spf13/cobra"
// 	"ai-orchestrator/internal/docker"
// 	"ai-orchestrator/internal/llm"
// )

// var runCmd = &cobra.Command{
// 	Use:   "run",
// 	Short: "Run the AI orchestrator",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		prompt, _ := cmd.Flags().GetString("prompt")
// 		apiKey := os.Getenv("GROQ_API_KEY")
// 		if apiKey == "" {
// 			fmt.Println("Set GROQ_API_KEY environment variable")
// 			return
// 		}

// 		fmt.Println("Prompt:", prompt)
// 		tasks, err := llm.CallLLM(prompt, apiKey)
// 		if err != nil {
// 			fmt.Println("Error calling LLM:", err)
// 			return
// 		}

// 		fmt.Println("Tasks received:", tasks)
// 		for _, task := range tasks {
// 			if err := docker.RunDockerTask(task); err != nil {
// 				fmt.Println("Error running task:", task, err)
// 				return
// 			}
// 		}

// 		fmt.Println("✅ All tasks complete.")
// 	},
// }

// package cmd

// import (
// 	"fmt"
// 	"os"

// 	"github.com/spf13/cobra"
// 	"ai-orchestrator/internal/docker"
// 	"ai-orchestrator/internal/llm"
// )

// var prompt string

// var runCmd = &cobra.Command{
// 	Use:   "run",
// 	Short: "Run the AI orchestrator",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if prompt == "" {
// 			fmt.Println("❌ Please provide a prompt using --prompt")
// 			return
// 		}

// 		apiKey := os.Getenv("GROQ_API_KEY")
// 		if apiKey == "" {
// 			fmt.Println("❌ Set your GROQ_API_KEY environment variable")
// 			return
// 		}

// 		fmt.Println("🧠 Sending prompt to LLM:", prompt)
// 		tasks, err := llm.CallLLM(prompt, apiKey)
// 		if err != nil {
// 			fmt.Println("LLM error:", err)
// 			return
// 		}

// 		fmt.Println("📦 Tasks returned by LLM:", tasks)

// 		for _, task := range tasks {
// 			if err := docker.RunDockerTask(task); err != nil {
// 				fmt.Println("Error running task:", task, err)
// 				return
// 			}
// 		}

// 		fmt.Println("✅ All tasks completed.")
// 	},
// }

// func init() {
// 	// Register the flag and bind it to the prompt variable
// 	runCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Prompt to process with the AI Orchestrator")

// 	// Attach runCmd to the root
// 	rootCmd.AddCommand(runCmd)
// }
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ai-orchestrator/internal/docker"
	"ai-orchestrator/internal/llm"
)

var prompt string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the AI orchestrator",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Check prompt flag
		if prompt == "" {
			fmt.Println("❌ Please provide a prompt using --prompt or -p")
			return
		}

		// 2. Check API key
		apiKey := os.Getenv("GROQ_API_KEY")
		if apiKey == "" {
			fmt.Println("❌ Set your GROQ_API_KEY environment variable")
			return
		}

		// 3. Write prompt to data/input.txt
		err := os.MkdirAll("data", os.ModePerm)
		if err != nil {
			fmt.Println("❌ Failed to create data folder:", err)
			return
		}
		err = os.WriteFile("data/input.txt", []byte(prompt), 0644)
		if err != nil {
			fmt.Println("❌ Failed to write prompt to input.txt:", err)
			return
		}

		// 4. Call LLM
		fmt.Println("🧠 Sending prompt to LLM:", prompt)
		tasks, err := llm.CallLLM(prompt, apiKey)
		if err != nil {
			fmt.Println("LLM error:", err)
			return
		}

		// 5. Show task plan
		fmt.Println("📦 Tasks returned by LLM:", tasks)

		// 6. Run each task
		for _, task := range tasks {
			fmt.Println("⚙️ Running task in container:", task)
			if err := docker.RunDockerTask(task); err != nil {
				fmt.Println("❌ Error running task:", task, err)
				return
			}
		}

		// 7. Done
		fmt.Println("✅ All tasks completed.")
	},
}

func init() {
	// Register the flag and bind it to the prompt variable
	runCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Prompt to process with the AI Orchestrator")

	// Attach runCmd to the root command
	rootCmd.AddCommand(runCmd)
}
