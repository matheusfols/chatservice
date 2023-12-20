package main

import (
	"database/sql"
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/matheusfols/chatservice/configs"
	"github.com/matheusfols/chatservice/internal/infra/grpc/server"
	"github.com/matheusfols/chatservice/internal/infra/repository"
	"github.com/matheusfols/chatservice/internal/infra/web"
	"github.com/matheusfols/chatservice/internal/infra/web/webserver"
	"github.com/matheusfols/chatservice/internal/usecase/chatcompletion"
	"github.com/matheusfols/chatservice/internal/usecase/chatcompletionstream"
	"github.com/sashabaranov/go-openai"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		configs.DBUser,
		configs.DBPassword,
		configs.DBHost,
		configs.DBPort,
		configs.DBName)

	conn, err := sql.Open(configs.DBDriver, url)
	fmt.Println("Connecting to database...")
	fmt.Println(conn)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := repository.NewChatRepositoryMySQL(conn)
	client := openai.NewClient(configs.OpenAIApiKey)

	chatConfig := chatcompletion.ChatCompletionConfigInputDTO{
		Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
	}

	chatConfigStream := chatcompletionstream.ChatCompletionConfigInputDTO{
		Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
	}

	usecase := chatcompletion.NewChatCompletionUseCase(repo, client)

	streamChannel := make(chan chatcompletionstream.ChatCompletionOutputDTO)
	usecaseStream := chatcompletionstream.NewChatCompletionUseCase(repo, client, streamChannel)

	fmt.Println("Starting gRPC server on port " + configs.GRPCServerPort)
	grpcServer := server.NewGRPCServer(
		*usecaseStream,
		chatConfigStream,
		configs.GRPCServerPort,
		configs.AuthToken,
		streamChannel,
	)
	go grpcServer.Start()

	webserver := webserver.NewWebServer(":" + configs.WebServerPort)
	webserverChatHandler := web.NewWebChatGPTHandler(*usecase, chatConfig, configs.AuthToken)
	webserver.AddHandler("/chat", webserverChatHandler.Handle)

	fmt.Println("Server running on port " + configs.WebServerPort)
	webserver.Start()
}
