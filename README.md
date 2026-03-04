# OTP Verification System

## 📱 Project Kya Hai? (What is this project?)

Yeh ek **OTP (One-Time Password) Verification System** hai jo SMS ke through OTP bhejta hai aur verify karta hai।

**Real-world example:**
- Jab aap Flipkart/Amazon mein account banate ho, phone number daalkr, unhe OTP bhejta hai
- Aap woh OTP enter karte ho
- System check karta hai ki OTP sahi hai ya nahi
- Agar sahi hai toh account verify ho jata hai

Yeh project bilkul wahi kaam karta hai! ✅

---

## 🛠️ Technology Stack (Technique kaun si use kri gyi?)

| Technology | Use |
|-----------|-----|
| **Go (Golang)** | Backend programming language |
| **Gin Framework** | Web server banane ke liye |
| **Twilio** | SMS service (OTP bhejne ke liye) |
| **Godotenv** | Environment variables (.env file se data padhne ke liye) |
| **Postman** | Testing ke liye |

---

## 🤔 Kyu Go aur Gin use kiya?

### Go use kiya kyu?
- **Fast** - Bohot tez hai
- **Simple** - Padhai aasaan hai
- **Lightweight** - Kam memory use karta hai
- **Production ready** - Big companies use karte hain (Google, Netflix)

### Gin use kiya kyu?
- Python ke Flask jaisa, simple aur easy
- Bohot fast requests handle karta hai
- Beginner-friendly

### Twilio use kiya kyu?
- SMS service karna padta hai, aur Twilio already banaya hua SMS service hai
- Apne aap SMS send karega - humein sirf API call karna padta hai

---

## 📁 Project Structure

```
go-sms-verify-yt/
├── cmd/
│   ├── main.go          (Server start hota hai yahan se)
│   └── .env            (Password aur keys likhe hote hain)
│
├── api/
│   ├── config.go        (Twilio set up)
│   ├── handler.go       (OTP send/verify logic)
│   ├── helper.go        (JSON response banane ke functions)
│   ├── route.go         (API endpoints)
│   └── service.go       (Twilio calls)
│
├── data/
│   └── model.go         (Data structure)
│
└── go.mod              (Dependencies)
```

---

## 🚀 Kaise Use Karein?

### 1️⃣ Server Start Karo
```bash
cd C:\Users\shivk\Downloads\Otp-Verification\go-sms-verify-yt\cmd
go run main.go
```

### 2️⃣ Postman mein Test Karo

**Step 1: OTP Send Karo**
- **Method:** POST
- **URL:** `http://localhost:8080/send-otp`
- **Body:**
```json
{
  "phoneNumber": "+91xxxxxxxxxx"
}
```
✅ SMS mein OTP aayega

**Step 2: OTP Verify Karo**
- **Method:** POST
- **URL:** `http://localhost:8080/verify-otp`
- **Body:**
```json
{
  "phoneNumber": "+91xxxxxxxxxx",
  "code": "123456"
}
```
✅ Agar code sahi hai toh "OTP verified successfully" aayega

**Step 3: Health Check**
- **Method:** GET
- **URL:** `http://localhost:8080/`
✅ Server chalti hai ya nahi check karo

---

## 💻 Code Explain (Code Samjhao)

### Handler (handler.go) mein kya hota hai?
```go
func (app *Config) SendSMS() gin.HandlerFunc {
    // Phone number validate karo
    // Twilio ko call karo "SMS bhej"
    // Response bhej
}
```

### Config (config.go) mein kya hota hai?
```go
func envACCOUNTSID() string {
    // .env file se Twilio Account ID read karo
}
```

### Helper (helper.go) mein kya hota hai?
```go
func (app *Config) writeJSON() {
    // Nice JSON format mein response bhej
}

func (app *Config) validateBody() {
    // Check karo ke phone number properly likha hai ya nahi
}
```

---

## 🔐 Environment Variables (.env file)

`.env` mein yeh likha hai:
```
TWILIO_ACCOUNT_SID=ACxxxxxxx          (Aapka Account ID)
TWILIO_AUTHTOKEN=xxxxx               (Secret password)
TWILIO_SERVICE_ID=VAxxxxxxx          (SMS Service ID)
```

⚠️ **Important:** Ye secret keys hain, kisi ko mat batana!

---

## 📊 API Endpoints

| Endpoint | Method | Kaam Kya Hai |
|----------|--------|-------------|
| `/` | GET | Check karo server chal raha hai |
| `/send-otp` | POST | SMS bhej OTP ke saath |
| `/verify-otp` | POST | OTP check karo aur verify karo |

---

## 🆘 Agar Error Aaye?

### ❌ "bind: Only one usage of each socket address"
= Port 8080 already use ho raha hai
= Purana server kill karo:
```powershell
Get-NetTCPConnection -LocalPort 8080 | Stop-Process -Force
```

### ❌ "Authentication Error"
= Twilio credentials galat hain
= `.env` file check karo

### ❌ "Phone number is invalid"
= Phone number format galat hai
= `+91xxxxxxxxxx` format likho (country code zaroori)

---

## ✅ Kya-Kya Seekhoge

- ✔️ Go mein backend banane ka tarika
- ✔️ REST API kaise banate hain
- ✔️ Environment variables kaise use karte hain
- ✔️ Third-party service (Twilio) kaise integrate karte hain
- ✔️ Error handling aur validation

---

## 📚 References

- [Go Official](https://golang.org)
- [Gin Framework](https://gin-gonic.com)
- [Twilio Docs](https://www.twilio.com/docs)

---

**Made with ❤️ for Learning**
