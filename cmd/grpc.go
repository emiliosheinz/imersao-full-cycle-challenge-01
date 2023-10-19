/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/emiliosheinz/imersao-full-cycle-challenge-01/application/grpc"
	"github.com/emiliosheinz/imersao-full-cycle-challenge-01/infra/db"
	"github.com/spf13/cobra"
)

var portNumber int

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Starts a gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB()
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC Server port")
}
