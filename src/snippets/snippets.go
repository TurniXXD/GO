package functions

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/jasonlvhit/gocron"
)

func SaveOnFourHoursPeriod() {
	for time.Now().Format("04") != "00" {
		time.Sleep(1 * time.Minute)
	}
	gocron.Every(4).Hours().Do(fmt.Println(time.Now()))
	<-gocron.Start()
}

func OutputHTML(w http.ResponseWriter, r *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, r, file.Name(), fi.ModTime(), file)
}

func IsValidEmail(email string) bool {
	var re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,}[a-zA-Z0-9])?)*(?:\\.[a-zA-Z]{2,})$")
	return re.MatchString(email)
}

func IsValidDomain(domain string) bool {
	pattern := `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(pattern, domain)
	return match
}

func IsValidHttpsUrl(input string) (err error) {
	pattern := `^https?:\/\/(www\.)?([a-zA-Z0-9-_\.]+\.[a-zA-Z]{2,})(:[0-9]+)?([\/\?\#].*)?$`
	match, err := regexp.MatchString(pattern, input)
	u, parseErr := url.Parse(input)

	if err != nil || !match || parseErr != nil {
		return fmt.Errorf("IsValidHttpsUrl", fmt.Sprintf("error: %v is not valid url", input))
	}

	if u.Scheme != "https" {
		return fmt.Errorf("IsValidHttpsUrl", "only https URLs are allowed")
	}

	return nil
}

type SecureString struct {
	Value  string
	Length int
}

type SecureStringGenEnum struct {
	ClientId   string
	Hash       string
	Secret     string
	LongSecret string
}

var (
	SecureStringType = SecureStringGenEnum{
		Hash:       "HASH",
		Secret:     "SECRET",
		LongSecret: "LONG_SECRET",
	}
)

func GenerateSecureString(stringType string) *string {
	stringLength := 32

	if stringType == SecureStringType.Hash {
		stringLength = 64
	}

	if stringType == SecureStringType.Secret {
		stringLength = 128
	}

	if stringType == SecureStringType.LongSecret {
		stringLength = 256
	}

	randomBytes := make([]byte, stringLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Errorf("GenerateSecureString", err.Error())
	}

	if stringType == SecureStringType.Hash ||
		stringType == SecureStringType.ClientId {
		base64Hash := base64.URLEncoding.EncodeToString(randomBytes)
		return &base64Hash
	}

	hash := sha512.Sum512(randomBytes)
	hashString := hex.EncodeToString(hash[:])
	return &hashString
}

func ExtractDomain(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	domain := u.Hostname()
	parts := strings.Split(domain, ".")
	if len(parts) > 2 {
		domain = strings.Join(parts[len(parts)-2:], ".")
	}

	return domain
}

func CreateDirIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func squashFiles(src string) (allFilesContent string) {
	// Walk through all files in the source directory and append their content to the allFilesContent string
	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileContent, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			allFilesContent += string(fileContent)
		}

		return nil
	})

	return allFilesContent
}

func CreateFile(out, file string, content []byte) error {
	return os.WriteFile(out+"/"+file, content, 0755)
}

func CreateSquashedFile(src, out, file string) {
	squashedFileContent := squashFiles(src)
	err := CreateFile(out, file, []byte(squashedFileContent))
	if err != nil {
		fmt.Errorf("CreateSquashedFile", err.Error())
	}
}

func FindDifferentItems(compareToArray *[]string, itemsArray *[]string) *[]string {
	diff := []string{}
	map1 := make(map[string]bool)
	map2 := make(map[string]bool)

	for _, item := range *compareToArray {
		map1[item] = true
	}

	for _, item := range *itemsArray {
		map2[item] = true
	}

	for item := range map1 {
		if _, ok := map2[item]; !ok {
			diff = append(diff, item)
		}
	}

	return &diff
}

func ConcateStrings() (result string) {
	// This is not optimized, because each operation causes a new memory allocation
	for i := 0; i < 1000; i++ {
		result += "test"
	}

	// Instead create list of strings and then joint all of them at once
	buffer := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		buffer = append(buffer, "test")
	}

	result = strings.Join(buffer, "")

	// Or use builder from standard library https://github.com/golang/go/blob/master/src/strings/builder.go
	// This doesn't expose the underlying byte slice it uses as buffer, bytes.Buffer has its own method which returns its internal buffer
	// Builder also has checks that the builder values are not copied
	var standardStringBuffer strings.Builder
	for i := 0; i < 1000; i++ {
		standardStringBuffer.WriteString("test")
	}

	result = standardStringBuffer.String()
	return
}
