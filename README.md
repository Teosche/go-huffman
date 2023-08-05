# Huffman Compression

Huffman Compression is a Go implementation of the Huffman coding algorithm, which is used for lossless data compression. This project provides a simple API endpoint for compressing strings using Huffman coding and MVC pattern.

## Prerequisites

To run this project, you need to have the following dependencies installed:

- Go (at least version 1.16)
- Beego
- Postman (optional)

## Getting Started

Follow the steps below to get started with the Huffman Compression project:

1. Clone the repository:
  ```
 git clone https://github.com/your-username/huffman-compression.git
  ```
   
2. Change into the project directory:

3. Run the project:

  ```
  bee run
  ```

4. The project should now be running on http://localhost:8080.

# API Endpoint

The project exposes a single API endpoint for Huffman compression:

`POST /api/v1/huffman/huffmanCompression` : Compresses a string using Huffman coding.

Example Request:

```
POST /api/v1/huffman/huffmanCompression
Content-Type: application/json

{
  "request": "your-string-to-compress"
}

```

Example Response:

```
{
    "codeMap": {
        "-": "100",
        "c": "11110",
        "e": "11111",
        "g": "0000",
        "i": "11100",
        "m": "11101",
        "n": "0010",
        "o": "010",
        "p": "1100",
        "r": "011",
        "s": "101",
        "t": "1101",
        "u": "0011",
        "y": "0001"
    }
}
```

# Usage

Make a POST request to the /api/v1/huffman/huffmanCompression endpoint with the string you want to compress.

```
curl -X POST -H "Content-Type: application/json" -d '{"request": "your-string-to-compress"}' http://localhost:8080/api/v1/huffman/huffmanCompression
```

You can also use Postman:

1. Create a new request and set the request method to POST.
2. Set the request URL to `http://localhost:8080/api/v1/huffman/huffmanCompression`.
3. Set the request body as a JSON object with the request field containing the string to compress.
4. Send the request and check the response for the compressed code.

# Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

# Authors
@Teosche
