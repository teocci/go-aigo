# AIGO: AI Search Engine

AIGO is a minimalistic AI-powered search engine implemented in Go using the Fiber framework. This project serves as a foundation for developing robust and scalable AI-driven applications.

## Features
- **Single Page Interface**: A simple, Google-like search input for user queries.
- **Lightweight Framework**: Built with Fiber, ensuring speed and efficiency.
- **Customizable**: Uses templates for UI and allows seamless CSS and JavaScript integration.
- **Modular Design**: Structured to support future expansions and additional features.

## Prerequisites
- Go 1.20 or later
- Node.js (optional, for managing JavaScript modules and CSS pre-processing)
- A web browser

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/teocci/aigo.git
   cd aigo
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create necessary directories for static files and templates:
   ```bash
   mkdir -p public/css public/js views
   ```

4. Add your custom CSS and JavaScript files to `public/css/` and `public/js/` respectively.

5. Ensure the `config.json` file exists in the root directory with appropriate settings.

## Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```

2. Open your browser and navigate to:
   ```
   http://localhost:3000
   ```

## Project Structure
```
AIGO
├── main.go               # Entry point of the application
├── src/
│   ├── cmd/              # Command-line commands and execution logic
│   ├── config/           # Configuration management
│   ├── webserver/        # Server logic and route handling
├── public/               # Public assets (CSS, JS, images, etc.)
│   ├── css/              # CSS files
│   ├── js/               # JavaScript modules
├── views/                # HTML templates
├── go.mod                # Go module file
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Configuration
The `config.json` file and `.env` file are used to manage application settings:

### Example `config.json`
```json
{
  "web": {
    "port": 3000
  },
  "profile": "prod",
  "profiles": {
    "prod": {
      "api": {
        "host": "localhost",
        "port": 8080
      }
    }
  }
}
```

### Example `.env`
```env
PROFILE=prod
PORT=3000
```

## Contributing
Feel free to fork this repository, submit pull requests, or file issues.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

---

Happy coding with AIGO! 🚀

