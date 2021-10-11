package main

import (
	"fmt"
	"log"
	"os"
	"go/printer"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Validate the go file name is passed as an argument
	if len(os.Args) > 2 {
		mainFilePath := workingDir + "\\..\\..\\test-files\\" + os.Args[2]
		// Validate the go file exists
		if _, err := os.Stat(mainFilePath); err == nil {

			destFilePath := workingDir + "\\..\\..\\test-files\\" + os.Args[3]

			// Generate the AST
			fset := token.NewFileSet()
			astFile, err := parser.ParseFile(fset, mainFilePath, nil, parser.ParseComments)
			if err != nil {
				log.Fatal(err)
			}

			// Add a print statement to each function
			ast.Inspect(astFile, func(n ast.Node) bool {
				functionDeclaration, ok := n.(*ast.FuncDecl)
				if ok {
					name := functionDeclaration.Name
	
				    printStmt := ast.ExprStmt{
					 	X: &ast.CallExpr{
					 		Fun:  &ast.SelectorExpr{
							  	X: &ast.Ident{Name: "fmt"},
					 			Sel: &ast.Ident{Name: "Println"}},
								Args: []ast.Expr {
									&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", name)}}}}

					functionDeclaration.Body.List = append([]ast.Stmt{&printStmt}, functionDeclaration.Body.List...)
				}
				return true
			})

			// Write the altered AST to the desired file path
			CreateAlteredFile(fset, astFile, destFilePath)

		} else if os.IsNotExist(err) {
			log.Fatal(os.Args[2] + " does not exist. Please make sure to place the go file in test-files folder.")
		}
	} else {
		log.Fatal("Please send the program name as an argument to do the dynamic analysis.")
	}
}

func CreateAlteredFile(fset *token.FileSet, file *ast.File, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)	
	}
	defer f.Close()

	if err := printer.Fprint(f, fset, file); err != nil {
		log.Fatal(err) 
	}
}
