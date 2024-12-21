package parser

import (
	"interpreter-book/ast"
	"interpreter-book/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 12345;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram returned nil program")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral got %s, want let", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not a *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value got %s, want %s", letStmt.Name.Value, name)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral got %s, want %s", letStmt.Name.TokenLiteral(), name)
		return false
	}

	return true
}
