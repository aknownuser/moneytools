package storage

import (
	"moneytool/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	GetAccounts() ([]models.Account, error)
	GetAccount(id string) (models.Account, error)
	CreateAccount(account models.Account) error
	//UpdateAccount(account models.Account) error
	//DeleteAccount(id string) error
	//GetTransactions(accountID string) ([]models.Transaction, error)
	//CreateTransaction(transaction models.Transaction) error
	//UpdateTransaction(transaction models.Transaction) error
	//DeleteTransaction(id string) error
	//GetCategories() ([]string, error)
	//GetCategory(category string) ([]models.Transaction, error)
	//GetTransactionsByCategory(category string) ([]models.Transaction, error)
	//GetTransactionsByDateRange(startDate, endDate time.Time) ([]models.Transaction, error)
	//GetTransactionsByDate(date time.Time) ([]models.Transaction, error)
	//GetTransactionByAccountAndDate(accountID string, date time.Time) ([]models.Transaction, error)
}

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) GetAccounts() ([]models.Account, error) {
	var accounts []models.Account
	result := s.db.Find(&accounts)
	return accounts, result.Error
}

func (s *PostgresStorage) GetAccount(id string) (models.Account, error) {
	var account models.Account
	result := s.db.First(&account, id)
	return account, result.Error
}

func (s *PostgresStorage) CreateAccount(account models.Account) error {
	result := s.db.Create(&account)
	return result.Error
}
