# URL Shortener

A simple URL shortener service built with Go (Golang).

## Features

- Shorten long URLs
- Redirect shortened URLs to the original URLs
- Basic in-memory storage for shortened URLs

## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/AustinKingOry/go-url-shortener.git
    cd url-shortener
    ```

2. Initialize the Go module:
    ```sh
    go mod init url-shortener
    ```

### Running the Application

1. Start the server:
    ```sh
    go run main.go
    ```

2. The server will run on `http://localhost:8080`.

### API Endpoints

#### Home

- **URL**: `/`
- **Method**: `GET`
- **Description**: Displays a welcome message.

#### Shorten URL

- **URL**: `/shorten`
- **Method**: `POST`
- **Request Body**: JSON object containing the long URL to be shortened.
    ```json
    {
        "url": "https://www.example.com"
    }
    ```
- **Response**: JSON object containing the shortened URL.
    ```json
    {
        "short_url": "http://localhost:8080/u/abc123"
    }
    ```

#### Redirect to Long URL

- **URL**: `/u/{shortURL}`
- **Method**: `GET`
- **Description**: Redirects to the original long URL.

### Example Usage

1. Shorten a URL:
    ```sh
    curl -X POST -d '{"url":"https://www.example.com"}' -H "Content-Type: application/json" http://localhost:8080/shorten
    ```

    Example response:
    ```json
    {
        "short_url": "http://localhost:8080/u/abc123"
    }
    ```

2. Visit the shortened URL:
    - Open `http://localhost:8080/u/abc123` in your browser. It should redirect to `https://www.example.com`.

### Additional Features to Implement

- Error handling and validation for the URL shortening process
- Persistent storage (e.g., database) for URLs
- Front-end interface for easier URL submission
- Rate limiting to prevent abuse
- Custom short URLs specified by users

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for their excellent documentation and tutorials.
