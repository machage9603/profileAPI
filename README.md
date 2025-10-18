# Profile API Endpoint

A simple RESTful API built with Go that returns profile information along with dynamic cat facts from an external API.

## Features

- **Profile Endpoint**: Returns user profile with dynamic data
- **Cat Facts Integration**: Fetches random cat facts on every request
- **ISO 8601 Timestamps**: Dynamic UTC timestamps
- **Error Handling**: Graceful handling of external API failures
- **CORS Support**: Cross-origin resource sharing enabled
- **Health Check**: `/health` endpoint for monitoring

## Requirements

- Go 1.21 or higher
- Internet connection (for Cat Facts API)

## Installation

### Clone the Repository

```bash
# Replace <your-repo-url> with the actual repository URL, e.g.:
# git clone https://github.com/yourusername/profile-api.git
git clone <your-repo-url>
cd profile-api
```

### Install Dependencies

This project uses only Go standard library, so no external dependencies are needed:

```bash
go mod download
```

## Configuration

### Environment Variables

Create a `.env` file or set these environment variables:

```bash
USER_EMAIL=your.email@example.com
USER_NAME=Your Full Name
USER_STACK=Go/Native HTTP
PORT=8080  # Optional, defaults to 8080
```

### Setting Environment Variables

**Linux/Mac:**
```bash
export USER_EMAIL="john.doe@example.com"
export USER_NAME="John Doe"
export USER_STACK="Go/Native HTTP"
```

**Windows (PowerShell):**
```powershell
$env:USER_EMAIL="john.doe@example.com"
$env:USER_NAME="John Doe"
$env:USER_STACK="Go/Native HTTP"
```

## Running Locally

### Option 1: Direct Run

```bash
go run main.go
```

### Option 2: Build and Run

```bash
# Build the binary
go build -o profile-api

# Run the binary
./profile-api  # Linux/Mac
profile-api.exe  # Windows
```

### Option 3: With Environment Variables

```bash
USER_EMAIL="jane@example.com" USER_NAME="Jane Smith" USER_STACK="Go/Chi" go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### GET /me

Returns profile information with a random cat fact.

**Request:**
```bash
curl http://localhost:8080/me
```

**Response (200 OK):**
```json
{
  "status": "success",
  "user": {
    "email": "your.email@example.com",
    "name": "Your Full Name",
    "stack": "Go/Native HTTP"
  },
  "timestamp": "2025-10-16T12:34:56.789123456Z",
  "fact": "Cats have over 20 vocalizations, including the purr, meow, and hiss."
}
```

**Headers:**
- `Content-Type: application/json`
- `Access-Control-Allow-Origin: *`

### GET /health

Health check endpoint for monitoring.

**Request:**
```bash
curl http://localhost:8080/health
```

**Response (200 OK):**
```json
{
  "status": "healthy",
  "time": "2025-10-16T12:34:56Z"
}
```

### GET /

Root endpoint with API information.

**Response (200 OK):**
```json
{
  "message": "Profile API - Use GET /me to retrieve profile information"
}
```

## 🧪 Testing

### Manual Testing

**Test the /me endpoint:**
```bash
curl -i http://localhost:8080/me
```

**Test multiple requests (timestamp should change):**
```bash
for i in {1..3}; do curl http://localhost:8080/me | jq '.timestamp'; sleep 1; done
```

**Test error handling (invalid method):**
```bash
curl -X POST http://localhost:8080/me
```

## 🚢 Deployment

### Railway

1. Create a new project on [Railway](https://railway.app)
2. Connect your GitHub repository
3. Add environment variables in Railway dashboard
4. Railway will automatically detect the Go application and deploy

### Heroku

1. Create a `Procfile`:
```
web: ./profile-api
```

2. Deploy:
```bash
heroku create your-app-name
heroku config:set USER_EMAIL="your@email.com"
heroku config:set USER_NAME="Your Name"
heroku config:set USER_STACK="Go/Native HTTP"
git push heroku main
```

### AWS EC2

1. SSH into your EC2 instance
2. Install Go
3. Clone repository and build
4. Run with systemd or supervisor

### PXXL App / Other Platforms

Follow platform-specific Go deployment guides. Most platforms:
- Auto-detect Go applications
- Use `PORT` environment variable
- Support environment variable configuration

## Project Structure

```
profile-api/
├── main.go          # Main application code
├── go.mod           # Go module definition
├── README.md        # This file
├── .gitignore       # Git ignore rules
└── .env.example     # Example environment variables
```

## 🔧 Environment Variables Reference

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `USER_EMAIL` | No | `your.email@example.com` | Your email address |
| `USER_NAME` | No | `Your Full Name` | Your full name |
| `USER_STACK` | No | `Go/Native HTTP` | Your backend stack |
| `PORT` | No | `8080` | Server port |

## Error Handling

The API handles various error scenarios:

- **Cat Facts API Down**: Returns fallback message
- **Network Timeout**: 5-second timeout for external API
- **Invalid HTTP Method**: Returns 405 Method Not Allowed
- **Invalid Route**: Returns 404 Not Found

## Acceptance Criteria Checklist

- ✅ Working GET `/me` endpoint accessible with 200 OK
- ✅ Response structure follows defined JSON schema
- ✅ All required fields present (status, user, timestamp, fact)
- ✅ User object contains email, name, and stack
- ✅ Timestamp in ISO 8601 format
- ✅ Timestamp updates dynamically
- ✅ Cat fact fetched from API
- ✅ New cat fact on every request
- ✅ Content-Type header is application/json
- ✅ Well-structured code following Go best practices

## Contributing

Feel free to submit issues and enhancement requests!

## License

MIT License - feel free to use this project for learning and development.

## Author

**Your Name**
- Email: mikemachage@gmail.com
- GitHub: [@machage9603](https://github.com/machage9603)
- LinkedIn: [Mike Machage](https://linkedin.com/in/mikemachage)

## Acknowledgments

- [Cat Facts API](https://catfact.ninja/) for providing cat facts
- Go community for excellent documentation
- HNG Internship for the challenge
