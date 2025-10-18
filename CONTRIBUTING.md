# Contributing to Profile API

Thank you for your interest in contributing to this project! 🎉

## Getting Started

1. **Fork the repository**
2. **Clone your fork:**
   ```bash
   git clone https://github.com/yourusername/profile-api.git
   cd profile-api
   ```

3. **Create a new branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Setup

### Prerequisites
- Go 1.21 or higher
- Git
- A code editor (VS Code, GoLand, etc.)

### Installation
```bash
# Install dependencies
go mod download

# Run the application
make run
# OR
go run main.go
```

### Running Tests
```bash
# Run all tests
make test

# Run tests with coverage
make test-cover
```

## Code Style

### Go Conventions
- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` to format your code
- Use meaningful variable and function names
- Add comments for exported functions and types

### Running Formatter
```bash
make fmt
# OR
go fmt ./...
```

### Linting
```bash
make lint
# OR
golangci-lint run
```

## Making Changes

### Commit Messages
Follow the conventional commits specification:

- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `style:` Code style changes (formatting, etc.)
- `refactor:` Code refactoring
- `test:` Adding or updating tests
- `chore:` Maintenance tasks

**Examples:**
```
feat: add rate limiting middleware
fix: handle nil pointer in fetchCatFact
docs: update README with deployment instructions
test: add unit tests for profileHandler
```

### Pull Request Process

1. **Update documentation** if needed
2. **Add tests** for new features
3. **Ensure all tests pass:**
   ```bash
   go test ./...
   ```

4. **Update the CHANGELOG.md** (if applicable)

5. **Submit your Pull Request** with:
   - Clear title and description
   - Reference to any related issues
   - Screenshots (if UI changes)

## Feature Requests

Have an idea? Open an issue with:
- Clear description of the feature
- Use case/motivation
- Possible implementation approach

## Bug Reports

Found a bug? Open an issue with:
- Steps to reproduce
- Expected behavior
- Actual behavior
- Environment details (OS, Go version)
- Error logs (if any)

## Code of Conduct

### Our Standards

- Be respectful and inclusive
- Accept constructive criticism
- Focus on what's best for the community
- Show empathy towards others

### Unacceptable Behavior

- Harassment or discriminatory language
- Trolling or insulting comments
- Public or private harassment
- Publishing others' private information

## Questions?

Feel free to open an issue for any questions about contributing!

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing! 🙏