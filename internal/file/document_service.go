package file

import (
	"fmt"
	"os"

	"github.com/mattmcsorley/jot/internal"
	"github.com/spf13/afero"
)

type DocumentService struct {
	appFs afero.Fs
}

// CreateDocumentService initializes a file based document service
func NewDocumentService(basePath string) *DocumentService {
	appFs := afero.NewBasePathFs(afero.NewOsFs(), basePath)
	return &DocumentService{appFs}
}

func (d *DocumentService) LoadDocument(path string) (*internal.Document, error) {
	file, err := d.appFs.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}
	defer file.Close()
	info, _ := file.Stat()
	p := make([]byte, info.Size())
	_, err1 := file.Read(p)
	if err1 != nil {
		fmt.Println(err1)
	}
	return &internal.Document{path, p}, err
}

func (d *DocumentService) SaveDocument(document *internal.Document) error {
	file, _ := d.appFs.OpenFile(document.FileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	_, err := file.Write(document.Contents)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *DocumentService) SyncDocuments() error {
	// no-op
	return nil
}
