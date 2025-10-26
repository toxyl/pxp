package language

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/toxyl/flo"
)

func loadFile(filePath string) (string, []byte, error) {
	filePath = strings.TrimSpace(filePath)
	isRemote := strings.HasPrefix(filePath, "http://") || strings.HasPrefix(filePath, "https://")

	var err error
	var data []byte
	if isRemote {
		resp, err := http.Get(filePath)
		if err != nil {
			return "", nil, fmt.Errorf("failed to download file '%s': %s", filePath, err.Error())
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", nil, fmt.Errorf("failed to download file: status %s", resp.Status)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", nil, fmt.Errorf("failed to read body: %s", err.Error())
		}

		hash := sha256.Sum256([]byte(filePath))
		tempFileName := fmt.Sprintf("%x-%s", hash, filepath.Base(filePath))
		filePath = filepath.Join(os.TempDir(), tempFileName)
		if err = flo.File(filePath).StoreBytes(data); err != nil {
			return "", nil, fmt.Errorf("could not store downloaded file: %s", err.Error())
		}
	}

	data, err = os.ReadFile(filePath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read file '%s': %s", filePath, err.Error())
	}
	return filePath, data, nil
}
