package gutil

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
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

// WriteMaptoFile
func WriteMaptoFile(m map[string]string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range m {
		lineStr := fmt.Sprintf("%s", v)
		fmt.Fprintln(w, lineStr)
	}
	return w.Flush()
}

// ReadMapToFile
func ReadMapToFile(fileName) map[string]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var (
		result map[string]string
	)
	result = make(map[string]string)

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		result[line] = line
	}
	return result
}
