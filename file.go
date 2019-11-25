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
	rows, err = csvReader.ReadAll()
	if err != nil {
		return rows, err
	}
	return rows, nil
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
func ReadMapToFile(fileName string) map[string]string {
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

// DewightDat .dat文件去重处理
func DewightDat(file1, file2 string, resfile string) error {
	// 样本(需要过滤掉的数据)
	f1, err := os.Open(file1)
	if err != nil {
		return err
	}
	defer f1.Close()

	var (
		equs   map[string]string
		result map[string]string
	)
	equs = make(map[string]string)
	rd := bufio.NewReader(f1)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		equs[line] = line
	}

	// 所有数据
	f, err := os.Open(file2)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 所有数据
	result = make(map[string]string)
	rd1 := bufio.NewReader(f)
	for {
		line, err := rd1.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		// 没有在样本里面
		if _, ok := equs[line]; !ok {
			result[line] = line
		}
	}
	err = WriteMaptoFile(result, resfile)
	if err != nil {
		return err
	}
	return nil
}
