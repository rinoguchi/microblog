package utils

import (
	"bufio"
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// 構造体定義
type StructDef struct {
	Name   string
	Fields []FieldDef
}

func (sd StructDef) Alias() string {
	return strings.ToLower(sd.Name[0:1]) + sd.Name[1:]
}

// 構造体のフィールド定義
type FieldDef struct {
	Name string
	Type string
}

func (fd FieldDef) Alias() string {
	return strings.ToLower(fd.Name[0:1]) + fd.Name[1:]
}

// コード生成をスキップすべきかどうかを判定
func ShouldSkip(outputFilePath string) (bool, error) {
	// ファイルがなければスキップしない
	if _, err := os.Stat(outputFilePath); err != nil {
		return false, nil
	}

	// 1行目に"DO NOT OVERWRITE"というコメントがあれば、スキップする
	f, err := os.Open(outputFilePath)
	if err != nil {
		return false, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var shouldSkip bool
	for scanner.Scan() {
		shouldSkip = strings.Contains(scanner.Text(), "DO NOT OVERWRITE")
		break // 1行目だけチェック
	}
	return shouldSkip, nil
}

// ファイルをパースして構造体の定義を返却
func ParseFirstStruct(fpath string) ([]StructDef, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fpath, nil, 0)
	if err != nil {
		return nil, err
	}

	list := []StructDef{}
	ast.Inspect(f, func(n ast.Node) bool {
		x, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		if y, ok := x.Type.(*ast.StructType); ok {
			sdef := StructDef{}
			sdef.Name = x.Name.Name
			for _, fld := range y.Fields.List {
				if fld.Names == nil {
					continue
				}
				var typeNameBuf bytes.Buffer
				err := printer.Fprint(&typeNameBuf, fset, fld.Type)
				if err != nil {
					log.Fatalf("failed printing %s", err)
				}
				sdef.Fields = append(sdef.Fields, FieldDef{Name: fld.Names[0].Name, Type: typeNameBuf.String()})
			}
			list = append(list, sdef)
		}
		return true
	})
	return list, nil
}

// 対象のフォルダからpatternにマッチするファイルを取得する
func ListUpFilePaths(targetFolder string, regexpPattern string) ([]string, error) {
	files, err := ioutil.ReadDir(targetFolder)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	r := regexp.MustCompile(regexpPattern)
	for _, file := range files {
		if !file.IsDir() && r.MatchString(file.Name()) {
			paths = append(paths, targetFolder+"/"+file.Name())
		}
	}
	return paths, nil
}
