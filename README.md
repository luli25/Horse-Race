# Horse Race Simulator 🐎

Concurrent horse race simulator built in Go, with **real-time updates** via WebSockets and a simple frontend using HTML/JS.

This project demonstrates:
- Concurrency in Go using **goroutines**.
- Safe synchronization with **sync.Mutex**.
- Real-time communication using **WebSockets**.
- Modular backend + frontend architecture.
- HTTP routing using `http.ServeMux`.

---

## 🔧 Technologies

- Go 1.22+
- WebSockets using [Gorilla WebSocket](https://github.com/gorilla/websocket)
- HTML and JavaScript for a simple frontend
- Modular backend (`internal/horse`, `internal/race`, `api`)

---

## 🚀 How to run

1. Clone the repository:

```bash
git clone https://github.com/yourusername/horse-race.git
cd horse-race
```

2. Initialize Go modules
   ```bash
   go mod tidy
   ```

3. Run the server
   ```bash
   go run ./cmd/server
   ```

4. Open your browser and visit:
   http://localhost:8080

---

🏁 How it works
Each horse runs in its own goroutine, advancing random positions.
The race is broadcasted in real-time using WebSockets.
The frontend displays the horses with progress bars and announces the winner.
A mutex is used to synchronize horse positions and avoid race conditions.

---

## 📂 Project Structure
```text
horse-race/
├── cmd/
│   └── server/      # Server entry point
├── internal/
│   ├── horse/       # Horse model
│   └── race/        # Race logic
├── api/             # WebSocket handlers
├── web/             # Frontend HTML/JS
├── go.mod           # Dependencies
└── README.md        # This file
