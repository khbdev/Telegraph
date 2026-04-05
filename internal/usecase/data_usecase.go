package usecase

import (
	"context"
	"telegraph-clone/internal/domain"
	"telegraph-clone/internal/models"
	"telegraph-clone/internal/numbergeneration"
)

type DataUsecase struct {
	dataRepo domain.DataRepository
	urlRepo  domain.URLRepository
}

func NewDataUsecase(dataRepo domain.DataRepository, urlRepo domain.URLRepository) *DataUsecase {
	return &DataUsecase{
		dataRepo: dataRepo,
		urlRepo:  urlRepo,
	}
}

// Input va Output struct handler bilan ishlash uchun
type CreateDataInput struct {
	Title     string
	YourName  string
	YourStory string
}

type CreateDataOutput struct {
	Data *models.Data
	URL  *models.URL
}

// CreateData usecase
func (uc *DataUsecase) CreateData(ctx context.Context, input CreateDataInput) (*CreateDataOutput, error) {
	// 1️⃣ Data model yaratish
	data := &models.Data{
		Title:     input.Title,
		YourName:  input.YourName,
		YourStory: input.YourStory,
	}

	// 2️⃣ DataRepository.Create chaqirish
	createdData, err := uc.dataRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	// 3️⃣ title ga random raqam qo'shish
	randomTitle := numbergeneration.AddRandomNumberPrefix(input.Title)

	// 4️⃣ URL model yaratish
	url := &models.URL{
		URL:    randomTitle,
		DataID: createdData.ID,
	}

	// 5️⃣ URLRepository.Create chaqirish
	if err := uc.urlRepo.Create(ctx, url); err != nil {
		return nil, err
	}

	// 6️⃣ Natijani qaytarish
	return &CreateDataOutput{
		Data: createdData,
		URL:  url,
	}, nil
}