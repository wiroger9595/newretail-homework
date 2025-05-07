package models

import "time"

type Customer struct {
	ID            uint      
	Name          string    
	Email         string    
	Phone         string    
	CreatedAt     time.Time
	LastActiveAt  time.Time
}