# API Documentation

Complete API reference for the Profile API.

## Base URL

```
https://your-api-url.com
```

For local development:
```
http://localhost:8080
```

---

## Endpoints

### 1. Get Profile (`/me`)

Returns user profile information with a dynamic cat fact.

#### Request

```http
GET /me HTTP/1.1
Host: your-api-url.com
```

#### cURL Example

```bash
curl -X GET https://your-api-url.com/me
```

#### Response

**Status Code:** `200 OK`

**Headers:**
```
Content-Type: application/json
Access-Control-Allow-Origin: *
```

**Body:**
```json
{
  "status": "success",
  "user": {
    "email": "john.doe@example.com",
    "name": "John Doe",
    "stack": "Go/Native HTTP"
  },
  "timestamp": "2025-10-16T14:23:45.123456789Z",
  "fact": "Cats have over 20 vocalizations, including the purr, meow, and hiss."
}
```

#### Response Fields

| Field | Type | Description |
|-------|------|-------------|
| `status` | string | Always "success" for successful requests |
| `user` | object | User profile information |
| `user.email` | string | User's email address |
| `user.name` | string | User's full name |
| `user.stack` | string | Backend technology stack |
| `timestamp` | string | Current UTC time in ISO 8601 format |
| `fact` | string | Random cat fact from Cat Facts API |

#### Error Responses

**Invalid Method (405)**
```json
{
  "status": "error",
  "message": "Method not allowed. Use GET."
}
```

**Cat Facts API Unavailable (200)**
```json
{
  "status": "success",
  "user": {
    "email": "john.doe@example.com",
    "name": "John Doe",
    "stack": "Go/Native HTTP"
  },
  "timestamp": "2025-10-16T14:23:45.123456789Z",
  "fact": "Unable to fetch cat fact at this moment. Please try again later."
}
```

---

### 2. Health Check (`/health`)

Returns the health status of the API.

#### Request

```http
GET /health HTTP/1.1
Host: your-api-url.com
```

#### cURL Example

```bash
curl -X GET https://your-api-url.com/health
```

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "status": "healthy",
  "time": "2025-10-16T14:23:45Z"
}
```

---

### 3. Root (`/`)

Returns basic API information.

#### Request

```http
GET / HTTP/1.1
Host: your-api-url.com
```

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "message": "Profile API - Use GET /me to retrieve profile information"
}
```

---

## HTTP Status Codes

| Code | Meaning | When It Occurs |
|------|---------|----------------|
| 200 | OK | Successful request |
| 404 | Not Found | Invalid endpoint |
| 405 | Method Not Allowed | Invalid HTTP method |
| 500 | Internal Server Error | Server error |

---

## Rate Limiting

Currently, no rate limiting is implemented. For production use, consider implementing rate limiting to prevent abuse.

**Recommended limits:**
- 100 requests per minute per IP
- 1000 requests per hour per IP

---

## CORS Support

The API supports Cross-Origin Resource Sharing (CORS) with the following headers:

```
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, OPTIONS
Access-Control-Allow-Headers: Content-Type
```

---

## External Dependencies

### Cat Facts API

The `/me` endpoint fetches cat facts from:

```
https://catfact.ninja/fact
```

**Timeout:** 5 seconds

**Fallback:** If the API is unavailable, a fallback message is returned instead of failing the request.

---

## Examples

### JavaScript (Fetch API)

```javascript
fetch('https://your-api-url.com/me')
  .then(response => response.json())
  .then(data => {
    console.log('Profile:', data.user);
    console.log('Cat Fact:', data.fact);
    console.log('Timestamp:', data.timestamp);
  })
  .catch(error => console.error('Error:', error));
```

### Python (requests)

```python
import requests

response = requests.get('https://your-api-url.com/me')
data = response.json()

print(f"Name: {data['user']['name']}")
print(f"Email: {data['user']['email']}")
print(f"Stack: {data['user']['stack']}")
print(f"Cat Fact: {data['fact']}")
print(f"Timestamp: {data['timestamp']}")
```

### Go

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ProfileResponse struct {
    Status    string `json:"status"`
    User      struct {
        Email string `json:"email"`
        Name  string `json:"name"`
        Stack string `json:"stack"`
    } `json:"user"`
    Timestamp string `json:"timestamp"`
    Fact      string `json:"fact"`
}

func main() {
    resp, err := http.Get("https://your-api-url.com/me")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var profile ProfileResponse
    json.NewDecoder(resp.Body).Decode(&profile)

    fmt.Printf("Name: %s\n", profile.User.Name)
    fmt.Printf("Cat Fact: %s\n", profile.Fact)
}
```

### cURL (Detailed)

```bash
# Basic request
curl https://your-api-url.com/me

# With headers
curl -i https://your-api-url.com/me

# Pretty print JSON (with jq)
curl -s https://your-api-url.com/me | jq .

# Extract specific field
curl -s https://your-api-url.com/me | jq '.fact'

# Save response to file
curl https://your-api-url.com/me -o response.json

# Test multiple requests
for i in {1..5}; do 
  curl -s https://your-api-url.com/me | jq '.timestamp'
  sleep 1
done
```

---

## Timestamp Format

All timestamps are returned in **ISO 8601 format** with nanosecond precision:

```
2025-10-16T14:23:45.123456789Z
```

**Format breakdown:**
- `2025-10-16` - Date (YYYY-MM-DD)
- `T` - Separator
- `14:23:45` - Time (HH:MM:SS)
- `.123456789` - Nanoseconds
- `Z` - UTC timezone indicator

### Parsing Timestamps

**JavaScript:**
```javascript
const timestamp = "2025-10-16T14:23:45.123456789Z";
const date = new Date(timestamp);
console.log(date.toLocaleString());
```

**Python:**
```python
from datetime import datetime

timestamp = "2025-10-16T14:23:45.123456789Z"
dt = datetime.fromisoformat(timestamp.replace('Z', '+00:00'))
print(dt.strftime('%Y-%m-%d %H:%M:%S'))
```

**Go:**
```go
import "time"

timestamp := "2025-10-16T14:23:45.123456789Z"
t, _ := time.Parse(time.RFC3339Nano, timestamp)
fmt.Println(t.Format("2006-01-02 15:04:05"))
```

---

## Testing

### Postman Collection

Import this JSON to test in Postman:

```json
{
  "info": {
    "name": "Profile API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Get Profile",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{base_url}}/me",
          "host": ["{{base_url}}"],
          "path": ["me"]
        }
      }
    },
    {
      "name": "Health Check",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{base_url}}/health",
          "host": ["{{base_url}}"],
          "path": ["health"]
        }
      }
    }
  ],
  "variable": [
    {
      "key": "base_url",
      "value": "https://your-api-url.com"
    }
  ]
}
```

---

## Support

For issues, questions, or contributions:
- **GitHub Issues:** [Your Repo URL]/issues
- **Email:** your.email@example.com

---
