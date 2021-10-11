package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func functionPathDfs(calledFunction *ast.FuncDecl, currentPath string){
	currentPath += "->" + calledFunction.Name.Name
	noCallees := true 

	for _, body := range calledFunction.Body.List{
		if exprStmt, ok := body.(*ast.ExprStmt); ok {
			if call, ok := exprStmt.X.(*ast.CallExpr); ok {
				if ident, ok := call.Fun.(*ast.Ident); ok {
					if funct, ok := ident.Obj.Decl.(*ast.FuncDecl); ok {
						noCallees = false
						functionPathDfs(funct, currentPath)}
				}
			}
		}
	}

	if noCallees {
		fmt.Println(currentPath)
	}
}

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Validate the go file name is passed as an argument
	if len(os.Args) > 2 {
		filePath := workingDir + "\\..\\..\\test-files\\" + os.Args[2]
		// Validate the go file exists
		if _, err := os.Stat(filePath); err == nil {
			
			// Compute call graph
			// Generate the AST
			fset := token.NewFileSet()
			fileAst, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
			if err != nil {
				log.Fatal(err)
			}

			// Get the main function to start traversal
			var mainFunction *ast.FuncDecl 
			for _, d := range fileAst.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					if fn.Name.Name == "main" {
						mainFunction = fn
					}
				}
			}

			// DFS to get the call graph
			for _, body := range mainFunction.Body.List {
				if exprStmt, ok := body.(*ast.ExprStmt); ok {
					if call, ok := exprStmt.X.(*ast.CallExpr); ok {
						if ident, ok := call.Fun.(*ast.Ident); ok {
								if funct, ok := ident.Obj.Decl.(*ast.FuncDecl); ok {
									functionPathDfs(funct, "main")}
						}
					}
				}
			}			

		} else if os.IsNotExist(err) {
			log.Fatal(os.Args[2] + " does not exist. Please make sure to place the go file in test-files folder.")
		}
	} else {
		log.Fatal("Please send the program name as an argument to do the static analysis.")
	}
}
