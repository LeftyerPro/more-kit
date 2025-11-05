/*
 * kitFile.go File-difine
 * by:Leftyer
 * dt:2025-11-01
 */
package kitFile

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

/* File-config-difine by:Leftyer dt:2025-11-01 */
type FileConfig struct{}

/* Folder-IsExist-difine by:Leftyer dt:2025-11-01 */
func IsFolderExist(path string, isClear int) bool {
	if st, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, os.ModePerm) == nil
		}
		return false
	} else if !st.IsDir() {
		return false
	}

	if isClear == 1 {
		entries, _ := os.ReadDir(path)
		for _, e := range entries {
			if os.RemoveAll(filepath.Join(path, e.Name())) != nil {
				return false
			}
		}
	}
	return true
}

/* Folder-Cpoy-difine by:Leftyer dt:2025-11-01 */
func CopyFolder(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		switch {
		case info.IsDir():
			return os.MkdirAll(target, info.Mode())
		default:
			return CopyFile(path, target)
		}
	})
}

/* File-IsExist-difine by:Leftyer dt:2025-11-01 */
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

/* File-Cpoy-difine by:Leftyer dt:2025-11-01 */
func CopyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = io.Copy(d, s)
	d.Close()
	return err
}

/* File-Read-difine by:Leftyer dt:2025-11-01 */
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err == nil {
		return string(data), nil
	}
	return "", err
}

/* File-Write-difine by:Leftyer dt:2025-11-01 */
func WriteFile(filePath string, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}

/* File-Read-Json-difine by:Leftyer dt:2025-11-01 */
func ReadJSON[T any](path string) (T, error) {
	var zero T
	b, err := os.ReadFile(path)
	if err != nil {
		return zero, err
	}
	err = json.Unmarshal(b, &zero)
	return zero, err
}

/* File-Write-Json-difine by:Leftyer dt:2025-11-01 */
func WriteJSON[T any](path string, obj T) error {
	b, _ := json.MarshalIndent(obj, "", "  ")
	return os.WriteFile(path, []byte(b), 0644)
}
