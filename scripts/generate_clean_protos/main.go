package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	if err := generateCleanProto("./proto/", "./gen/clean_proto/"); err != nil {
		panic(err)
	}
}

// generateCleanProto generates clean proto files for debugging purposes (via bloomrpc for example)
func generateCleanProto(in string, out string) error {
	re := regexp.MustCompile(`\[((.|\n)*?)\];`)                                         // удаляем параметры полей
	re2 := regexp.MustCompile(`option \(((.|\n)*?)};\n\n`)                              // удаляем глобальные опиции
	re3 := regexp.MustCompile(`option \(((.|\n)*?)};\n`)                                // удаляем локальные опции
	re4 := regexp.MustCompile(`import(( public)?) "github\.com(.*?).proto";`)           // удаляем лишние импорты
	re5 := regexp.MustCompile(`import(( public)?) "google/api/(.*?).proto";`)           // удаляем лишние импорты
	re6 := regexp.MustCompile(`import(( public)?) "protoc-gen-openapiv2/(.*?).proto";`) // удаляем лишние импорты
	re7 := regexp.MustCompile(`option \((.*?) = ((.|\n)*?);\n`)                         // удаляем лишние импорты

	files, err := ioutil.ReadDir(in)
	if err != nil {
		return err
	}
	for _, f := range files {
		if !f.IsDir() {
			filename := in + f.Name()
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}
			str := string(data)
			str = re.ReplaceAllString(str, ";")
			str = re2.ReplaceAllString(str, "")
			str = re3.ReplaceAllString(str, "")
			str = re4.ReplaceAllString(str, "")
			str = re5.ReplaceAllString(str, "")
			str = re6.ReplaceAllString(str, "")
			str = re7.ReplaceAllString(str, "")

			for strings.Contains(str, "\n\n\n") {
				str = strings.ReplaceAll(str, "\n\n\n", "\n\n")
			}
			str = strings.ReplaceAll(str, "                    }", "\t}")
			str = strings.ReplaceAll(str, "{\n        uint32", "{\n\tuint32")
			str = strings.ReplaceAll(str, "{\n\t\t", "{\n\t")
			str = strings.ReplaceAll(str, " ;", ";")

			if err := ioutil.WriteFile(out+f.Name(), []byte(str), 0600); err != nil {
				panic(err)
			}
		}
	}
	return nil
}
