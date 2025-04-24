# 🧠 Targeting Engine Service

A high-performance, horizontally scalable ad-serving microservice built in Go using [Go-Kit](https://gokit.io). Routes campaigns to user requests based on dynamic targeting rules.

---

## 🚀 Problem Statement

Build a lightweight and scalable delivery engine to:
- Handle **billions of delivery requests**
- Match **~1000 campaigns** based on targeting rules
- Maintain **low latency** and **high availability**

---

## 🛠️ Technologies & Architecture

| Component | Description |
|----------|-------------|
| **Go + Go-Kit** | Service layer, structured logging, middleware |
| **Postgres** | Source of truth for campaigns and targeting rules |
| **Redis** | Read-optimized cache for active campaigns and rules |
| **Docker Compose** | Local dev orchestration with persistent volumes |
| **Pg LISTEN/NOTIFY** | Realtime Redis cache refresh |
| **Circuit Breaker** | Fault tolerance for Redis/DB using `gobreaker` |
| **Rate Limiting** | API protection via `golang.org/x/time/rate` |
| **Panic Recovery** | Middleware-based graceful failure handling |

---

## 🔄 How It Works

1. Campaigns and rules are stored in Postgres.
2. Redis is used to serve all GET requests for `/v1/delivery`.
3. Postgres triggers send updates via `NOTIFY`.
4. Go service listens on channels (`campaign_change`, `targeting_rule_change`) to auto-refresh Redis cache.
5. API logic matches requests to cached campaigns using inclusion/exclusion rules.
6. Structured JSON logs with metrics are captured for all requests.

---

## ⚙️ Running Locally with Docker Compose

### 1. Clone + Build
```bash
git clone https://github.com/sashankdk/targetEngineAPI.git
cd targeting-engine-service
chmod +x run_with_tests.sh
./run_with_tests.sh
```
### 2. Interact with Services

|Component |Access|
|----------|------|
|API|	http://localhost:8080/v1/delivery|
|Redis CLI	|docker exec -it redis redis-cli|
|Postgres CLI	|docker exec -it postgres psql -U gg_user -d gg_campaigns|
|Logs	|docker compose logs -f api|

### 3. Sample Requests
```bash
curl "http://localhost:8080/v1/delivery?app=com.abc.xyz&country=germany&os=android"
```

# ⚙️ Future Enhancements

🔹 Prometheus + Grafana	Metrics dashboard (latency, throughput, circuit states)
🔹 Kubernetes Manifests	Autoscaling pods, resource limits
🔹 Tracing with OpenTelemetry	End-to-end distributed tracing