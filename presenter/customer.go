package presenter

import (
	"fmt"
	"log"
	"newretail-homework/view"
	"strings"
	"time"

	"gorm.io/gorm"
)


func FindTargetCustomers(db *gorm.DB, sinceDays int, minAmount float64) ([]view.CustomerResponse, map[uint]float64, error) {
    var results []view.CustomerResponse
    spendMap := make(map[uint]float64)
    since := time.Now().AddDate(0, 0, -sinceDays)

    rows, err := db.Raw(`
         SELECT c.name, c.area_code, c.phone, 
            p.customer_id, SUM(p.amount) AS total
            FROM purchase p
            INNER JOIN customer c ON c.id = p.customer_id
            WHERE p.purchased_at >= ?
            GROUP BY p.customer_id, c.name, c.area_code, c.phone
            HAVING SUM(p.amount) > ?
    `, since, minAmount).Rows()

    if err != nil {
		log.Fatal("no exist client")
		return nil, nil, err
	}
	defer rows.Close()

    for rows.Next() {
		var r view.CustomerResponse
		if err := rows.Scan(&r.Name, &r.AreaCode, &r.Phone, &r.CustomerId, &r.Total); err != nil {
			return nil, nil, err
		}
		results = append(results, r)
	}
    return results, spendMap, nil
}




func SendMarketingSMS(customers []view.CustomerResponse, template string) []string {
    var messages []string
	for _, c := range customers {
		msg := strings.ReplaceAll(template, "{{name}}", c.Name)
		msg = strings.ReplaceAll(msg, "{{amount}}", fmt.Sprintf("%.2f", c.Total))
		log.Printf("Send SMS to %s: %s", c.Phone, msg)
        message := fmt.Sprintf("Send SMS to %s: %s", c.Phone, msg)
		messages = append(messages, message)
	}
    return messages
}




