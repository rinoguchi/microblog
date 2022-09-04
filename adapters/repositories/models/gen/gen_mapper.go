//go:generate go run .
//go:generate gofmt -w ../

package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/rinoguchi/microblog/utils"
)

func createMapperFile(outputFilePath string, def utils.StructDef) error {
	shouldSkip, err := utils.ShouldSkip(outputFilePath)
	if err != nil {
		return err
	}
	if shouldSkip {
		fmt.Println("generate " + outputFilePath + " is skipped.")
		return nil
	}

	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.Must(template.ParseFiles("./mapper.tmpl"))
	data := map[string]interface{}{
		"ModelName":  def.Name[2:],
		"ModelArias": strings.ToLower(def.Name[2:][0:1]) + def.Name[2:][1:],
		"Fields":     def.Fields,
	}
	if err := t.Execute(file, data); err != nil {
		return err
	}
	fmt.Println(outputFilePath + " is generated.")
	return nil
}

func main() {
	fmt.Println("DB model mapper generation started.")

	inputFilePaths, err := utils.ListUpFilePaths("..", "models.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "model file not found.\n: %s", err)
		os.Exit(1)
	}

	for _, inputFilePath := range inputFilePaths {
		list, err := utils.ParseFirstStruct(inputFilePath)
		if err != nil || len(list) == 0 {
			fmt.Fprintf(os.Stderr, "model parse faild.\n: %s", err)
			os.Exit(1)
		}

		outputFilePath := "../" + strings.Split(strings.Split(inputFilePath, "/")[1], ".")[0] + "_mapper.gen.edt.go"
		err = createMapperFile(outputFilePath, list[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "code generate failed.\n: %s", err)
			os.Exit(1)
		}
	}

	fmt.Println("DB model mapper generation finished.")
}
