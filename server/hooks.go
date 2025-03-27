// Code generated by `go generate`. DO NOT EDIT.
// source: server/internal/gen/hooks.go.tmpl
package server

import (
	"github.com/aavaz-ai/mcp-go/mcp"
)

// BeforeAnyHookFunc is a function that is called after the request is
// parsed but before the method is called.
type BeforeAnyHookFunc func(id any, method mcp.MCPMethod, message any)

// OnSuccessHookFunc is a hook that will be called after the request
// successfully generates a result, but before the result is sent to the client.
type OnSuccessHookFunc func(id any, method mcp.MCPMethod, message any, result any)

// OnErrorHookFunc is a hook that will be called when an error occurs,
// either during the request parsing or the method execution.
//
// Example usage:
// ```
//
//	hooks.AddOnError(func(id any, method mcp.MCPMethod, message any, err error) {
//	  // Check for specific error types using errors.Is
//	  if errors.Is(err, ErrUnsupported) {
//	    // Handle capability not supported errors
//	    log.Printf("Capability not supported: %v", err)
//	  }
//
//	  // Use errors.As to get specific error types
//	  var parseErr = &UnparseableMessageError{}
//	  if errors.As(err, &parseErr) {
//	    // Access specific methods/fields of the error type
//	    log.Printf("Failed to parse message for method %s: %v",
//	               parseErr.GetMethod(), parseErr.Unwrap())
//	    // Access the raw message that failed to parse
//	    rawMsg := parseErr.GetMessage()
//	  }
//
//	  // Check for specific resource/prompt/tool errors
//	  switch {
//	  case errors.Is(err, ErrResourceNotFound):
//	    log.Printf("Resource not found: %v", err)
//	  case errors.Is(err, ErrPromptNotFound):
//	    log.Printf("Prompt not found: %v", err)
//	  case errors.Is(err, ErrToolNotFound):
//	    log.Printf("Tool not found: %v", err)
//	  }
//	})
type OnErrorHookFunc func(id any, method mcp.MCPMethod, message any, err error)

type OnBeforeInitializeFunc func(id any, message *mcp.InitializeRequest)
type OnAfterInitializeFunc func(id any, message *mcp.InitializeRequest, result *mcp.InitializeResult)

type OnBeforePingFunc func(id any, message *mcp.PingRequest)
type OnAfterPingFunc func(id any, message *mcp.PingRequest, result *mcp.EmptyResult)

type OnBeforeListResourcesFunc func(id any, message *mcp.ListResourcesRequest)
type OnAfterListResourcesFunc func(id any, message *mcp.ListResourcesRequest, result *mcp.ListResourcesResult)

type OnBeforeListResourceTemplatesFunc func(id any, message *mcp.ListResourceTemplatesRequest)
type OnAfterListResourceTemplatesFunc func(id any, message *mcp.ListResourceTemplatesRequest, result *mcp.ListResourceTemplatesResult)

type OnBeforeReadResourceFunc func(id any, message *mcp.ReadResourceRequest)
type OnAfterReadResourceFunc func(id any, message *mcp.ReadResourceRequest, result *mcp.ReadResourceResult)

type OnBeforeListPromptsFunc func(id any, message *mcp.ListPromptsRequest)
type OnAfterListPromptsFunc func(id any, message *mcp.ListPromptsRequest, result *mcp.ListPromptsResult)

type OnBeforeGetPromptFunc func(id any, message *mcp.GetPromptRequest)
type OnAfterGetPromptFunc func(id any, message *mcp.GetPromptRequest, result *mcp.GetPromptResult)

type OnBeforeListToolsFunc func(id any, message *mcp.ListToolsRequest)
type OnAfterListToolsFunc func(id any, message *mcp.ListToolsRequest, result *mcp.ListToolsResult)

type OnBeforeCallToolFunc func(id any, message *mcp.CallToolRequest)
type OnAfterCallToolFunc func(id any, message *mcp.CallToolRequest, result *mcp.CallToolResult)

type Hooks struct {
	OnBeforeAny                   []BeforeAnyHookFunc
	OnSuccess                     []OnSuccessHookFunc
	OnError                       []OnErrorHookFunc
	OnBeforeInitialize            []OnBeforeInitializeFunc
	OnAfterInitialize             []OnAfterInitializeFunc
	OnBeforePing                  []OnBeforePingFunc
	OnAfterPing                   []OnAfterPingFunc
	OnBeforeListResources         []OnBeforeListResourcesFunc
	OnAfterListResources          []OnAfterListResourcesFunc
	OnBeforeListResourceTemplates []OnBeforeListResourceTemplatesFunc
	OnAfterListResourceTemplates  []OnAfterListResourceTemplatesFunc
	OnBeforeReadResource          []OnBeforeReadResourceFunc
	OnAfterReadResource           []OnAfterReadResourceFunc
	OnBeforeListPrompts           []OnBeforeListPromptsFunc
	OnAfterListPrompts            []OnAfterListPromptsFunc
	OnBeforeGetPrompt             []OnBeforeGetPromptFunc
	OnAfterGetPrompt              []OnAfterGetPromptFunc
	OnBeforeListTools             []OnBeforeListToolsFunc
	OnAfterListTools              []OnAfterListToolsFunc
	OnBeforeCallTool              []OnBeforeCallToolFunc
	OnAfterCallTool               []OnAfterCallToolFunc
}

func (c *Hooks) AddBeforeAny(hook BeforeAnyHookFunc) {
	c.OnBeforeAny = append(c.OnBeforeAny, hook)
}

func (c *Hooks) AddOnSuccess(hook OnSuccessHookFunc) {
	c.OnSuccess = append(c.OnSuccess, hook)
}

// AddOnError registers a hook function that will be called when an error occurs.
// The error parameter contains the actual error object, which can be interrogated
// using Go's error handling patterns like errors.Is and errors.As.
//
// Example:
// ```
// // Create a channel to receive errors for testing
// errChan := make(chan error, 1)
//
// // Register hook to capture and inspect errors
// hooks := &Hooks{}
//
//	hooks.AddOnError(func(id any, method mcp.MCPMethod, message any, err error) {
//	    // For capability-related errors
//	    if errors.Is(err, ErrUnsupported) {
//	        // Handle capability not supported
//	        errChan <- err
//	        return
//	    }
//
//	    // For parsing errors
//	    var parseErr = &UnparseableMessageError{}
//	    if errors.As(err, &parseErr) {
//	        // Handle unparseable message errors
//	        fmt.Printf("Failed to parse %s request: %v\n",
//	                   parseErr.GetMethod(), parseErr.Unwrap())
//	        errChan <- parseErr
//	        return
//	    }
//
//	    // For resource/prompt/tool not found errors
//	    if errors.Is(err, ErrResourceNotFound) ||
//	       errors.Is(err, ErrPromptNotFound) ||
//	       errors.Is(err, ErrToolNotFound) {
//	        // Handle not found errors
//	        errChan <- err
//	        return
//	    }
//
//	    // For other errors
//	    errChan <- err
//	})
//
// server := NewMCPServer("test-server", "1.0.0", WithHooks(hooks))
// ```
func (c *Hooks) AddOnError(hook OnErrorHookFunc) {
	c.OnError = append(c.OnError, hook)
}

func (c *Hooks) beforeAny(id any, method mcp.MCPMethod, message any) {
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeAny {
		hook(id, method, message)
	}
}

func (c *Hooks) onSuccess(id any, method mcp.MCPMethod, message any, result any) {
	if c == nil {
		return
	}
	for _, hook := range c.OnSuccess {
		hook(id, method, message, result)
	}
}

// onError calls all registered error hooks with the error object.
// The err parameter contains the actual error that occurred, which implements
// the standard error interface and may be a wrapped error or custom error type.
//
// This allows consumer code to use Go's error handling patterns:
// - errors.Is(err, ErrUnsupported) to check for specific sentinel errors
// - errors.As(err, &customErr) to extract custom error types
//
// Common error types include:
// - ErrUnsupported: When a capability is not enabled
// - UnparseableMessageError: When request parsing fails
// - ErrResourceNotFound: When a resource is not found
// - ErrPromptNotFound: When a prompt is not found
// - ErrToolNotFound: When a tool is not found
func (c *Hooks) onError(id any, method mcp.MCPMethod, message any, err error) {
	if c == nil {
		return
	}
	for _, hook := range c.OnError {
		hook(id, method, message, err)
	}
}
func (c *Hooks) AddBeforeInitialize(hook OnBeforeInitializeFunc) {
	c.OnBeforeInitialize = append(c.OnBeforeInitialize, hook)
}

func (c *Hooks) AddAfterInitialize(hook OnAfterInitializeFunc) {
	c.OnAfterInitialize = append(c.OnAfterInitialize, hook)
}

func (c *Hooks) beforeInitialize(id any, message *mcp.InitializeRequest) {
	c.beforeAny(id, mcp.MethodInitialize, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeInitialize {
		hook(id, message)
	}
}

func (c *Hooks) afterInitialize(id any, message *mcp.InitializeRequest, result *mcp.InitializeResult) {
	c.onSuccess(id, mcp.MethodInitialize, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterInitialize {
		hook(id, message, result)
	}
}
func (c *Hooks) AddBeforePing(hook OnBeforePingFunc) {
	c.OnBeforePing = append(c.OnBeforePing, hook)
}

func (c *Hooks) AddAfterPing(hook OnAfterPingFunc) {
	c.OnAfterPing = append(c.OnAfterPing, hook)
}

func (c *Hooks) beforePing(id any, message *mcp.PingRequest) {
	c.beforeAny(id, mcp.MethodPing, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforePing {
		hook(id, message)
	}
}

func (c *Hooks) afterPing(id any, message *mcp.PingRequest, result *mcp.EmptyResult) {
	c.onSuccess(id, mcp.MethodPing, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterPing {
		hook(id, message, result)
	}
}
func (c *Hooks) AddBeforeListResources(hook OnBeforeListResourcesFunc) {
	c.OnBeforeListResources = append(c.OnBeforeListResources, hook)
}

func (c *Hooks) AddAfterListResources(hook OnAfterListResourcesFunc) {
	c.OnAfterListResources = append(c.OnAfterListResources, hook)
}

func (c *Hooks) beforeListResources(id any, message *mcp.ListResourcesRequest) {
	c.beforeAny(id, mcp.MethodResourcesList, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeListResources {
		hook(id, message)
	}
}

func (c *Hooks) afterListResources(id any, message *mcp.ListResourcesRequest, result *mcp.ListResourcesResult) {
	c.onSuccess(id, mcp.MethodResourcesList, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterListResources {
		hook(id, message, result)
	}
}
func (c *Hooks) AddBeforeListResourceTemplates(hook OnBeforeListResourceTemplatesFunc) {
	c.OnBeforeListResourceTemplates = append(c.OnBeforeListResourceTemplates, hook)
}

func (c *Hooks) AddAfterListResourceTemplates(hook OnAfterListResourceTemplatesFunc) {
	c.OnAfterListResourceTemplates = append(c.OnAfterListResourceTemplates, hook)
}

func (c *Hooks) beforeListResourceTemplates(id any, message *mcp.ListResourceTemplatesRequest) {
	c.beforeAny(id, mcp.MethodResourcesTemplatesList, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeListResourceTemplates {
		hook(id, message)
	}
}

func (c *Hooks) afterListResourceTemplates(id any, message *mcp.ListResourceTemplatesRequest, result *mcp.ListResourceTemplatesResult) {
	c.onSuccess(id, mcp.MethodResourcesTemplatesList, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterListResourceTemplates {
		hook(id, message, result)
	}
}
func (c *Hooks) AddBeforeReadResource(hook OnBeforeReadResourceFunc) {
	c.OnBeforeReadResource = append(c.OnBeforeReadResource, hook)
}

func (c *Hooks) AddAfterReadResource(hook OnAfterReadResourceFunc) {
	c.OnAfterReadResource = append(c.OnAfterReadResource, hook)
}

func (c *Hooks) beforeReadResource(id any, message *mcp.ReadResourceRequest) {
	c.beforeAny(id, mcp.MethodResourcesRead, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeReadResource {
		hook(id, message)
	}
}

func (c *Hooks) afterReadResource(id any, message *mcp.ReadResourceRequest, result *mcp.ReadResourceResult) {
	c.onSuccess(id, mcp.MethodResourcesRead, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterReadResource {
		hook(id, message, result)
	}
}
func (c *Hooks) AddBeforeListPrompts(hook OnBeforeListPromptsFunc) {
	c.OnBeforeListPrompts = append(c.OnBeforeListPrompts, hook)
}

func (c *Hooks) AddAfterListPrompts(hook OnAfterListPromptsFunc) {
	c.OnAfterListPrompts = append(c.OnAfterListPrompts, hook)
}

func (c *Hooks) beforeListPrompts(id any, message *mcp.ListPromptsRequest) {
	c.beforeAny(id, mcp.MethodPromptsList, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeListPrompts {
		hook(id, message)
	}
}

func (c *Hooks) afterListPrompts(id any, message *mcp.ListPromptsRequest, result *mcp.ListPromptsResult) {
	c.onSuccess(id, mcp.MethodPromptsList, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterListPrompts {
		hook(id, message, result)
	}
}
func (c *Hooks) AddBeforeGetPrompt(hook OnBeforeGetPromptFunc) {
	c.OnBeforeGetPrompt = append(c.OnBeforeGetPrompt, hook)
}

func (c *Hooks) AddAfterGetPrompt(hook OnAfterGetPromptFunc) {
	c.OnAfterGetPrompt = append(c.OnAfterGetPrompt, hook)
}

func (c *Hooks) beforeGetPrompt(id any, message *mcp.GetPromptRequest) {
	c.beforeAny(id, mcp.MethodPromptsGet, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeGetPrompt {
		hook(id, message)
	}
}

func (c *Hooks) afterGetPrompt(id any, message *mcp.GetPromptRequest, result *mcp.GetPromptResult) {
	c.onSuccess(id, mcp.MethodPromptsGet, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterGetPrompt {
		hook(id, message, result)
	}
}
func (c *Hooks) AddBeforeListTools(hook OnBeforeListToolsFunc) {
	c.OnBeforeListTools = append(c.OnBeforeListTools, hook)
}

func (c *Hooks) AddAfterListTools(hook OnAfterListToolsFunc) {
	c.OnAfterListTools = append(c.OnAfterListTools, hook)
}

func (c *Hooks) beforeListTools(id any, message *mcp.ListToolsRequest) {
	c.beforeAny(id, mcp.MethodToolsList, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeListTools {
		hook(id, message)
	}
}

func (c *Hooks) afterListTools(id any, message *mcp.ListToolsRequest, result *mcp.ListToolsResult) {
	c.onSuccess(id, mcp.MethodToolsList, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterListTools {
		hook(id, message, result)
	}
}
func (c *Hooks) AddBeforeCallTool(hook OnBeforeCallToolFunc) {
	c.OnBeforeCallTool = append(c.OnBeforeCallTool, hook)
}

func (c *Hooks) AddAfterCallTool(hook OnAfterCallToolFunc) {
	c.OnAfterCallTool = append(c.OnAfterCallTool, hook)
}

func (c *Hooks) beforeCallTool(id any, message *mcp.CallToolRequest) {
	c.beforeAny(id, mcp.MethodToolsCall, message)
	if c == nil {
		return
	}
	for _, hook := range c.OnBeforeCallTool {
		hook(id, message)
	}
}

func (c *Hooks) afterCallTool(id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
	c.onSuccess(id, mcp.MethodToolsCall, message, result)
	if c == nil {
		return
	}
	for _, hook := range c.OnAfterCallTool {
		hook(id, message, result)
	}
}
