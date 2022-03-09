package pkg

type Sum struct {
	Num1 int
	Num2 int
}

func (s *Sum) Add() int {
	return s.Num1 + s.Num2
}
