package simpdb

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type SimpDBProvider struct {
	dbPath string
}

func NewSimpDBProvider(dbPath string) *SimpDBProvider {
	return &SimpDBProvider{dbPath: dbPath}
}

func (p *SimpDBProvider) CreateOne(newRecord string) error {
	filePath := fmt.Sprint(p.dbPath, string(newRecord[0]))
	// create a "map" storage for records (emails) for quicker lookups (very theoretically)
	// emails starting with "a" will be stored in file "a", starting with "b" in "b" etc.
	partition, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0643)
	if err != nil {
		return fmt.Errorf("unexpected error: failed to open/create a db partition: %w", err)
	}
	defer partition.Close()

	scanner := bufio.NewScanner(partition)
	// check if a record is already stored in file
	for scanner.Scan() {
		record := scanner.Text()
		if record == newRecord {
			return fmt.Errorf("db already contains record: %s", newRecord)
		}
	}

	_, err = fmt.Fprintf(partition, "%s\n", newRecord)
	if err != nil {
		return fmt.Errorf("unexpected error: failed to write a new record to db partition: %w", err)
	}

	return nil
}

func (p *SimpDBProvider) GetAll() (*[]string, error) {
	var records []string
	err := filepath.Walk(p.dbPath, p.readPartitionFunc(&records))
	if err != nil {
		return nil, err
	}

	return &records, nil
}

func (p *SimpDBProvider) readPartitionFunc(
	records *[]string,
) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("unexpected error: failed to walk a db partition: %w", err)
		}

		if info.IsDir() {
			return nil
		}

		partition, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("unexpected error: failed to read a partition: %w", err)
		}
		defer partition.Close()

		scanner := bufio.NewScanner(partition)
		for scanner.Scan() {
			record := scanner.Text()
			*records = append(*records, record)
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("unexpected error: scanner encountered error: %w", err)
		}

		return nil
	}
}
