package ir

import (

)

type Add struct {
	CommandRRR
}

func NewAdd(target int, reg2 int, reg3 int) *Add {
	res := &Add{}
	res.CommandRRR = *NewCommandRRR("add", target, reg2, reg3)
	return res
}

type AddC struct {
	CommandRRC
}

func NewAddC(target int, reg2 int, constant int) *AddC {
	res := &AddC{}
	res.CommandRRC = *NewCommandRRC("add", target, reg2, constant)
	return res
}

type Sub struct {
	CommandRRR
}

func NewSub(target int, reg2 int, reg3 int) *Sub {
	res := &Sub{}
	res.CommandRRR = *NewCommandRRR("sub", target, reg2, reg3)
	return res
}

type SubC struct {
	CommandRRC
}

func NewSubC(target int, reg2 int, constant int) *SubC {
	res := &SubC{}
	res.CommandRRC = *NewCommandRRC("sub", target, reg2, constant)
	return res
}

type Mul struct {
	CommandRRR
}

func NewMul(target int, reg2 int, reg3 int) *Mul {
	res := &Mul{}
	res.CommandRRR = *NewCommandRRR("mul", target, reg2, reg3)
	return res
}

type MulC struct {
	CommandRRC
}

func NewMulC(target int, reg2 int, constant int) *MulC {
	res := &MulC{}
	res.CommandRRC = *NewCommandRRC("mul", target, reg2, constant)
	return res
}

type Div struct {
	CommandRRR
}

func NewDiv(target int, reg2 int, reg3 int) *Div {
	res := &Div{}
	res.CommandRRR = *NewCommandRRR("div", target, reg2, reg3)
	return res
}

type DivC struct {
	CommandRRC
}

func NewDivC(target int, reg2 int, constant int) *DivC {
	res := &DivC{}
	res.CommandRRC = *NewCommandRRC("div", target, reg2, constant)
	return res
}

type And struct {
	CommandRRR
}

func NewAnd(target int, reg2 int, reg3 int) *And {
	res := &And{}
	res.CommandRRR = *NewCommandRRR("and", target, reg2, reg3)
	return res
}

type AndC struct {
	CommandRRC
}

func NewAndC(target int, reg2 int, constant int) *AndC {
	res := &AndC{}
	res.CommandRRC = *NewCommandRRC("and", target, reg2, constant)
	return res
}

type Or struct {
	CommandRRR
}

func NewOr(target int, reg2 int, reg3 int) *Or {
	res := &Or{}
	res.CommandRRR = *NewCommandRRR("or", target, reg2, reg3)
	return res
}

type OrC struct {
	CommandRRC
}

func NewOrC(target int, reg2 int, constant int) *OrC {
	res := &OrC{}
	res.CommandRRC = *NewCommandRRC("or", target, reg2, constant)
	return res
}

type Cmp struct {
	CommandRR
}

func NewCmp(reg1 int, reg2 int) *Cmp {
	res := &Cmp{}
	res.CommandRR = *NewCommandRR("cmp", reg1, reg2)
	return res
}

type CmpC struct {
	CommandRC
}

func NewCmpC(reg1 int, constant int) *CmpC {
	res := &CmpC{}
	res.CommandRC = *NewCommandRC("cmp", reg1, constant)
	return res
}

type MoveReg struct{
	CommandRR
}

func NewMoveReg(target int, reg2 int) *MoveReg {
	res := &MoveReg{}
	res.CommandRR = *NewCommandRR("mov", target, reg2)
	return res
}

type Move struct{
	CommandRC
}

func NewMove(target int, operand int) *Move {
	res := &Move{}
	res.CommandRC = *NewCommandRC("mov", target, operand)
	return res
}

type MoveEq struct{
	CommandRC
}

func NewMoveEq(target int, operand int) *MoveEq {
	res := &MoveEq{}
	res.CommandRC = *NewCommandRC("moveq", target, operand)
	return res
}

type MoveNe struct{
	CommandRC
}

func NewMoveNe(target int, operand int) *MoveNe {
	res := &MoveNe{}
	res.CommandRC = *NewCommandRC("movne", target, operand)
	return res
}

type MoveGt struct{
	CommandRC
}

func NewMoveGt(target int, operand int) *MoveGt {
	res := &MoveGt{}
	res.CommandRC = *NewCommandRC("movgt", target, operand)
	return res
}

type MoveGe struct{
	CommandRC
}

func NewMoveGe(target int, operand int) *MoveGe {
	res := &MoveGe{}
	res.CommandRC = *NewCommandRC("movge", target, operand)
	return res
}

type MoveLt struct{
	CommandRC
}

func NewMoveLt(target int, operand int) *MoveLt {
	res := &MoveLt{}
	res.CommandRC = *NewCommandRC("movlt", target, operand)
	return res
}

type MoveLe struct{
	CommandRC
}

func NewMoveLe(target int, operand int) *MoveLe {
	res := &MoveLe{}
	res.CommandRC = *NewCommandRC("movle", target, operand)
	return res
}

type Not struct{
	CommandRR
}

func NewNot(target int, reg2 int) *Not {
	res := &Not{}
	res.CommandRR = *NewCommandRR("not", target, reg2)
	return res
}

type NotC struct{
	CommandRC
}

func NewNotC(target int, constant int) *NotC {
	res := &NotC{}
	res.CommandRC = *NewCommandRC("not", target, constant)
	return res
}

type Read struct {
	CommandR
}

func NewRead(target int) *Read {
	res := &Read{}
	res.CommandR = *NewCommandR("read", target)
	return res
}

type Print struct {
	CommandR
}

func NewPrint(target int) *Print {
	res := &Print{}
	res.CommandR = *NewCommandR("print", target)
	return res
}

type Println struct {
	CommandR
}

func NewPrintln(target int) *Println {
	res := &Println{}
	res.CommandR = *NewCommandR("println", target)
	return res
}

type Return struct {
	CommandR
}

func NewReturn(reg int) *Return {
	res := &Return{}
	res.CommandR = *NewCommandR("ret", reg)
	return res
}