package utils

import (
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
)

// UploadFile uploads a file to the specified destination
func UploadFile(file *multipart.FileHeader, destination string) error {
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    out, err := os.Create(filepath.Join(destination, file.Filename))
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, src)
    return err
}
