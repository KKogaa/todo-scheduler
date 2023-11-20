package clients

type FileClient interface {
	CreateNewFile(fileName string) error
	UpdateFile(fileName string) error
	DeleteFile(fileName string) error
	GetFile(fileName string) error
}
