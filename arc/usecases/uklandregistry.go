package usecases

import (
	"InlineView/arc/entities"
	"InlineView/arc/infrastructure"
)

var priceData entities.PriceStore

// PricesAndStore accepts a URL of a UK Land Registry prices file and converts the contents to prices paid stores.
func PricesAndStore(resource string) {
	repo := infrastructure.NewFileStore()
	storeIndex := entities.NewIndex(repo)

	retrieveIndex("new store", storeIndex)

}

func retrieveIndex(storename string, storer *entities.ManageIndex) {
	storer.ReadIndex("file")

	priceData.Index = storer.IndexData
}
