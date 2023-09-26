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
	ttl     = time.Second * 8
	checkId = "check_health"
)

type Service struct {
	consulClient *api.Client
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

func (s *Service) Start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	s.registerService()
	go s.healthCheck()
	s.acceptLoop(ln)
}

func (s *Service) acceptLoop(ln net.Listener) {
	for {
		_, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
	}
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

func (s *Service) registerService() {
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(),
		TLSSkipVerify:                  true,
		TTL:                            ttl.String(),
		CheckID:                        checkId,
	}

	register := &api.AgentServiceRegistration{
		ID:      "login_service",
		Name:    "Login Service",
		Tags:    []string{"login"},
		Address: "127.0.0.1",
		Port:    3000,
		Check:   check,
	}

	query := map[string]any{
		"type":        "service",
		"service":     "Login Service",
		"passingonly": true,
	}
	plan, err := watch.Parse(query)
	if err != nil {
		log.Fatal(err)
	}

	plan.HybridHandler = func(index watch.BlockingParamVal, result any) {
		switch msg := result.(type) {
		case []*api.ServiceEntry:
			for _, entry := range msg {
				fmt.Print("new service", entry.Service)
			}
		}
	}

	go func() {
		_ = plan.RunWithConfig("", &api.Config{})
	}()

	if err = s.consulClient.Agent().ServiceRegister(register); err != nil {
		log.Fatal(err)
	}

}
