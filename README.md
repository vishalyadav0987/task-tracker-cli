
# 🧠 Task Tracker CLI (Golang)

🔗 **Project URL:**  
https://github.com/vishalyadav0987/task-tracker-cli

---

A simple yet powerful **Command Line Interface (CLI)** application built in **Go (Golang)** to manage your daily tasks efficiently.

---

## 🚀 Features

- ✅ Add new tasks
- 📋 List all tasks (table format)
- ✏️ Update task description
- ❌ Delete tasks
- 🔄 Mark task as:
  - Todo
  - In Progress
  - Done
- 🔍 Filter tasks by status
- 🎨 Colored terminal output (better UX)
- 🧾 Persistent storage using JSON file

---

## 📁 Project Structure

```

task-tracker/
│
├── cmd/
│   └── task-tracker/
│       └── main.go
│
├── internal/
│   ├── domain/
│   ├── application/
│   ├── infrastructure/
│   └── interfaces/
│
├── data/
│   └── tasks.json
│
└── README.md

````

---

## ⚙️ Installation

### 1. Clone repo

```bash
git clone https://github.com/vishalyadav0987/task-tracker-cli.git
cd task-tracker-cli
````

---

### 2. Build binary

```bash
go build -o task-cli cmd/task-tracker/main.go
```

---

### 3. Make it global (Mac/Linux)

```bash
sudo mv task-cli /usr/local/bin/
```

---

## 🧪 Usage

### ➕ Add Task

```bash
task-cli add "Buy groceries"
```

---

### 📋 List Tasks

```bash
task-cli list
```

---

### ✏️ Update Task

```bash
task-cli update <task_id> "New description"
```

---

### ❌ Delete Task

```bash
task-cli delete <task_id>
```

---

### 🔄 Mark Task Status

```bash
task-cli mark-done <task_id>
task-cli mark-in-progress <task_id>
```

---

### 🔍 Filter by Status

```bash
task-cli status todo
task-cli status in-progress
task-cli status done
```

---

### 🔎 Get Task by ID

```bash
task-cli task <task_id>
```

---

## 🖥️ Example Output

```
+--------------------------------------+----------------------+----------------+
| ID                                   | DESCRIPTION          | STATUS         |
+--------------------------------------+----------------------+----------------+
| eb79310d-0c36-4fbf...                | Gym jaana            | 📌 Todo        |
| f1eab601-eacc-4b36...                | Learn Go             | ⏳ In Progress |
| dd4396c2-7995-4592...                | Build project        | ✔ Done        |
+--------------------------------------+----------------------+----------------+
```

---

## 🧠 Tech Stack

* Go (Golang)
* Clean Architecture
* JSON File Storage
* CLI Interface

---

## 🔥 Future Improvements

* UUID short IDs
* Interactive CLI (arrow navigation)
* Cobra CLI integration
* REST API version
* Database support (SQLite/Postgres)

---

## 👨‍💻 Author

**Vishal Yadav**

---

## ⭐ If you like this project

Give it a ⭐ on GitHub!

