package api

import (
	"APIS/dataservice"
	"APIS/model"
	"APIS/queue"
	"database/sql"
	"fmt"

	"github.com/IBM/sarama"
)

type IBizLogic interface {
	CreateBookLogic(book model.Book) error
}

type BizLogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

func NewBizLogic(db *sql.DB, producer sarama.SyncProducer) *BizLogic {
	return &BizLogic{DB: db, Producer: producer}
}

func (bl *BizLogic) CreateBookLogic(book model.Book) error {
	// validation by making a get request
	if book.Title == "" {
		return fmt.Errorf("Title should be present")
	}

	if err := dataservice.CreateBook(bl.DB, book); err != nil {
		return err
	}

	// produce the message to kafka
	message := fmt.Sprintf("Tile: %s by Author: %s", book.Title, book.Author)
	err := queue.ProduceKafkaMessage("book_created_topic", message, bl.Producer)
	if err != nil {
		return fmt.Errorf("failed to produce kafka message: %v", err)
	}

	return nil
}
