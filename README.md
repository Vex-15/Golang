# Golang Masterclass :

This is where I actually learn Go — not by taking notes I'll never open again, but by building things, breaking them, and fixing them until they make sense.

I build projects as I learn, extending many of them beyond the original idea to explore how Go works in practice. Whenever a concept clicks, I try to push it a little further, turn it into something practical, and write about the projects that teach me the most.
This repository is the running record of that—a public learning journal and practical reference for anyone learning Go alongside me.

If you're learning Go too, I hope something in here saves you an hour of confusion I had to sit through myself.

---

## How This Is Organized

Each top-level folder is a concept area. Inside, you'll usually find a small standalone Go program exploring one idea, and sometimes a full project where a few ideas come together.

```text
Golang-Masterclass
│
├── goroutines/                     # Concurrency: the core of this repo so far
│   ├── main.go                     # Ping-pong with context cancellation
│   ├── Mutex/                      # Thread-safe bank account with sync.Mutex
│   └── Project - FileDownloader/   # Concurrent file downloader (featured project)
│
├── Practise/                       # Early warm-up snippets, kept as a record of day one
│
└── README.md
```

The repository grows naturally as I explore new concepts and build more projects.

---

## What's Inside So Far

### Concurrency

Concepts explored:

- Goroutines
- WaitGroups
- Channels
- Mutexes
- Context Cancellation
- Buffered Channels

Projects:

- Ping-Pong with Context Cancellation
- Mutex-Protected Bank Account
- Concurrent File Downloader


### Early Practice (`Practise/`)

Some of my very first Go snippets from before I had any real footing in the language. I'm keeping these instead of deleting them because they're an honest starting point — it's easier to see how far something has come when the beginning is still there.

More folders will be added as I continue exploring new areas of Go, from networking and data handling to backend development.

---

## Featured Project: Concurrent File Downloader

The project I'm proudest of so far. It downloads multiple files at once instead of one-by-one, and along the way covers:

- Goroutines and WaitGroups for running downloads in parallel
- A buffered-channel semaphore to cap how many downloads run concurrently
- Channels for collecting results (size, duration, errors) from each worker
- A basic sequential-vs-concurrent comparison to see the actual speedup

I also wrote a detailed walkthrough covering how it works, the design decisions behind it, and a review of the mistakes I made while building it.

📖 **[Building a Concurrent File Downloader in Go: A Practical Guide to Goroutines](https://medium.com/@vexstack/building-a-concurrent-file-downloader-in-go-a-practical-guide-to-goroutines-4071dd01f946)**

## Running a Project

Clone the repo:

```bash
git clone https://github.com/Vex-15/Golang-Masterclass.git
cd Golang-Masterclass
```

Each project is self-contained.

Navigate to the project you want to run:

```bash
cd <project-directory>
go run .
```

For example:

```bash
cd "goroutines/Project - FileDownloader"
go run .
```

---

## Why I'm Doing This in the Open

I could keep this all local and never show anyone the messy parts. But building in public keeps me honest about what I actually understand versus what I've just copied from a tutorial, and it means the mistakes are as visible as the wins — which is usually the more useful part for anyone else learning the same thing.

If you spot something I got wrong, or a cleaner way to do what I did, I'd genuinely like to hear it. Open an issue or a PR.

---

## Connect

- GitHub: [github.com/Vex-15](https://github.com/Vex-15)
- Medium: [medium.com/@vexstack](https://medium.com/@vexstack)
- X: [x.com/vexstack](https://x.com/vexstack)

---

If any of this helped you get unstuck on something, a ⭐ on the repo is always appreciated.
