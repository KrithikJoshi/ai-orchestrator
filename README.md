# 🧠 AI Orchestrator (DevOps Backend Assignment)

This is a CLI-based AI orchestrator that uses a Large Language Model (LLM) to decide which data processing tasks to run inside Docker containers. The orchestrator receives a natural language prompt, consults an LLM, and runs containerized tasks like `clean_data` and `sentiment_analysis` based on the LLM's plan.

---

## 🚀 Features

- 🧠 LLM integration via Groq API
- 🐳 Task-based container orchestration with Docker
- 📦 Modular containerized microservices
- 🖥️ CLI interface with Cobra (Go)
- 📂 Data passed between containers via shared volume
- ✅ Outputs final result to both file and terminal

---

## 📁 Project Structure

ai-orchestrator/ ├── main.go ├── cmd/ │ └── run.go ├── internal/ │ ├── llm/ # Calls Groq API │ └── docker/ # Runs Docker containers ├── tasks/ │ ├── clean_data/ │ │ ├── clean.py │ │ ├── Dockerfile │ │ └── requirements.txt │ └── sentiment_analysis/ │ ├── sentiment.py │ ├── Dockerfile │ └── requirements.txt ├── data/ │ ├── input.txt # Prompt is saved here │ ├── output.txt # Cleaned text │ └── sentiment.txt # Final sentiment result ├── go.mod └── README.md