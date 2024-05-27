package interface_learning

import "fmt"

type StorageProvider interface {
	UploadFile(fileName string, data []byte) error
	DownloadFile(fileName string) ([]byte, error)
}

type LocalStorageProvider struct {
	UploadPath string `json:"upload_path"`
}

type CloudStorageProvider struct {
	Key           string `json:"key"`
	CloudProvider string `json:"cloud_provider"`
}

func NewLocalStorageProvider(uploadLoad string) *LocalStorageProvider {
	return &LocalStorageProvider{
		UploadPath: uploadLoad,
	}
}

func NewCloudStorageProvider(key, cloudProvider string) *CloudStorageProvider {
	return &CloudStorageProvider{
		Key:           key,
		CloudProvider: cloudProvider,
	}
}

func (c *LocalStorageProvider) UploadFile(fileName string, data []byte) error {
	fmt.Printf("uploaded  fileName: %s data: %s", fileName, data)
	return nil
}
func (c *CloudStorageProvider) UploadFile(fileName string, data []byte) error {
	key := c.Key
	cloudProvider := c.CloudProvider
	fmt.Printf("uploaded file for key %s cloudprovider %s fileName: %s data: %s", key, cloudProvider, fileName, data)
	return nil
}
func (c *LocalStorageProvider) DownloadFile(fileName string) ([]byte, error) {
	return nil, nil
}
func (c *CloudStorageProvider) DownloadFile(fileName string) ([]byte, error) {
	return nil, nil
}

func InterfaceExample() {
	localStorage := NewLocalStorageProvider("/path/to/local_storage")
	cloudStorage := NewCloudStorageProvider("key", "cloud_provider")
	data := []byte("ayush")
	localStorage.UploadFile("test.txt", data)
	localStorage.DownloadFile(localStorage.UploadPath)

	cloudStorage.UploadFile("test.txt", data)
	cloudStorage.DownloadFile(localStorage.UploadPath)
	map1 := map[string]int{
		"first":  11,
		"second": 22,
	}
	map2 := map[string]float64{
		"first":  1.2,
		"second": 22.2,
	}
	answer := make(chan float64)
	go RoundPercent(2, 3, answer)
	x := <-answer
	fmt.Println("\n", x)
	fmt.Println(Sum(map1))
	fmt.Println(Sum(map2))

}

type number interface {
	int | int32 | int64 | float32 | float64
}

func Sum[A string, B number](m map[string]B) B {
	var sum B
	sum = 0
	for _, v := range m {
		sum = sum + v
	}
	return sum
}

func RoundPercent[A number, B number](numerator A, denominator B, answer chan float64) {
	if denominator != 0 {
		percentage := float64(numerator) / float64(denominator) * 100
		answer <- percentage
	}
	answer <- 0
}
