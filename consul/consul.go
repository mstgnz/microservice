package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

const (
	ttl     = time.Second * 10
	checkId = "check_health"
)

type Service struct {
	consulClient *api.Client
	services     []api.AgentServiceRegistration
}

func NewService() *Service {
	client, err := api.NewClient(&api.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Service{
		consulClient: client,
	}
}

func (s *Service) AddService(service api.AgentServiceRegistration) {
	s.services = append(s.services, service)
}

func (s *Service) Start() {
	s.registerServices()
	go s.healthCheck()
	s.acceptLoop()
}

func (s *Service) healthCheck() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		err := s.consulClient.Agent().UpdateTTL(checkId, "update", api.HealthPassing)
		if err != nil {
			log.Fatal(err)
		}
		<-ticker.C
	}
}

func (s *Service) registerServices() {

	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(),
		TLSSkipVerify:                  true,
		TTL:                            ttl.String(),
		CheckID:                        checkId,
	}
	for _, service := range s.services {

		service.Check = check

		s.query(map[string]any{
			"type":        "service",
			"service":     service.Name,
			"passingonly": true,
		})

		if err := s.consulClient.Agent().ServiceRegister(&service); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Service) query(query map[string]any) {

	plan, err := watch.Parse(query)
	if err != nil {
		log.Fatal(err)
	}

	plan.HybridHandler = func(index watch.BlockingParamVal, result any) {
		switch msg := result.(type) {
		case []*api.ServiceEntry:
			for _, entry := range msg {
				fmt.Print("service", entry.Service)
			}
		}
	}

	go func() {
		_ = plan.RunWithConfig("", &api.Config{})
	}()
}

func (s *Service) acceptLoop() {
	ln, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, err = ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
	}
}
