package model

import "time"



type MessageLog struct {
	ID         uint      
	CustomerID uint
	Message    string
	SentAt     time.Time
}