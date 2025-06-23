package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func AskForConfirm() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 用于接收系统中断信号（Ctrl+C）
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 用户输入通道
	inputChan := make(chan string, 1)

	// 开 goroutine 读取用户输入
	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Confirm (y/n): ")
		for {
			text, _ := reader.ReadString('\n')
			text = strings.ToLower(strings.TrimSpace(text))
			inputChan <- text
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nTimeout: No input received within 10 seconds.")
			return false, ctx.Err()

		case sig := <-sigChan:
			fmt.Printf("\nReceived signal: %v, exiting.\n", sig)
			return false, fmt.Errorf("interrupted by signal: %v", sig)

		case input := <-inputChan:
			switch input {
			case "y":
				return true, nil
			case "n":
				return false, nil
			default:
				fmt.Print("Invalid input, please enter 'y' or 'n':")
			}
		}
	}
}

func main() {
	proceed, err := AskForConfirm()
	if err != nil {
		fmt.Println("Exited with error:", err)
		return
	}
	if proceed {
		fmt.Println("User chose to continue.")
	} else {
		fmt.Println("User chose to stop.")
	}
}
