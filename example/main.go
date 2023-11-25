package main

import next_api "github.com/dkZzzz/openai_next_api"

func main() {
	questions := make([]string, 0)
	next_api_key := ""
	next_api.Chat(next_api_key, questions)
}
