package pattern

import (
	"fmt"
	"log"
)

// Фасад
type WalletFacadeInterface interface {
	AddMoneyToWallet(accountID string, securityCode int, amount int) error
	DeductMoneyFromWallet(accountID string, securityCode int, amount int) error
}

type walletFacade struct {
	account      AccountInterface
	wallet       WalletInterface
	securityCode SecurityCodeInterface
	notification NotificationInterface
	ledger       LedgerInterface
}

func NewWalletFacade(accountID string, code int) WalletFacadeInterface {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      NewAccount(accountID),
		securityCode: NewSecurityCode(code),
		wallet:       NewWallet(),
		notification: NewNotification(),
		ledger:       NewLedger(),
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *walletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

// Различные классы

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
		return fmt.Errorf("security Code is incorrect")
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
		return fmt.Errorf("account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

func driverCode() {
	fmt.Println()
	walletFacade := NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
