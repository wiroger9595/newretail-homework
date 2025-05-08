package model

import "time"

type Customer struct {
	ID            uint      
	Name          string    
	AreaCode	  string
	Email         string    
	Phone         string    
	CreatedAt     time.Time
	LastActiveAt  time.Time
}