package main

import (
	"fmt"
)

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
