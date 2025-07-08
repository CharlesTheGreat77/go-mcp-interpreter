package main

import (
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"go-mcp-interpreter/handler"
)

func main() {
	mcpServer := server.NewMCPServer(
		"golang-interpreter",
		"1.0.0",
	)

	golangTool := mcp.NewTool(
		"execute-go",
		mcp.WithDescription(
			"Execute go code in an isolated environment. Yaegi is an interpreter for go and will be used to execute the code provided. Only native packages are supported.",
		),
		mcp.WithString(
			"code",
			mcp.Description("The golang code to execute"),
			mcp.Required(),
		),
	)

	mcpServer.AddTool(golangTool, handler.HandleGolangExecution)

	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatalf("[-] Server error: %v", err)
	}
}
