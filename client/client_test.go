//go:build integration

package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/bartossh/The-Accountant/fileoperations"
	"github.com/bartossh/The-Accountant/wallet"
	"github.com/stretchr/testify/assert"
)

func TestAlive(t *testing.T) {
	t.Parallel()
	c := NewRest("http://localhost:8080", 5*time.Second, wallet.Helper{}, fileoperations.Helper{}, wallet.New)
	err := c.ValidateApiVersion()
	assert.Nil(t, err)
}

func BenchmarkAlive(b *testing.B) {
	c := NewRest("http://localhost:8080", 5*time.Second, wallet.Helper{}, fileoperations.Helper{}, wallet.New)
	for i := 0; i < b.N; i++ {
		_ = c.ValidateApiVersion()
	}
}

func TestFullClientApiCycle(t *testing.T) {
	issuer := NewRest("http://localhost:8080", 5*time.Second, wallet.Helper{}, fileoperations.Helper{}, wallet.New)
	err := issuer.ValidateApiVersion()
	assert.Nil(t, err)
	err = issuer.NewWallet("wpg6d0grqJjyRicC8oI0/w6IGivm5ypFNTO/wwPGW9A=")
	assert.Nil(t, err)

	receiver := NewRest("http://localhost:8080", 5*time.Second, wallet.Helper{}, fileoperations.Helper{}, wallet.New)
	err = receiver.ValidateApiVersion()
	assert.Nil(t, err)
	err = receiver.NewWallet("GWFuhvyFnmMg1/vhPCfoa9ct1pAMC1pWwlRg4kt0D/w=")
	assert.Nil(t, err)

	receiverAddr, err := receiver.Address()
	assert.Nil(t, err)
	err = issuer.ProposeTransaction(receiverAddr, "text", []byte("test_transaction_data"))
	assert.Nil(t, err)
	issuedTrx, err := issuer.ReadIssuedTransactions()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(issuedTrx))

	awaitedTrx, err := receiver.ReadWaitingTransactions()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(awaitedTrx))

	fmt.Printf("Awaited transaction: %#v\n", awaitedTrx[0])

	err = receiver.ConfirmTransaction(awaitedTrx[0])
	fmt.Println(err)
	assert.Nil(t, err)

	issuer.FlushWalletFromMemory()
	receiver.FlushWalletFromMemory()
}
