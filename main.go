package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func pastMonthCommits() (int, error) {
	wc := exec.Command("wc", "-l")
	wcPipe, err := wc.StdinPipe()
	if err != nil {
		fmt.Println("error creating pipe")
		return 0, err
	}
	commits, err := exec.Command("git", "log", "--since='last month'", "pretty=format:'%s'").Output()
	if err != nil {
		fmt.Println("error executing git log")
		return 0, err
	}
	go func() {
		defer wcPipe.Close()
		wcPipe.Write(commits)

	}()
	amount, err := wc.CombinedOutput()
	if err != nil {
		fmt.Println("error getting output")
		return 0, err
	}
	return int(binary.BigEndian.Uint16(amount)), nil
}

func main() {
	salary := os.Getenv("SALARY")
	fmt.Println("HERE")
	fmt.Println(salary)
	intSalary, err := strconv.Atoi(salary)
	if err != nil {
		println("WOOLOL")
		fmt.Printf("Check your env variable is a valid int: %s", salary)
		os.Exit(1)
	}
	commits, err := pastMonthCommits()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		os.Exit(1)
	}
	fmt.Printf("YOU WON %d WITH THIS COMMIT\n", intSalary/commits)
}
