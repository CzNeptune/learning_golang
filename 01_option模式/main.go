package main

import (
	"fmt"
)

type Model struct {
	Tempreture float64
	MaxTokens  int
}

func (m Model) Run(message string) {
	fmt.Printf("Running model with tempreture: %f, max tokens: %d, message: %s\n", m.Tempreture, m.MaxTokens, message)
}

type Tool struct {
}

type Agent struct {
	Model Model
	Tools []Tool
}

type ModelOption func(*Model)

func WithTempreture(tempreture float64) ModelOption {
	return func(m *Model) {
		m.Tempreture = tempreture
	}
}

func WithMaxTokens(maxTokens int) ModelOption {
	return func(m *Model) {
		m.MaxTokens = maxTokens
	}
}

func (a Agent) Query(message string, options ...ModelOption) {
	for _, option := range options {
		option(&a.Model)
	}
	a.Model.Run(message)

}

func main() {
	agent := Agent{
		Model: Model{
			Tempreture: 0.5,
			MaxTokens:  1024,
		},
	}
	agent.Query("Hello")

	agent.Query("Hello", WithTempreture(0.7), WithMaxTokens(2048))
	agent.Query("Hello", WithTempreture(0.9))
}
