package main

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://github.com/AlexeyBazhin/wbL2/blob/master/develop/dev08/task_test.go"
	filename := "task_test"

	err := downloadFile(filename,url, "go")
	if err != nil {
		t.Errorf("Ошибка при скачивании файла: %s", err.Error())
	}

	// Проверяем, что файл был создан
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("Файл не был создан")
	}

}