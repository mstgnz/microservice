package main

import (
	"github.com/hashicorp/consul/api"
)

func main() {
	// Init
	s := NewService()

	// Add Service
	s.AddService(api.AgentServiceRegistration{
		ID:      "postgres",
		Name:    "Postgres",
		Tags:    []string{"database", "relational database"},
		Address: "microservice-postgres",
		Port:    5432,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "mongo",
		Name:    "Mongo",
		Tags:    []string{"database", "document database"},
		Address: "microservice-mongo",
		Port:    27017,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "redis",
		Name:    "Redis",
		Tags:    []string{"database", "key-value database"},
		Address: "microservice-redis",
		Port:    6379,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "kafka",
		Name:    "Kafka",
		Tags:    []string{"queue", "message broker"},
		Address: "microservice-kafka",
		Port:    9092,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "listener",
		Name:    "Listener",
		Tags:    []string{"service", "consumer"},
		Address: "microservice-listener",
		Port:    80,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "mail",
		Name:    "Mail",
		Tags:    []string{"service", "consumer"},
		Address: "microservice-mail",
		Port:    80,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "sms",
		Name:    "Sms",
		Tags:    []string{"service", "consumer"},
		Address: "microservice-sms",
		Port:    80,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "logger",
		Name:    "Logger",
		Tags:    []string{"service", "logger"},
		Address: "microservice-logger",
		Port:    80,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "auth",
		Name:    "Auth",
		Tags:    []string{"service", "backend"},
		Address: "microservice-auth",
		Port:    80,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "blog-api",
		Name:    "Blog Api",
		Tags:    []string{"service", "backend"},
		Address: "microservice-blog-api",
		Port:    80,
	})

	s.AddService(api.AgentServiceRegistration{
		ID:      "blog-web",
		Name:    "Blog Web",
		Tags:    []string{"client", "frontend"},
		Address: "microservice-blog-web",
		Port:    80,
	})

	// start
	s.Start()
}
