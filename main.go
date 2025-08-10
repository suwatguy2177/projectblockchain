package main

import (
	"fmt"
	"os"
	"os/exec"
)

type AESHandler struct{}

// ฟังก์ชันเข้ารหัส
func (a AESHandler) EncryptModel(inputPath, outputPath, password string) error {
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("input file not found: %s", inputPath)
	}

	fmt.Println("Using input file path :", inputPath)
	fmt.Println("Using output file path:", outputPath)

	cmd := exec.Command("openssl", "enc", "-aes-256-cbc",
		"-salt",
		"-in", inputPath,
		"-out", outputPath,
		"-pass", fmt.Sprintf("pass:%s", password),
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("openssl encrypt error: %v\noutput: %s", err, string(out))
	}

	return nil
}

// ฟังก์ชันถอดรหัส
func (a AESHandler) DecryptModel(inputPath, outputPath, password string) error {
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("input file not found: %s", inputPath)
	}

	fmt.Println("Using input file path :", inputPath)
	fmt.Println("Using output file path:", outputPath)

	cmd := exec.Command("openssl", "enc", "-d", "-aes-256-cbc",
		"-in", inputPath,
		"-out", outputPath,
		"-pass", fmt.Sprintf("pass:%s", password),
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("openssl decrypt error: %v\noutput: %s", err, string(out))
	}

	return nil
}

func main() {
	handler := AESHandler{}
	password := "mypassword123" // กำหนดรหัสผ่าน

	// 🔹 1. เข้ารหัสไฟล์
	inputPlain := "C:/Users/DEll/Documents/meshmixer/models/bunnyr.fbx"
	encryptedFile := "C:/Users/DEll/Documents/meshmixer/models/bunnyr_encrypted.fbx"

	fmt.Println("🔐 Starting encryption...")
	if err := handler.EncryptModel(inputPlain, encryptedFile, password); err != nil {
		fmt.Println("❌ Encryption error:", err)
		return
	}
	fmt.Println("✅ Encryption complete:", encryptedFile)

	// 🔹 2. ถอดรหัสไฟล์
	decryptedFile := "C:/Users/DEll/Documents/meshmixer/models/bunnyr_dec.fbx"

	fmt.Println("🔓 Starting decryption...")
	if err := handler.DecryptModel(encryptedFile, decryptedFile, password); err != nil {
		fmt.Println("❌ Decryption error:", err)
		return
	}
	fmt.Println("✅ Decryption complete:", decryptedFile)
}
