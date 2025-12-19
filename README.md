# 🔗 Go URL Shortener

A full-stack URL shortener application built using **Go** with a clean backend architecture and a modern, lightweight frontend.  
It converts long URLs into short, shareable links with persistent storage.

---

## 🚀 Features

- Shorten long URLs into unique short links
- Redirect short URLs to the original destination
- Persistent storage using SQLite
- RESTful API built with Go
- Clean layered architecture (handler, service, repository)
- CORS-enabled backend
- Modern frontend UI for easy testing
- Copy-to-clipboard support
- Docker-ready backend

---

## 🖥️ Demo

![Go URL Shortener Demo](./screenshots/demo.png)

---

## 🛠️ Tech Stack

### Backend
- Go
- Gorilla Mux
- SQLite
- REST API
- Clean Architecture

### Frontend
- HTML
- CSS
- Vanilla JavaScript (Fetch API)

---

## 📂 Project Structure

```text
url-shortener/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   └── middleware/
│       ├── cors.go
│       └── logger.go
├── migrations/
│   └── schema.sql
├── frontend/
│   ├── index.html
│   ├── style.css
│   └── script.js
├── screenshots/
│   └── demo.png
├── Dockerfile
├── go.mod
├── go.sum
├── README.md
└── .gitignore