package env

import (
	"bufio"
	"crypt/pkg/cipher"
	"log"
	"os"
	"strings"
)

const (
	prefix string = "[ENC--"
	suffix string = "--ENC]"
)

func ReadEnv(path string) map[string]string {
	env := make(map[string]string)

	readFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() != "" && strings.Contains(fileScanner.Text(), "=") {
			l := strings.Split(fileScanner.Text(), "=")
			val := mergeStringValue(l)
			env[l[0]] = val
		}
	}

	readFile.Close()
	return env
}

func SaveEnv(envs map[string]string, path string) {
	var str []string
	for key, value := range envs {
		str = append(str, key+"="+value+"\n")
	}

	stringByte := strings.Join(str, "")
	err := os.WriteFile(path, []byte(stringByte), 0644)
	if err != nil {
		panic(err)
	}
}

func EncodeEnv(envs map[string]string) map[string]string {
	for key, value := range envs {
		valueEnc := []byte(value)
		if !strings.HasPrefix(value, prefix) && !strings.HasSuffix(value, suffix) {
			envs[key] = prefix + string(cipher.Ecnrypt([]byte("B3FFCA612CD0C3D9050A4DE3"), valueEnc)) + suffix
		}
	}

	return envs
}

func DecodeEnv(envs map[string]string) map[string]string {
	for key, value := range envs {
		if strings.HasPrefix(value, prefix) && strings.HasSuffix(value, suffix) {
			value = strings.Replace(value, prefix, "", -1)
			value = strings.Replace(value, suffix, "", -1)
			envs[key] = string(cipher.Decrypt([]byte("B3FFCA612CD0C3D9050A4DE3"), value))
		}
	}

	return envs
}

func mergeStringValue(text []string) string {
	// Remove first element
	textslice := text[1:]

	// If lk length is more than 1
	// means the value of the secret
	// contains '=' strings
	if len(textslice) > 1 {
		return strings.Join(textslice[:], "=")
	}

	return textslice[0]
}
