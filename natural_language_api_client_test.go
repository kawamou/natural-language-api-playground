package main

import (
	"context"
	"log"
	"testing"
)

func TestClient_AnalyzeSyntax(t *testing.T) {
	ctx := context.Background()
	client := NewClient(ctx)
	output, err := client.AnalyzeSyntax(ctx, "こんにちは。私の名前は太郎です。楽しい、嬉しい！やったあ")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", output.Language)
	log.Printf("%#v", output.OriginalSentences[0])
	for _, token := range output.Tokens {
		log.Printf("%#v", *token)
	}
}

func TestClient_AnalyzeSentiment(t *testing.T) {
	ctx := context.Background()
	client := NewClient(ctx)
	output, err := client.AnalyzeSentiment(ctx, "こんにちは。私の名前は太郎です。楽しい、嬉しい！やったあ")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", output.Language)
	for _, sentence := range output.OriginalSentences {
		log.Printf("%#v", sentence)
	}
	log.Printf("%#v", output.Sentiment.Magnitude)
	log.Printf("%#v", output.Sentiment.Score)
}
