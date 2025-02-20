// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	llm, err := gollm.NewVertexAIClient(ctx)
	if err != nil {
		return fmt.Errorf("creating vertexai client: %w", err)
	}
	systemPrompt := ""
	chat := llm.StartChat(systemPrompt)

	response, err := chat.SendMessage(ctx, "Hello, world!")
	if err != nil {
		return fmt.Errorf("sending message: %w", err)
	}
	fmt.Printf("response: %v\n", response)

	return nil
}
