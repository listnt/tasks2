package main

import "fmt"

type WalletInterface interface {
	CreditBalance(amount int)
	DebitBalance(amount int) error
}

type SecurityCodeInterface interface {
	CheckCode(incomingCode int) error
}

type NotificationInterface interface {
	SendWalletCreditNotification()
	SendWalletDebitNotification()
}

type LedgerInterface interface {
	MakeEntry(accountID, txnType string, amount int)
}

type AccountInterface interface {
	CheckAccount(accountName string) error
}

type wallet struct {
	balance int
}

func NewWallet() WalletInterface {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

type securityCode struct {
	code int
}

func NewSecurityCode(code int) SecurityCodeInterface {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) CheckCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

type notification struct {
}

func NewNotification() NotificationInterface {
	return &notification{}
}

func (n *notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

type ledger struct {
}

func NewLedger() LedgerInterface {
	return &ledger{}
}

func (s *ledger) MakeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
}

type account struct {
	name string
}

func NewAccount(accountName string) AccountInterface {
	return &account{
		name: accountName,
	}
}

func (a *account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
