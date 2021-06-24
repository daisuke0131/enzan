package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func evaluateBasicLit(basic *ast.BasicLit) float64{
	switch basic.Kind {
	case token.INT:
		v, err := strconv.ParseInt(basic.Value,10,64)
		if err != nil{
			panic("int parse error")
		}
		return float64(v)

	case token.FLOAT:
		v, err := strconv.ParseFloat(basic.Value,64)
		if err != nil{
			panic("float parse error")
		}
		return v
	default:
		panic("BasicLit error")
	}
}

func evaluateBinaryExpr(expr *ast.BinaryExpr) float64{
	x := evaluateExpr(expr.X)
	y := evaluateExpr(expr.Y)

	switch expr.Op {
	case token.ADD:
		return x + y
	case token.SUB:
		return x - y
	case token.MUL:
		return x * y
	case token.QUO:
		return x / y
	case token.REM:
		return float64(int(x) % int(y))
	default:
		panic("binaryExpr error")
	}
}

func evaluateUnaryExpr(expr *ast.UnaryExpr) float64{
	x := evaluateExpr(expr.X)

	switch expr.Op {
	case token.ADD:
		return x
	case token.SUB:
		return -x
	default:
		panic("unaryExpr error")
	}
}


func evaluateExpr(expr ast.Expr) float64 {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return evaluateBasicLit(e)
	case *ast.BinaryExpr:
		return evaluateBinaryExpr(e)
	case *ast.UnaryExpr:
		return evaluateUnaryExpr(e)
	case *ast.ParenExpr:
		return evaluateExpr(e.X)
	default:
		panic("evaluateExpr error")
	}
}


func main(){
	expr , err := parser.ParseExpr("1+(-2)+3*(3+1)+1")
	if err != nil{
		panic("error")
	}
	v := evaluateExpr(expr)
	println(v)
}
