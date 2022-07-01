package api

import (
	"go-project-structure/internal/ports"
)

type Adapter struct {
	arith ports.ArithmaticPort
	db    ports.DbProt
}

func NewAdapter(arith ports.ArithmaticPort, db ports.DbProt) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (ad Adapter) GetAddition(a, b int32) (int32, error) {
	ans, err := ad.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = ad.db.AddToHistory(ans, "addition")
	if err != nil {
		return 0, nil
	}

	return ans, nil
}

func (ad Adapter) GetSubtraction(a, b int32) (int32, error) {
	ans, err := ad.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = ad.db.AddToHistory(ans, "subtraction")
	if err != nil {
		return 0, nil
	}

	return ans, nil
}

func (ad Adapter) GetMultiplication(a, b int32) (int32, error) {
	ans, err := ad.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = ad.db.AddToHistory(ans, "multiplication")
	if err != nil {
		return 0, nil
	}

	return ans, nil
}

func (ad Adapter) GetDivision(a, b int32) (int32, error) {
	ans, err := ad.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = ad.db.AddToHistory(ans, "divivsion")
	if err != nil {
		return 0, nil
	}

	return ans, nil
}
