# 🛒 Distributed Order Processing System

A **production-grade microservices project** built with **Go**, **Kafka**, **Cassandra**, **Envoy Proxy**, and **Kubernetes**.  
This project simulates a real-world e-commerce order processing system with **event-driven architecture**, **horizontal scalability**, and **fault tolerance**.

---

## 🚀 Overview

This project demonstrates how large-scale companies design **scalable, resilient, and maintainable backends**.  
It processes orders asynchronously, integrates multiple microservices, and uses a message queue to decouple services for better performance.

---

## 🏗 Architecture

```text
[ React Frontend ]
       |
       v
[ Envoy Proxy (API Gateway) ]
       |
+----------------------------+
|  Order Service (Go)        | ---> Kafka ---> Payment Service (Go)
|  - Create/Cancel Orders    |                  - Process payments
|  - Update Status           |                  - Publish payment result
|                            |
|----------------------------|
| Cassandra DB (Order Data)  |
+----------------------------+

Inventory Service (Go) <-- Kafka <-- Order & Payment Events
- Deducts stock
- Updates inventory status

All services run in Docker containers and are orchestrated with Kubernetes.
