

實作的客戶篩選與行銷簡訊和發放優惠券 API 專案。


資料庫啟動(Postgres, Redis)
```
- docker-compose up
```
專案啟動
```
- go run main.go
```

## 🚀 功能簡介

### 1. 客戶篩選查詢和發送行銷簡訊發送（模擬）

根據使用者提供的：

- `since_days`：過去幾天內的消費紀錄
- `amount`：消費金額門檻

來查找消費總金額大於指定數值的客戶。

針對符合條件的客戶，根據模板發送行銷簡訊（支援變數 `{{name}}`, `{{amount}}`）。

---

## 📡 API 說明

### `POST /customer/send`

查詢並發送簡訊給目標客戶。

#### 🔸 Request 

```json
{
    "SinceDays": 20,
    "Amount": 100 
}
```

#### 🔸 Response

```json
{
  "sms_logs": [
        "Send SMS to 0955667788: 親愛的 Eric Tsai，您在最近的消費金額為 299.00 元，感謝您的支持！",
        "Send SMS to 0933444555: 親愛的 Charlie Lin，您在最近的消費金額為 600.00 元，感謝您的支持！",
        "Send SMS to 0912345678: 親愛的 Alice Chen，您在最近的消費金額為 440.00 元，感謝您的支持！"
    ]
}

```
---

### 2. 發送優惠券

根據使用者提供的：

- `userId`：客戶id
- `couponId`：優惠券id

發放優惠券給客戶，要確保優惠券不會超發。

假設優惠券沒有，過期要跳回錯誤訊息。

#### 🔸 Request 


```json
{
    "userId": 3,
	"couponId": 3,
    "userLevel": "VIP"
}
```

#### 🔸 Response

```json
{
    "message": "get free coupon success",
    "response": [
        {
            "UserId": 3,
            "CouponId": 3,
            "Quantity": 0,
        }
    ]
}
---



