package internal

// Document hold all information pertaining to a jot document
type Document struct {
	FileName string
	Contents []byte
}

func (d *Document) Write(p []byte) (int, error) {
	if len(d.Contents) < 1 {
		d.Contents = p
	} else {
		d.Contents = append(d.Contents, p...)
	}
	return len(p), nil
}

// DocumentService provides document operations
type DocumentService interface {
	LoadDocument(path string) (*Document, error)
	SaveDocument(d *Document) error
	SyncDocuments() error
}

// Sync service keeps files in sync
type SyncService interface {
	Sync()
}
