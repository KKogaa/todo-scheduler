package clients

import "time"

type JsonClient struct {
}

type Pomo struct {
	StartTime time.Time
}

func (j JsonClient) CreateNewFile(fileName string) error {
	return nil
}

func (j JsonClient) UpdateFile(fileName string) error {
	return nil
}

func (j JsonClient) DeleteFile(fileName string) error {
	return nil
}

func (j JsonClient) GetFile(fileName string) error {
	return nil
}
