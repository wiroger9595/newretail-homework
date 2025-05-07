package models

type CustomerTag struct {
	CustomerID uint   // 外鍵對應 customers
	Tag        string 
}