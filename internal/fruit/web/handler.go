package web

import (
	"bufio"
	"errors"
	"fmt"
	"gitlab.cmpayments.local/creditcard/fruit-price-calculator/internal/entities"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type service interface {
	CalculateAveragePrice(fruits entities.Fruits) entities.Fruits
}

type fruitHandler struct {
	service service
}

func NewFruitHandler(service service) *fruitHandler {
	return &fruitHandler{service: service}
}

func (h fruitHandler) PrintFruitsAverageFromFile(args []string) {

	filePath, err := getFilePath(args)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("failed reading given file path", err)
	}
	defer file.Close()

	fruits, err := loadFileContent(file)
	if err != nil {
		log.Println("failed loading fruits file: ", err)
		return
	}

	for _, fruit := range h.service.CalculateAveragePrice(fruits) {
		fmt.Printf("Average price of a %s is %.2f", fruit.Type, fruit.Price)
		fmt.Println()
	}
}

func getFilePath(args []string) (string, error) {
	switch {
	case len(args) > 1:
		filePath := args[1]
		if !filepath.IsAbs(filePath) {
			filePath, _ = filepath.Abs(filePath)
		}

		if fileInfo, err := os.Stat(filePath); os.IsNotExist(err) || fileInfo.IsDir() {
			return "", fmt.Errorf("invalid given file path %s", filePath)
		}
		return filePath, nil
	default:
		return "", errors.New("program missing argument: file path required")
	}
}

func loadFileContent(file *os.File) (entities.Fruits, error) {
	var (
		scanner = bufio.NewScanner(file)
		fruits  entities.Fruits
	)

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ",")

		if len(data) < 2 {
			return nil, fmt.Errorf("error parsing file line %s", line)
		}

		price, err := strconv.ParseFloat(strings.TrimSpace(data[1]), 32)
		if err != nil {
			return nil, fmt.Errorf("error parsing fruit: %s price: %s", data[0], data[1])
		}

		fruits = append(fruits, entities.Fruit{Type: data[0], Price: price})
	}

	return fruits, nil
}
