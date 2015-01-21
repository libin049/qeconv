package qeconv

import (
	"errors"
	"strings"
	"strconv"
)

type smt2Conv struct {
	err error
}

func (m *smt2Conv) All(f Formula, co *cnv_out) {
	prefixm(f, m, "(forall ", " ", ")", co)
}

func (m *smt2Conv) Ex(f Formula, co *cnv_out) {
	prefixm(f, m, "(exists ", " ", ")", co)
}

func (m *smt2Conv) And(f Formula, co *cnv_out) {
	prefixm(f, m, "(and ", " ", ")", co)
}

func (m *smt2Conv) Or(f Formula, co *cnv_out) {
	prefixm(f, m, "(or ", " ", ")", co)
}

func (m *smt2Conv) Not(f Formula, co *cnv_out) {
	prefixm(f, m, "(not ", " ", ")", co)
}

func (m *smt2Conv) Impl(f Formula, co *cnv_out) {
	prefixm(f, m, "(=> ", " ", ")", co)
}

func (m *smt2Conv) Equiv(f Formula, co *cnv_out) {
	co.append("(and (=> ")
	conv2(f.args[0], m, co)
	co.append(" ")
	conv2(f.args[1], m, co)
	co.append(") (=> ")
	conv2(f.args[1], m, co)
	co.append(" ")
	conv2(f.args[0], m, co)
	co.append("))")
}

func (m *smt2Conv) Abs(f Formula, co *cnv_out) {
	m.err = errors.New("unsupport the abs function")
}

func (m *smt2Conv) Leop(f Formula, co *cnv_out) {
	prefixm(f, m, "(<= ", " ", ")", co)
}

func (m *smt2Conv) Ltop(f Formula, co *cnv_out) {
	prefixm(f, m, "(< ", " ", ")", co)
}

func (m *smt2Conv) Eqop(f Formula, co *cnv_out) {
	prefixm(f, m, "(= ", " ", ")", co)
}

func (m *smt2Conv) Neop(f Formula, co *cnv_out) {
	prefixm(f, m, "(not (= ", " ", "))", co)
}

func (m *smt2Conv) List(f Formula, co *cnv_out) {
	prefixm(f, m, "(", " ", ")", co)
}

func (m *smt2Conv) Plus(fml Formula, co *cnv_out) {
	prefixm(fml, m, "(+ ", " ", ")", co)
}
func (m *smt2Conv) Minus(fml Formula, co *cnv_out) {
	prefixm(fml, m, "(- ", " ", ")", co)
}
func (m *smt2Conv) Mult(fml Formula, co *cnv_out) {
	prefixm(fml, m, "(* ", " ", ")", co)
}
func (m *smt2Conv) Div(fml Formula, co *cnv_out) {
	prefixm(fml, m, "(/ ", " ", ")", co)
}
func (m *smt2Conv) Pow(fml Formula, co *cnv_out) {
	exp := fml.args[1]
	if exp.cmd != NUMBER {
		m.err = errors.New("unsupport rational exponential")
	}
	co.append("(*")
	n,_ := strconv.Atoi(exp.str)
	for i := 0; i < n; i++ {
		co.append(" ")
		conv2(fml.args[0], m, co)
	}
	co.append(")")
}

func (m *smt2Conv) Ftrue() string {
	return "true"
}

func (m *smt2Conv) Ffalse() string {
	return "false"
}
func (m *smt2Conv) Comment(str string) string {
	return ";" + str
}

func ToSmt2(str string) (string, error) {
	stack = new(Stack)
	l := new(SynLex)
	l.Init(strings.NewReader(str))
	yyParse(l)
	fml := tofml(stack)
	qc := new(smt2Conv)
	qc.err = nil
	qstr := conv(fml, qc, l.comment)
	return qstr, qc.err
}