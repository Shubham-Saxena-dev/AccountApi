package services

import (
	"GolangProject/apis"
	"GolangProject/repositories"
	"errors"
)

type AccountService interface {
	GetAccounts() ([]apis.AccountCreateRequest, error)
	GetAccount(id string) (apis.AccountCreateRequest, error)
	CreateAccount(request apis.AccountCreateRequest) (apis.AccountCreateRequest, error)
	DeleteAccount(id string) error
	UpdateAccount(id string, request apis.AccountUpdateRequest) (apis.AccountCreateRequest, error)
}

func NewService(repo repositories.Repositories) AccountService {
	return &accountService{
		repo: repo,
	}
}

type accountService struct {
	repo repositories.Repositories
}

func (a *accountService) GetAccounts() ([]apis.AccountCreateRequest, error) {
	var accounts []apis.AccountCreateRequest

	acc, err := a.repo.GetAccounts()

	if err != nil {
		return []apis.AccountCreateRequest{}, err
	}

	for _, account := range acc {
		accounts = append(accounts, account)
	}

	if len(accounts) < 1 {
		return []apis.AccountCreateRequest{}, errors.New("no account found")
	}

	return accounts, nil
}

func (a *accountService) GetAccount(id string) (apis.AccountCreateRequest, error) {
	acc, err := a.repo.GetAccount(id)

	if err != nil {
		return apis.EmptyCreateAccount, err
	}

	return acc, nil
}

func (a *accountService) CreateAccount(request apis.AccountCreateRequest) (apis.AccountCreateRequest, error) {
	acc, err := a.repo.CreateAccount(request)

	if err != nil {
		return apis.EmptyCreateAccount, err
	}

	return acc, nil
}

func (a *accountService) DeleteAccount(id string) error {
	err := a.repo.DeleteAccount(id)
	if err != nil {
		return err
	}
	return nil
}

func (a *accountService) UpdateAccount(id string, request apis.AccountUpdateRequest) (apis.AccountCreateRequest, error) {
	acc, err := a.repo.UpdateAccount(id, request)

	if err != nil {
		return apis.EmptyCreateAccount, err
	}
	return acc, nil
}
