package usecases

import (
	"io"

	"google.golang.org/api/drive/v3"
)

func CreateFolder(svc *drive.Service, folderName string, parentId string) (*drive.File, error) {
	f := &drive.File{
		Name:     folderName,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentId},
	}

	fileSvc := drive.NewFilesService(svc)
	file, err := fileSvc.Create(f).Do()
	if err != nil {
		return nil, err
	}

	return file, nil
}

// list folder by name and id
func ListFolders(svc *drive.Service, pageSize *int64) (*drive.FileList, error) {
	ps := int64(10)
	if pageSize != nil {
		ps = *pageSize
	}
	fileSvc := drive.NewFilesService(svc)
	fl, err := fileSvc.List().PageSize(ps).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return nil, err
	}
	return fl, nil
}

func CreateDocument(svc *drive.Service, docName string, mimeType string, doc io.Reader, parentId string) (*drive.File, error) {
	f := &drive.File{
		Name:     docName,
		MimeType: mimeType,
		Parents:  []string{parentId},
	}

	fileSvc := drive.NewFilesService(svc)
	file, err := fileSvc.Create(f).Media(doc).Do()
	if err != nil {
		return nil, err
	}

	return file, err

}
