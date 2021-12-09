package main

import (
    "flag"
    "proj/coordinator"
)

var (
    lex bool
    ast bool
    iloc bool
)

func init() {
    flag.BoolVar(&lex, "lex", false, "only lex")
    flag.BoolVar(&ast, "ast", false, "only ast")
    flag.BoolVar(&iloc, "iloc", false, "only iloc")
}

func main() {
    flag.Parse()
    filePath := flag.Args()[0]
    if lex {
        coordinator.DoLex(filePath)
    } else if ast {
        coordinator.ConstructAST(filePath, true)
    } else if iloc {
        coordinator.GenerateILoc(filePath)
    }
} 
