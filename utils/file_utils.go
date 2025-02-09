package utils

import (
    "fmt"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
)

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

func DownloadFile(filename string, destination string) (string, error) {
    src, err := os.Open(filepath.Join(destination, filename))
    if err != nil {
        return "", err
    }
    defer src.Close()

    out, err := os.Create(filepath.Join("downloads", filename))
    if err != nil {
        return "", err
    }
    defer out.Close()

    _, err = io.Copy(out, src)
    return out.Name(), err
}