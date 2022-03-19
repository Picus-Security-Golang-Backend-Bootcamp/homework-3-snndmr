package helper

import (
	"encoding/csv"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-snndmr/internal/domain/book"
	"os"
	"strconv"
	"sync"
)

// FillDBFromCSV To fill database with given csv file.
func FillDBFromCSV(repository *book.Repository, path string) error {
	jobs := make(chan []string, 5)
	results := make(chan book.Book)
	waitGroup := new(sync.WaitGroup)

	for w := 1; w <= 3; w++ {
		waitGroup.Add(1)
		go convertToBookStruct(jobs, results, waitGroup)
	}

	go func() {
		file, _ := os.Open(path)
		defer file.Close()

		lines, _ := csv.NewReader(file).ReadAll()
		isFirstRow := true
		for _, line := range lines {
			if isFirstRow {
				isFirstRow = false
				continue
			}

			jobs <- line
		}
		close(jobs)
	}()

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	repository.InsertSampleData(results)
	return nil
}

// To convert string to book struct.
func convertToBookStruct(jobs <-chan []string, results chan<- book.Book, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for job := range jobs {
		stockId, _ := strconv.Atoi(job[1])
		pageCount, _ := strconv.Atoi(job[3])
		stockCount, _ := strconv.Atoi(job[5])
		price, _ := strconv.ParseFloat(job[4], 32)
		results <- book.Book{ID: 0, Title: job[0], StockId: stockId, ISBN: job[2], PageCount: pageCount, Price: float32(price), StockCount: stockCount, Author: job[6], IsDeleted: false}
	}
}
