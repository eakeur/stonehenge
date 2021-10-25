package domain

import (
	"crypto/md5"
	"encoding/hex"
	model "stonehenge/model"
	"stonehenge/providers"
	"strconv"
	"strings"
)

// How much money will the users get when creating a new account
const INITIAL_BUDGET int = 500

// Adds a new account if there is not any other with the same CPF and increments its balance
func AddNewAccount(account model.Account) (*string, error) {
	account.Cpf = NormalizeCPF(account.Cpf)
	if !IsCPFValid(account.Cpf) {
		return nil, model.ErrCPFInvalid
	}

	exists, err := providers.AccountsRepository.AccountExists(account.Cpf)
	if err != nil {
		return nil, err
	}
	if !exists {
		// All accounts start with a initial budget set by the INITIAL_BUDGET property
		// and the secret is always saved as a hash
		account.Balance = int64(INITIAL_BUDGET)
		account.Secret = HashSecret(account.Secret)
		res, err := providers.AccountsRepository.AddAccount(account)
		if err != nil {
			return nil, err
		}
		return &res, nil
	}

	return nil, model.ErrAccountExists
}

// Gets all accounts existing - no specific rule
func GetAllAccounts() ([]model.Account, error) {
	accounts, err := providers.AccountsRepository.GetAccounts(nil)
	if err != nil {
		return nil, err
	}

	for _, acc := range accounts {
		acc.Cpf = ""
		acc.Secret = ""
		acc.Balance = 0

	}

	return accounts, nil
}

// Gets the balance of the account of the id passed
func GetAccountBalance(id string) (int64, error) {
	acc, err := GetAccountById(id)
	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}

// Gets the account by id - no specific rule
func GetAccountById(id string) (*model.Account, error) {
	return providers.AccountsRepository.GetAccountById(id)
}

// Removes any special character from the CPF string
func NormalizeCPF(cpf string) string {
	return strings.Trim(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(cpf, ".", ""), "-", ""), "/", ""), ",", ""), " ")
}

// Transforms a secret string into a MD5 hash
func HashSecret(secret string) string {
	hash := md5.Sum([]byte(secret))
	return hex.EncodeToString(hash[:])
}

// This function validates if the CPF document is valid. On purpose it validates only its length, so that you, the tester, don't need to provide an
// existing CPF
func IsCPFValid(cpf string) bool {
	for _, digit := range cpf {
		_, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
	}
	return len(cpf) == 11
}
