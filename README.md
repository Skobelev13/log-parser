# Log Parser Service

Микросервис на Go для парсинга логов и CSV topology-файлов.

---

# Возможности

- health endpoint
- parsing text logs
- parsing CSV topology files
- JSON API
- layered architecture

---

# Стек

- Go
- net/http
- encoding/json
- encoding/csv

---

# Структура проекта

```text
log-parser/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── data/
│   ├── test.log
│   └── ibdiagnet2.db_csv
│
├── internal/
│   ├── api/
│   │   └── handler.go
│   │
│   ├── model/
│   │   ├── node.go
│   │   └── link.go
│   │
│   ├── parser/
│   │   ├── parser.go
│   │   └── csv_parser.go
│   │
│   ├── service/
│   │   └── parser_service.go
│   │
│   ├── logger/
│   ├── storage/
│   └── topology/
│
├── migrations/
│   └── 001_init.sql
│
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md




Запуск
go run ./cmd/app

Сервис стартует на:

localhost:8080
Проверка health endpoint
curl localhost:8080/health

Ответ:

ok
Parse text log
Запрос
curl -X POST localhost:8080/api/v1/parse \
-H "Content-Type: application/json" \
-d '{"path":"data/test.log"}'
Ответ
{
  "lines": [
    {
      "name": "switch-1",
      "ip": "10.0.0.1",
      "port": "eth0"
    },
    {
      "name": "host-1",
      "ip": "10.0.0.2",
      "port": "eth1"
    }
  ],
  "path": "data/test.log",
  "status": "accepted"
}
Parse CSV topology file
Запрос
curl localhost:8080/api/v1/csv
Ответ
[
  {
    "switch": "SWITCH_1",
    "port": "65",
    "peer": "2"
  },
  {
    "switch": "SWITCH_2",
    "port": "65",
    "peer": "2"
  }
]
Архитектура
HTTP request
    ↓
handler
    ↓
service
    ↓
parser
    ↓
model
    ↓
JSON response
Endpoints
GET /health

Healthcheck endpoint.

POST /api/v1/parse

Парсинг текстового лог-файла.

GET /api/v1/csv

Парсинг CSV topology файла.

Обработка ошибок

Сервис возвращает ошибки при:

отсутствии файла
битом JSON
поврежденном CSV
внутренних ошибках сервера
Пример запуска
go run ./cmd/app
curl localhost:8080/health
curl localhost:8080/api/v1/csv
