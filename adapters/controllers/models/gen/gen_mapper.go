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
		"ModelName":  def.Name,
		"ModelArias": def.Alias(),
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

	inputFilePaths, err := utils.ListUpFilePaths("..", "models\\.gen\\.go")
	if err != nil || len(inputFilePaths) == 0 {
		fmt.Fprintf(os.Stderr, "model file not found.\n: %s", err)
		os.Exit(1)
	}

	defs, err := utils.ParseFirstStruct(inputFilePaths[0])
	if err != nil || len(defs) == 0 {
		fmt.Fprintf(os.Stderr, "model parse faild.\n: %s", err)
		os.Exit(1)
	}

	for _, def := range defs {
		if def.Name == "CommonProps" || def.Name == "Error" || strings.HasPrefix(def.Name, "New") {
			continue
		}
		outputFilePath := "../" + utils.ToSnakeCase(def.Name) + "_mapper.gen.edt.go"
		err = createMapperFile(outputFilePath, def)
		if err != nil {
			fmt.Fprintf(os.Stderr, "code generate failed.\n: %s", err)
			os.Exit(1)
		}

	}

	fmt.Println("DB model mapper generation finished.")
}
