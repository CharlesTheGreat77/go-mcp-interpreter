# go-mcp-interpreter

Execute Go code directly from your LLM using the Yaegi interpreter via MCP.

## Usage

This MCP server provides a single tool called `execute-go` that lets LLMs run Go code in real-time. No compilation needed - just pure Go interpretation.

## Quick Start

```bash
git clone https://github.com/CharlesTheGreat77/go-mcp-interpreter
cd go-mcp-interpreter
go mod init go-mcp-interpreter
go mod tidy
go build -o go-mcp-interpreter cmd/main.go
./go-mcp-interpreter
```

The server runs via stdio and exposes one tool:
- `execute-go` - Takes Go code as input, returns the output

## Openwebui
For integration with [openwebui](https://openwebui.com) is as easy as installing `mcpo` via *pip*
- `pip3 install mcpo` -> `mcpo --port 8000 -- go run cmd/main.go`


## How it works

1. LLM sends Go code to the `execute-go` tool
2. Code gets executed in Yaegi (Go interpreter)
3. Both stdout and stderr are captured and returned
4. 5-second timeout prevents infinite loops

## Features

- **No compilation** - Yaegi interprets Go code directly
- **Standard library** - Full access to Go stdlib
- **MCP compatible** - Works with Claude Desktop and other MCP clients

## Limitations

- Only Go standard library packages
- No CGO or external dependencies
- 5-second execution limit
- No persistent state between executions

## Dependencies

- `github.com/mark3labs/mcp-go` - MCP protocol
- `github.com/traefik/yaegi` - Go interpreter

