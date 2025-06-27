package main

import "fmt"

// Interface for logging
type Logger interface {
	Log(message string)
}

// Concrete implementation
// This defines a new struct type called ConsoleLogger.

// It has no fields ({} is empty), so it's just a type to implement the interface.
type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

// Business logic depends only on interface
type Service struct {
	logger Logger
}

func (s Service) DoWork() {
	s.logger.Log("Service is doing work")
}

func main() {
	cl := ConsoleLogger{}
	cl.Log("print")
	svc := Service{logger: cl}
	svc.DoWork()
}

// The Service struct has a field logger of interface type Logger.

// In Go, any value that implements the methods required by an interface can be assigned to a variable of that interface type.

// cl is of type ConsoleLogger, which has a method Log(string), so it satisfies the Logger interface.
