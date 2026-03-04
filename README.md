# OTP Verification System

A production-ready OTP (One-Time Password) verification service built with Go and Gin, leveraging Twilio's SMS gateway for secure phone number authentication.

## Tech Stack

| Component | Technology | Purpose |
|-----------|-----------|---------|
| **Runtime** | Go 1.x | Backend server |
| **Framework** | Gin Web Framework | HTTP REST API |
| **Third-party Service** | Twilio SMS API | SMS OTP delivery & verification |
| **Configuration** | godotenv | Environment variable management |

---

## Project Structure

```
go-sms-verify-yt/
├── cmd/
│   └── main.go              # Server entry point
├── api/
│   ├── config.go            # Twilio client initialization
│   ├── handler.go           # HTTP request handlers
│   ├── helper.go            # JSON response utilities & validation
│   ├── route.go             # Route definitions
│   └── service.go           # SMS send/verify business logic
├── data/
│   └── model.go             # Data structures (OTP request/response)
└── go.mod                   # Module dependencies
```

---

## API Endpoints

| Endpoint | Method | Description | Payload |
|----------|--------|-------------|---------|
| `/` | GET | Health check | — |
| `/send-otp` | POST | Initiate OTP flow | `{ "phoneNumber": "+91xxxxxxxxxx" }` |
| `/verify-otp` | POST | Validate OTP code | `{ "phoneNumber": "+91xxxxxxxxxx", "code": "123456" }` |

**Base URL:** `http://localhost:8080`

---

## Setup & Execution

### Prerequisites
- Go 1.16+
- Twilio Account (free tier available at [twilio.com](https://www.twilio.com))

### Installation

1. **Clone & Navigate:**
```bash
cd go-sms-verify-yt
```

2. **Configure Environment Variables:**
Create `.env` file in root:
```env
TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
TWILIO_AUTHTOKEN=your_auth_token_here
TWILIO_SERVICE_ID=VAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

3. **Install Dependencies:**
```bash
go mod download
```

4. **Run Server:**
```bash
cd cmd
go run main.go
```
Server runs on `http://localhost:8080`

---

## Usage Example

### Send OTP
```bash
curl -X POST http://localhost:8080/send-otp \
  -H "Content-Type: application/json" \
  -d '{"phoneNumber": "+91XXXXXXXXXX"}'
```

**Response:**
```json
{
  "success": true,
  "message": "OTP sent successfully"
}
```

### Verify OTP
```bash
curl -X POST http://localhost:8080/verify-otp \
  -H "Content-Type: application/json" \
  -d '{"phoneNumber": "+91XXXXXXXXXX", "code": "123456"}'
```

**Response:**
```json
{
  "success": true,
  "message": "OTP verified successfully"
}
```

---

## Error Handling

| Error | Cause | Resolution |
|-------|-------|-----------|
| `Port already in use` | Port 8080 occupied | Kill process: `netstat -ano \| findstr :8080` (Windows) |
| `Authentication failed` | Invalid Twilio credentials | Verify `.env` variables match Twilio dashboard |
| `Invalid phone number` | Malformed input | Use format: `+{country_code}{number}` |

---

## Implementation Details

### Handler Layer (`handler.go`)
- Receives HTTP requests
- Validates request payloads
- Orchestrates send/verify operations
- Returns JSON responses

### Service Layer (`service.go`)
- Encapsulates Twilio API calls
- Manages OTP generation & verification
- Handles error responses from Twilio

### Configuration (`config.go`)
- Loads Twilio credentials from environment
- Initializes API client
- Manages configuration state

---

## Security Considerations

⚠️ **Critical:**
- Never commit `.env` file (add to `.gitignore`)
- Store secrets in environment variables only
- Rotate Twilio credentials after exposure
- Use HTTPS in production
- Validate phone numbers server-side

---

## Key Features

✅ RESTful API design  
✅ Request validation & error handling  
✅ Twilio SMS integration  
✅ Environment-based configuration  
✅ JSON request/response format  

---

## Dependencies

See `go.mod` for complete dependency list. Key imports:
- `github.com/gin-gonic/gin` - Web framework
- `github.com/joho/godotenv` - .env file parser
- `github.com/twilio/twilio-go` - Twilio SDK

---

## References

- [Go Documentation](https://golang.org/doc)
- [Gin Framework](https://gin-gonic.com)
- [Twilio Verify API](https://www.twilio.com/docs/verify/api)
- [REST API Best Practices](https://restfulapi.net)

---

**Status:** Production-Ready | **License:** MIT
