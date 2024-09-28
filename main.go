package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const GB_SIZE = 1024 * 1024 * 1024

func main() {
	fmt.Println("RAM Size (GB)")
	printRAMSize()
	fmt.Println()

	cpu := getCPU()
	fmt.Printf("Processor %s", cpu)

	drive := getDiskDrive()
	fmt.Printf("Drive %s", drive)

	printDDCapacity()
	fmt.Println()

	sn := getSerialNumber()
	fmt.Printf("%s", sn)

	os := getOSName()
	fmt.Printf("OS %s", os)

	gpu := getGPU()
	fmt.Printf("GPU %s", gpu)

	exitCMD()
}

func printRAMSize() {
	ram, err := exec.Command("wmic", "memorychip", "get", "capacity").Output()
	if err != nil {
		log.Fatal(err)
	}

	sticks := strings.Split(string(ram), "\n")

	for _, stick := range sticks[1:] {
		stick = strings.TrimSpace(stick)

		if stick != "" {
			memSize, err := strconv.ParseFloat(stick, 64)
			if err != nil {
				log.Fatal(err)
			}

			memSizeGB := memSize / GB_SIZE

			fmt.Printf("%.2f\n", memSizeGB)
		}
	}
}

func getCPU() string {
	cpu, err := exec.Command("wmic", "cpu", "get", "name").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(cpu)
}

func getDiskDrive() string {
	drive, err := exec.Command("wmic", "diskdrive", "get", "model").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(drive)
}

func printDDCapacity() {
	dd, err := exec.Command("wmic", "logicaldisk", "get", "size,freespace,caption").Output()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(dd), "\n")

	fmt.Println("Drive\tFree Space (GB)\tSize (GB)")

	for _, line := range lines[1:] {
		fields := strings.Fields(line)

		if len(fields) == 3 {
			caption := fields[0]

			freeSpaceBytes, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				log.Fatal(err)
			}

			sizeBytes, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				log.Fatal(err)
			}

			freeSpaceGB := freeSpaceBytes / GB_SIZE
			sizeGB := sizeBytes / GB_SIZE

			fmt.Printf("%s\t%.2f\t\t%.2f\n", caption, freeSpaceGB, sizeGB)
		}
	}
}

func getSerialNumber() string {
	sn, err := exec.Command("wmic", "bios", "get", "serialnumber").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(sn)
}

func getOSName() string {
	os, err := exec.Command("wmic", "os", "get", "name").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(os)
}

func getGPU() string {
	gpu, err := exec.Command("wmic", "path", "win32_videocontroller", "get", "name").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(gpu)
}

func exitCMD() {
	fmt.Print("Press Enter to exit")
	fmt.Scanln()
}
