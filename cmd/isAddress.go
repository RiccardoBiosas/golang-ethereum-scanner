/*
Copyright © 2020 NAME HERE <rbios@protonmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"

	client "github.com/RiccardoBiosas/golang-ethereum-scanner/client"
)

// isAddressCmd represents the isAddress command
var isAddressCmd = &cobra.Command{
	Use:   "isAddress",
	Short: "CLI command to check whether a given address is a contract or an account",
	Long: `The isAddress CLI command takes a valid ethereum address as an argument and
			subsequently check whethere there's bytecode stored on the given address. If
			positive, the address is a contract and the output is true, otherwise the address
			is an account and the output is false.
			Example: golang-ethereum-scanner isAddress 0x773cc2e2cbda9945f4e69e26e516708d66e45dc2
			`,
	Run: func(cmd *cobra.Command, args []string) {
		c := client.Client{}
		c.Mount()
		checkAddress(args[0], c.EthereumClient)
	},
}

func init() {
	rootCmd.AddCommand(isAddressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isAddressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isAddressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkAddress(pb string, client *ethclient.Client) {
	fmt.Println("address as argument is ", pb)
	address := common.HexToAddress(pb)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("The address is a contract: %v \n", isContract)

}
