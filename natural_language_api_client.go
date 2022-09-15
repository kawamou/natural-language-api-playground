package main

import (
	"context"
	"log"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

type AnalyzeSyntaxOutput struct {
	OriginalSentences Sentences
	Tokens            Tokens
	Language          string // e.g. "ja" "en"
}

type AnalyzeSentimentOutput struct {
	// OriginalSentences は入力文章が文の区切りとして戻ってきます
	// ex. こんにちは。私の名前は太郎です。楽しい、嬉しい！やったあ
	// &main.Sentence{Text:"こんにちは。"}
	// &main.Sentence{Text:"私の名前は太郎です。"}
	// &main.Sentence{Text:"楽しい、嬉しい！"}
	// &main.Sentence{Text:"やったあ"}
	OriginalSentences Sentences
	Sentiment         Sentiment
	Language          string // e.g. "ja" "en"
}

type AnalyzeEntityOutput struct {
	Entities Entities
	Language string // e.g. "ja" "en"
	Sentiment
}

type Sentiment struct {
	Magnitude float32 // [0, +inf)
	// Score is -1.0 (negative sentiment) and 1.0
	// Score >= 0 -> positive
	Score float32
}

type Entities []*Entity

type Entity struct {
	Name     string
	Type     string
	Metadata map[string]string
	// Sentiment は感情を表します
	// EntitySentimentのAPIを叩いたときにのみ反映される値
	Sentiment Sentiment
	// Mentions等は一旦省略
}

type Tokens []*Token

type Token struct {
	Text           Text
	PartOfSpeech   PartOfSpeech
	DependencyEdge DependencyEdge
}

type Text struct {
	Content     string
	BeginOffset int
}

type PartOfSpeech struct {
	Tag    int
	Proper int
}

type DependencyEdge struct {
	HeadTokenIndex int
	Label          int
}

type Sentences []*Sentence

type Sentence struct {
	Text string
}

type Client struct {
	c *language.Client
}

func NewClient(ctx context.Context) *Client {
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		c: client,
	}
}

func NewAnalyzeSyntaxOutput(analyzeSyntaxResp *languagepb.AnalyzeSyntaxResponse) *AnalyzeSyntaxOutput {
	return &AnalyzeSyntaxOutput{
		OriginalSentences: NewSentencesFromSentencesPb(analyzeSyntaxResp.Sentences),
		Tokens:            NewTokensFromTokensPb(analyzeSyntaxResp.Tokens),
		Language:          analyzeSyntaxResp.Language,
	}
}

func NewAnalyzeSentimentOutput(analyzeSentimentResp *languagepb.AnalyzeSentimentResponse) *AnalyzeSentimentOutput {
	return &AnalyzeSentimentOutput{
		OriginalSentences: NewSentencesFromSentencesPb(analyzeSentimentResp.Sentences),
		Sentiment:         NewSentimentFromSentimentPb(analyzeSentimentResp.DocumentSentiment),
		Language:          analyzeSentimentResp.Language,
	}
}

func NewAnalyzeEntityOutput(analyzeEntityResp *languagepb.AnalyzeEntitiesResponse) *AnalyzeEntityOutput {
	return &AnalyzeEntityOutput{
		Entities: NewEntitiesFromEntitiesPb(analyzeEntityResp.Entities),
		Language: analyzeEntityResp.Language,
	}
}

func NewAnalyzeEntitySentimentOutput(analyzeEntityResp *languagepb.AnalyzeEntitySentimentResponse) *AnalyzeEntityOutput {
	return &AnalyzeEntityOutput{
		Entities: NewEntitiesFromEntitiesPb(analyzeEntityResp.Entities),
		Language: analyzeEntityResp.Language,
	}
}

func NewSentimentFromSentimentPb(sentimentPb *languagepb.Sentiment) Sentiment {
	return Sentiment{
		Magnitude: sentimentPb.Magnitude,
		Score:     sentimentPb.Score,
	}
}

func NewSentencesFromSentencesPb(sentencesPb []*languagepb.Sentence) (sentences Sentences) {
	for _, sentencePb := range sentencesPb {
		sentence := &Sentence{
			Text: sentencePb.Text.Content,
		}
		sentences = append(sentences, sentence)
	}
	return
}

func NewEntitiesFromEntitiesPb(entitiesPb []*languagepb.Entity) (entities Entities) {
	for _, entityPb := range entitiesPb {
		sentiment := entityPb.GetSentiment()
		entity := &Entity{
			Name: entityPb.Name,
			Type: entityPb.Type.String(),
			Sentiment: Sentiment{
				Magnitude: magnitude(sentiment),
				Score:     score(sentiment),
			},
		}
		entities = append(entities, entity)
	}
	return
}

func magnitude(sentiment *languagepb.Sentiment) float32 {
	if sentiment == nil {
		return 0
	}
	return sentiment.Magnitude
}

func score(sentiment *languagepb.Sentiment) float32 {
	if sentiment == nil {
		return 0
	}
	return sentiment.Score
}

func NewTokensFromTokensPb(tokensPb []*languagepb.Token) (tokens Tokens) {
	for _, tokenPb := range tokensPb {
		token := &Token{
			Text: Text{
				Content:     tokenPb.Text.Content,
				BeginOffset: int(tokenPb.Text.BeginOffset),
			},
			PartOfSpeech: PartOfSpeech{
				Tag:    int(tokenPb.PartOfSpeech.Tag),
				Proper: int(tokenPb.PartOfSpeech.Proper),
			},
			DependencyEdge: DependencyEdge{
				HeadTokenIndex: int(tokenPb.DependencyEdge.HeadTokenIndex),
				Label:          int(tokenPb.DependencyEdge.Label),
			},
		}
		tokens = append(tokens, token)
	}
	return
}

func (c *Client) AnalyzeSyntax(ctx context.Context, targetContent string) (*AnalyzeSyntaxOutput, error) {
	req := &languagepb.AnalyzeSyntaxRequest{
		Document: &languagepb.Document{
			Type: languagepb.Document_PLAIN_TEXT,
			Source: &languagepb.Document_Content{
				Content: targetContent,
			},
		},
		EncodingType: languagepb.EncodingType_UTF8,
	}
	resp, err := c.c.AnalyzeSyntax(ctx, req)
	if err != nil {
		return nil, err
	}
	output := NewAnalyzeSyntaxOutput(resp)
	return output, nil
}

func (c *Client) AnalyzeSentiment(ctx context.Context, targetContent string) (*AnalyzeSentimentOutput, error) {
	req := &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Type: languagepb.Document_PLAIN_TEXT,
			Source: &languagepb.Document_Content{
				Content: targetContent,
			},
		},
		EncodingType: languagepb.EncodingType_UTF8,
	}
	resp, err := c.c.AnalyzeSentiment(ctx, req)
	if err != nil {
		return nil, err
	}
	output := NewAnalyzeSentimentOutput(resp)
	return output, nil
}

func (c *Client) AnalyzeEntity(ctx context.Context, targetContent string) (*AnalyzeEntityOutput, error) {
	req := &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Type: languagepb.Document_PLAIN_TEXT,
			Source: &languagepb.Document_Content{
				Content: targetContent,
			},
		},
		EncodingType: languagepb.EncodingType_UTF8,
	}
	resp, err := c.c.AnalyzeEntities(ctx, req)
	if err != nil {
		return nil, err
	}
	output := NewAnalyzeEntityOutput(resp)
	return output, nil
}

func (c *Client) AnalyzeEntitySentiment(ctx context.Context, targetContent string) (*AnalyzeEntityOutput, error) {
	req := &languagepb.AnalyzeEntitySentimentRequest{
		Document: &languagepb.Document{
			Type: languagepb.Document_PLAIN_TEXT,
			Source: &languagepb.Document_Content{
				Content: targetContent,
			},
		},
		EncodingType: languagepb.EncodingType_UTF8,
	}
	resp, err := c.c.AnalyzeEntitySentiment(ctx, req)
	if err != nil {
		return nil, err
	}
	output := NewAnalyzeEntitySentimentOutput(resp)
	return output, nil
}

func (c *Client) Nyuwaize(ctx context.Context, targetContent string) (retVal string, err error) {
	analyzed, err := c.AnalyzeSyntax(ctx, targetContent)
	if err != nil {
		return "", err
	}
	for _, token := range analyzed.Tokens {
		switch token.Text.Content {
		case "。":
			retVal += "😆"
		case "、":
			retVal += "🥰"
		default:
			retVal += token.Text.Content
		}
	}
	return
}
