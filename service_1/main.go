package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serviceBConn := createServiceBConnection()
	defer serviceBConn.Close()
	loadConfig()
	server := fasthttp.Server{
		Handler: requestHandler,
	}
	if err := server.ListenAndServe("127.0.0.1:8080"); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
}

func loadConfig() {
	// Установка значений по умолчанию для конфигурационных параметров
	viper.SetDefault("service_a.port", 8080)
	// Загрузка файлов конфигурации .env
	if err := godotenv.Load(); err != nil {
		log.Println("Ошибка загрузки .env")
	}
	// Чтение переменных окружения
	viper.AutomaticEnv()
}

func createServiceBConnection() *grpc.ClientConn {
	// Получение адреса сервера ServiceB из конфигурации
	serviceBAddress := viper.GetString("service_b.address")
	creds := credentials.NewTLS(nil)
	// Установка параметров подключения к gRPC серверу ServiceB
	conn, err := grpc.Dial(serviceBAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf("Failed to connect to ServiceB: %v", err)
	}
	return conn
}