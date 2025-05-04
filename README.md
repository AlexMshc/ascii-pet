# ASCII Pet App

## Требования

- **Go** 1.24
- **Node.js** ≥ 18 и **npm**  
- (опционально) **Docker** & **Docker Compose**

---

## Запуск в режиме разработки

### 1. Бэкенд

```bash
cd backend
go run cmd/server/main.go
````

### 2. Фронтенд

```bash
cd frontend
npm install
npm run dev          
```

* Откройте в браузере: `http://localhost:5173`

---

## Запуск через Docker Compose

```bash
docker-compose up --build
```

* **Откройте в браузере:**:  `http://localhost/`
---