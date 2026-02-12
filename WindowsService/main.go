package main

import (
	"log"
	"time"
	"os"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
)

const serviceName = "GoHeartBeatService"

func main() {
	isService, err := svc.IsWindowsService()

	if err != nil {
		log.Fatalf("Failed to detect service context: %v", err)
	}

	if isService {
		runService()
		return
	}

	debug.Run(serviceName, &myService{})
}


func runService() {
	err := svc.Run(serviceName, &myService{})
	if err != nil {
		log.Fatalf("Service failed: %v", err)
	}
}

type myService struct{}

func (m *myService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
	s <- svc.Status{State: svc.StartPending}

	logFile, err := os.OpenFile("D:\\go-service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 
	if err != nil {
		return false, 1
	}

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println("Service hearbeat engine starting")

	s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

	ticket := time.NewTicker(5 * time.Second)
	defer ticket.Stop()

loop:
	for {
		select {
		case <-ticket.C:
			logger.Println("Heartbate: service is alive")
		case c:= <-r:
			switch c.Cmd {
			case svc.Stop, svc.Shutdown:
				logger.Println("Service stopping...")
				break loop
			}
		}
	}


	s <- svc.Status{State: svc.StopPending}
	logger.Println("Service stopped")
	return false, 0
}