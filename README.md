#  AI Orchestrator (DevOps Backend Assignment)

This is a CLI-based AI orchestrator that uses a Large Language Model (LLM) to decide which data processing tasks to run inside Docker containers. The orchestrator receives a natural language prompt, consults an LLM, and runs containerized tasks like `clean_data` and `sentiment_analysis` based on the LLM's plan.

---

## 🚀 Features

-  LLM integration via Groq API
-  Task-based container orchestration with Docker
-  Modular containerized microservices
-  CLI interface with Cobra (Go)
-  Data passed between containers via shared volume
-  Outputs final result to both file and terminal

---

##  Project Structure

ai-orchestrator/
├── cmd/
│   ├── run.go                    # Cobra CLI logic
│   └── root.go                   # Root command setup
├── internal/
│   ├── docker/
│   │   └── runner.go             # Runs Docker containers from Go
│   └── llm/
│       └── llm.go                # Calls Groq API and extracts task list
├── tasks/
│   ├── clean_data/
│   │   ├── clean.py              # Cleans the input text
│   │   ├── Dockerfile
│   │   └── requirements.txt
│   └── sentiment_analysis/
│       ├── sentiment.py          # Analyzes sentiment of cleaned text
│       ├── Dockerfile
│       └── requirements.txt
├── data/
│   ├── input.txt                 # User prompt gets saved here
│   ├── output.txt                # Output from clean_data
│   └── sentiment.txt             # Output from sentiment_analysis
├── main.go                       # CLI entrypoint
├── go.mod                        # Go module file
├── go.sum                        # Go module dependencies
├── architecture.png              # System design diagram
└── README.md                     # You're reading it 🎉


##  Setup Instructions

###  Install Requirements

- [Go](https://go.dev/doc/install)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Groq API Key](https://console.groq.com)

---

###  Step 1: Clone & Build Containers

# Navigate to each task and build
cd tasks/clean_data
docker build -t clean_data .

cd ../sentiment_analysis
docker build -t sentiment_analysis .

###  Step 2: Set Your API Key

# For Bash/Git Bash
export GROQ_API_KEY=your_groq_api_key_here

# For Windows CMD
set GROQ_API_KEY=your_groq_api_key_here

### Step 3: Run the CLI

go run main.go run --prompt "I love the product but hate the shipping."

### Example output

Sending prompt to LLM: I love the product but hate the shipping.
Tasks returned by LLM: [clean_data sentiment_analysis]
Running task in container: clean_data
Running task in container: sentiment_analysis
Final Sentiment Output:
Sentiment: Negative (score: -0.15)

## Architecture

1. User provides CLI prompt

2. Orchestrator calls Groq LLM

3. LLM returns ordered task list

4. Docker containers are spun up in sequence

5. Data passed between containers via volume

6. Final output returned to user