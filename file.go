package gutil

import (
	"encoding/csv"
	"os"
)

// WriteCsv 写入数据到csv文件
func WriteCsv(allData [][]string, file string) error {
	// 需要写入的文件
	csvWriteFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer csvWriteFile.Close()

	// 防止乱码
	csvWriteFile.WriteString("\xEF\xBB\xBF")

	csvWriter := csv.NewWriter(csvWriteFile)
	err = csvWriter.WriteAll(allData)
	if err != nil {
		return err
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}
	return nil
}

// ReadCsv 读取csv文件
func ReadCsv(file string) ([][]string, error) {
	var (
		rows [][]string
		err  error
	)
	csvFile, err := os.Open(file)
	if err != nil {
		return rows, err
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return rows, err
	}
	return rows
}
