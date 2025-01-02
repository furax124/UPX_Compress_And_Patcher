package UPX_retrieve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const upxReleasesURL = "https://api.github.com/repos/upx/upx/releases/latest"

func DownloadAndInstallUPX() error {
	fmt.Println("[*] Fetching the latest release information")
	resp, err := http.Get(upxReleasesURL)
	if err != nil {
		return fmt.Errorf("[-] failed to fetch UPX releases: %v", err)
	}
	defer resp.Body.Close()

	var release struct {
		Assets []struct {
			BrowserDownloadURL string `json:"browser_download_url"`
			Name               string `json:"name"`
		} `json:"assets"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return fmt.Errorf("[-] failed to decode UPX release information: %v", err)
	}

	if len(release.Assets) == 0 {
		return fmt.Errorf("[-] no assets found in the latest UPX release")
	}

	fmt.Println("[*] Finding the appropriate asset for Windows 64-bit")
	var downloadURL string
	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, "win64.zip") {
			downloadURL = asset.BrowserDownloadURL
			fmt.Printf("[*] Found asset: %s\n", asset.Name)
			break
		}
	}

	if downloadURL == "" {
		return fmt.Errorf("[-] no suitable asset found for Windows 64-bit")
	}

	fmt.Printf("[*] Downloading UPX from %s\n", downloadURL)
	cmd := exec.Command("curl", "-L", "-o", "upx.zip", downloadURL)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("[-] failed to download UPX: %v", err)
	}

	fmt.Println("[*] Unzipping the downloaded file")
	cmd = exec.Command("powershell", "Expand-Archive", "-Path", "upx.zip", "-DestinationPath", ".")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("[-] failed to unzip UPX: %v", err)
	}

	fmt.Println("[*] Finding the extracted directory name dynamically")
	files, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("[-] failed to read current directory: %v", err)
	}

	var extractedDir string
	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "upx-") && strings.HasSuffix(file.Name(), "-win64") {
			extractedDir = file.Name()
			fmt.Printf("[*] Found extracted directory: %s\n", extractedDir)
			break
		}
	}

	if extractedDir == "" {
		return fmt.Errorf("[-] failed to find extracted UPX directory")
	}

	fmt.Println("[*] Moving upx.exe to the current directory")
	upxExePath := filepath.Join(".", extractedDir, "upx.exe")
	err = os.Rename(upxExePath, "./upx.exe")
	if err != nil {
		return fmt.Errorf("[-] failed to move upx.exe: %v", err)
	}

	fmt.Println("[*] Cleaning up the downloaded zip file")
	err = os.Remove("upx.zip")
	if err != nil {
		return fmt.Errorf("[-] failed to remove upx.zip: %v", err)
	}

	fmt.Println("[*] Cleaning up the extracted directory")
	err = os.RemoveAll(extractedDir)
	if err != nil {
		return fmt.Errorf("[-] failed to remove extracted directory: %v", err)
	}

	fmt.Println("[+] UPX installation completed successfully")
	return nil
}

func CompressWithUPX(filePath string) error {
	fmt.Printf("[*] Compressing file %s with UPX\n", filePath)
	cmd := exec.Command("./upx.exe", filePath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("[-] failed to compress file with UPX: %v", err)
	}

	fmt.Println("[+] File compressed successfully")
	return nil
}
