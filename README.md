# Horse Race Simulator 🐎

Simulador concurrente de carreras de caballos desarrollado en Go, con **transmisión en tiempo real** mediante WebSockets y frontend simple en HTML/JS.

El proyecto demuestra:
- Concurrencia en Go usando **goroutines**.
- Sincronización segura con **sync.Mutex**.
- Comunicación en tiempo real con **WebSockets**.
- Arquitectura modular backend + frontend.
- Manejo de rutas y servidor HTTP con `http.ServeMux`.

---

## 🔧 Tecnologías

- Go 1.22+
- WebSockets con [Gorilla WebSocket](https://github.com/gorilla/websocket)
- HTML y JavaScript para frontend simple
- Backend modular (`internal/horse`, `internal/race`, `api`)

---

## 🚀 Cómo ejecutar

1. Clonar el repositorio:

```bash
git clone https://github.com/tuusuario/horse-race.git
cd horse-race
```

2. Inicializar GO modules
```bash
go mod tidy
```

3. Ejecutar el servidor
```bash
go run ./cmd/server
```

4. Abrir el navegador
http://localhost:8080

---

🏁 Cómo funciona
Cada caballo se ejecuta en su propia goroutine, avanzando posiciones aleatorias.
La carrera se transmite en tiempo real mediante WebSockets.
El frontend muestra los caballos con barras de progreso y anuncia al ganador.
Se utiliza mutex para sincronizar acceso a la posición de los caballos y evitar race conditions.

---

📂 Estructura del proyecto
horse-race/
├── cmd/server/          # Entry point del servidor
├── internal/horse/      # Modelo de caballo
├── internal/race/       # Lógica de carrera
├── api/                 # Handlers de WebSocket
├── web/                 # Frontend HTML/JS
├── go.mod               # Dependencias
└── README.md            # Este archivo
