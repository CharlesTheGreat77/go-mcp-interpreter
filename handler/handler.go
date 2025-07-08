package handler

import (
	"context"
	"fmt"
	"go-mcp-interpreter/internal"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleGolangExecution(
	ctx context.Context,
	request mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	code, err := request.RequireString("code")
	if err != nil || code == "" {
		return mcp.NewToolResultError("Missing or invalid `code` argument"), nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second) // execution timeout
	defer cancel()
	resultCh := make(chan *mcp.CallToolResult, 1)
	go func() {
		outputStr, err := internal.ExecuteGoCodeInYaegi(timeoutCtx, code)

		if err != nil {
			resultCh <- mcp.NewToolResultError(fmt.Sprintf("Execution error: %v\nCaptured output:\n%s", err, outputStr))
		} else {
			if outputStr == "" {
				resultCh <- mcp.NewToolResultText("Code executed successfully, but no output was captured.")
			} else {
				resultCh <- mcp.NewToolResultText(outputStr)
			}
		}
	}()

	select {
	case <-timeoutCtx.Done():
		return mcp.NewToolResultError("Execution timed out after 5 seconds"), nil
	case res := <-resultCh:
		return res, nil
	}
}
