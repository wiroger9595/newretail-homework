

實作的客戶篩選與行銷簡訊和發放優惠券 API 專案。


資料庫啟動(Postgres, Redis)

- docker-compose up


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

#### 🔸 Request Body

```json
{
  "since_days": 30,
  "amount": 1000,
  "template": "Hi {{name}}, 您在本月消費達到 {{amount}} 元，感謝您的支持！"
}

---

### 2. 發送優惠券

根據使用者提供的：

- `userId`：客戶id
- `couponId`：優惠券id

發放優惠券給客戶，要確保優惠券不會超發。

假設優惠券沒有，過期要跳回錯誤訊息。

---



