# Smart IAM (Okta-like Access Management Platform)

An open-source identity and access management system built with **Go**, **Kafka**, and **React**. Designed to scale to millions of users using microservices, JWT-based auth, and cloud-native architecture.

---

## 🎯 Features

- ✅ Multi-tenant user management
- ✅ Role-based access control (RBAC)
- ✅ JWT Authentication & Authorization
- ✅ Kafka-based audit/event streaming
- ✅ Admin dashboard (React)
- ✅ Redis caching layer
- ✅ Scalable to 1M+ users
- ✅ Monitoring with Prometheus + Grafana
- ✅ Docker & Kubernetes support

---

## 🔧 Tech Stack

| Layer         | Stack                          |
| ------------- | ------------------------------ |
| Backend       | Go 1.21+, Fiber Framework       |
| Communication | Apache Kafka                   |
| Databases     | PostgreSQL (Users, Tenants)    |
| Frontend      | React + Tailwind CSS           |
| Caching       | Redis                          |
| Observability | Prometheus, Grafana, Loki      |
| CI/CD         | Docker Compose, Kubernetes     |
| Load Testing  | k6, Locust                     |

---

## 🗂️ Folder Structure (v0.1 Skeleton)

```bash
smart-iam/
├── auth-service/          # JWT auth, signup/login (Go)
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod
├── tenant-service/        # Multi-tenant + role mgmt (Go)
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod
├── gateway-service/       # API Gateway + RBAC enforcement (Go)
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod
├── shared-libs/           # Common structs & utils
│   └── (to be added)
├── audit-service/         # Kafka event consumer/logger (Go)
│   └── (coming soon)
├── frontend/              # React UI (Scaffolded)
├── infra/
│   └── k8s/               # Kubernetes manifests (WIP)
├── docker-compose.yml     # Spin up all services
└── README.md              # Root doc
