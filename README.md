# SMART-TODO

### A simple app built to deploy and interract with an ERC20 Token Smart Contract on ETH compatible testnets including Rinkeby testnet written in Go.

This is a simple smart contract to test out the workings of Go interractions with ERC20 Tokens

### Usage
 $ ./smartdo <command> <options>

list of functionalities and the commands to test them out


1. Create account - Create a new encrypted ethereum compatible wallet
    $ ./smartdo adduser username password

2. Deploy contract
    Deploys the pre written ERC20 Toke to the active Testnet or mainnet
	$ ./smartdo deploy username password <contractname>
3. Check Wallet Address
    $ ./smartdo mywallet username password
4. Check balance
    $ ./smartdo balance username password
5. Send wei
    $ ./smartdo sendwei username password <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <sendAmount e.g. 200>
6. Mint Token
    $ ./smartdo contract-mint username password <contractName e.g. logi> <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <amountOfTokens e.g. 200000000000000>
7. Transfer Tokens
    $ ./smartdo contract-transfer username password <contractName e.g. logi> <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <amountOfTokens e.g. 2000000000000000000>



... A work in progress


