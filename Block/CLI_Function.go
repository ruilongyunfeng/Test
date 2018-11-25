package Block

import "fmt"

/* *************************************************************
 * Copyright  2018 Bridge-ruijiezhi@163.com. All rights reserved.
 *
 * FileName: CLI_Function
 *
 * @Author: Bridge 2018/11/22 19:58
 *
 * @Version: 1.0
 * *************************************************************/

func (cli *CLI) createWallet(nodeID string) {

	wallets, _ := NewWallets(nodeID)

	wallets.CreateNewWallet(nodeID)

	fmt.Println(len(wallets.walletsMap))
}

func (cli *CLI) createGenesisBlockchain(address string, nodeID string) {
	blockChain := CreateBlockchainWithGenesisBlock(address, nodeID)
	defer blockChain.DB.Close()

	utxoSet := &UTXOSet{blockChain}

	utxoSet.ResetUTXOSet()
}

func (cli *CLI) addressList(nodeID string) {
	fmt.Println("All wallet address:")

	wallets, _ := NewWallets(nodeID)

	for address, _ := range wallets.walletsMap {
		fmt.Println(address)
	}
}

func (cli *CLI) getBalance(address string, nodeID string) {
	fmt.Println("address : " + address)

	blockChain := BlockChainObject(nodeID)
	defer blockChain.DB.Close()

	utxoSet := &UTXOSet{blockChain}

	amount := utxoSet.GetBalance(address)

	fmt.Printf("%s have %d tokens.\n", address, amount)
}

func (cli *CLI) printChain(nodeID string) {
	blockChain := BlockChainObject(nodeID)

	defer blockChain.DB.Close()

	blockChain.PrintChain()
}

func (cli *CLI) send(from []string, to []string, amount []string, nodeID string, mineNow bool) {
	blockChain := BlockChainObject(nodeID)

	utxoSet = &UTXOSet{blockChain}
	defer blockChain.DB.Close()

	if mineNow {
		blockChain.MineNewBlock(from, to, amount, nodeID)
		utxoSet.Update()
	} else {
		value, _ := strconv.Atoi(amount[0])
		tx := NewSimpleTransaction(from[0], to[0], value, utxoSet, []*Transaction{}, nodeID)

		sendTx(knowNodes[0], tx)
	}
}
