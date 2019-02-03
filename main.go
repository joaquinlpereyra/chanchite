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
		return 0, err
	}
	commits, err := exec.Command("git", "log", "--since=\"last month\"", "--prety=format:'%s'").Output()
	if err != nil {
		return 0, err
	}
	wcPipe.Write(commits)
	amount, err := wc.CombinedOutput()
	if err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint16(amount)), nil
}

func main() {
	salary := os.Getenv("salary")
	if salary == "" {
		salary = os.Getenv("notmysalary")
	}
	if salary == "" {
		salary = os.Getenv("chanchite")
	}
	intSalary, err := strconv.Atoi(salary)
	if err != nil {
		fmt.Printf("Check your env variable is a valid int")
		os.Exit(1)
	}
	commits, err := pastMonthCommits()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		os.Exit(1)
	}
	fmt.Printf("YOU WON %d WITH THIS COMMIT\n", intSalary/commits)
}
