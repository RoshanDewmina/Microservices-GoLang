

# Microservices Application

This project is a comprehensive example of a microservices architecture implemented in Go. It includes separate services for user authentication, order management, payments, notifications, and inventory. Each service is containerized using Docker and deployed on Kubernetes for scalability and resilience.

## Features

- **Authentication Service:**  
  Handles user authentication, issues JWT tokens, and verifies user identities.

- **Order Service:**  
  Manages customer orders and coordinates with the payment, inventory, and notification services.

- **Payment Service:**  
  Processes payments and integrates with external payment gateways.

- **Notification Service:**  
  Sends notifications to users about orders, payments, and inventory updates.

- **Inventory Service:**  
  Tracks product stock levels and provides inventory data to other services.

## Technologies Used

- **Programming Language:** Go (Golang)  
- **Containerization:** Docker  
- **Orchestration:** Kubernetes  
- **Monitoring:** Prometheus, Grafana  
- **API Communication:** REST

## Project Structure

```
app/
  auth-service/
    main.go
    Dockerfile
  order-service/
    main.go
    Dockerfile
  payment-service/
    main.go
    Dockerfile
  notification-service/
    main.go
    Dockerfile
  inventory-service/
    main.go
    Dockerfile
k8s/
  auth-service.yaml
  order-service.yaml
  payment-service.yaml
  notification-service.yaml
  inventory-service.yaml
```

## Getting Started

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/RoshanDewmina/microservices-app.git
   cd microservices-app
   ```

2. **Build the Docker Images:**
   ```bash
   docker build -t auth-service:latest ./app/auth-service
   docker build -t order-service:latest ./app/order-service
   docker build -t payment-service:latest ./app/payment-service
   docker build -t notification-service:latest ./app/notification-service
   docker build -t inventory-service:latest ./app/inventory-service
   ```

3. **Deploy on Kubernetes:**
   ```bash
   kubectl apply -f k8s/auth-service.yaml
   kubectl apply -f k8s/order-service.yaml
   kubectl apply -f k8s/payment-service.yaml
   kubectl apply -f k8s/notification-service.yaml
   kubectl apply -f k8s/inventory-service.yaml
   ```

4. **Verify the Deployment:**
   ```bash
   kubectl get pods
   kubectl get services
   ```

5. **Access the Services:**
   Use `kubectl port-forward` or a Kubernetes ingress to expose the services and interact with them through their RESTful APIs.
