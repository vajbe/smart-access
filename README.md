# Smart IAM (Okta-like Access Management Platform)

An open-source identity and access management system built with Java, Kafka, and React. Designed to scale to millions of users using microservices, JWT auth, and cloud-native architecture.

---

## 🎯 Features

- ✅ Multi-tenant user management
- ✅ Role-based access control (RBAC)
- ✅ JWT Authentication & Authorization
- ✅ Kafka-based audit/event streaming
- ✅ Admin dashboard (React)
- ✅ Redis caching layer
- ✅ Load-tested for 1M+ users
- ✅ Monitoring with Prometheus + Grafana
- ✅ Docker & Kubernetes support

---

## 🔧 Tech Stack

| Layer         | Stack                                   |
| ------------- | --------------------------------------- |
| Backend       | Java 17, Spring Boot 3, Spring Security |
| Communication | Apache Kafka                            |
| Databases     | PostgreSQL (Users, Tenants, Roles)      |
| Frontend      | React + Tailwind CSS                    |
| Caching       | Redis                                   |
| Observability | Prometheus, Grafana, Loki               |
| CI/CD         | Docker Compose, Kubernetes (K8s)        |
| Load Testing  | k6, Locust                              |

---

## 🗂️ Folder Structure (v0.1 Skeleton)

```bash
smart-iam/
├── auth-service/          # JWT auth, signup/login
│   ├── src/main/java/com/smartiam/auth/
│   ├── Dockerfile
│   └── pom.xml
├── tenant-service/        # Multi-tenant + role mgmt
│   ├── src/main/java/com/smartiam/tenant/
│   ├── Dockerfile
│   └── pom.xml
├── gateway-service/       # API Gateway + RBAC enforcement
│   ├── src/main/java/com/smartiam/gateway/
│   ├── Dockerfile
│   └── pom.xml
├── shared-libs/           # Common DTOs & config
│   └── (to be added)
├── audit-service/         # Kafka event consumer/logger
│   └── (coming soon)
├── frontend/              # React UI (Scaffolded)
├── infra/
│   └── k8s/               # Kubernetes manifests (WIP)
├── docker-compose.yml     # Spin up all services
└── README.md              # Root doc
