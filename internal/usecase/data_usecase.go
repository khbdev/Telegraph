package usecase

import (
	"context"
	"telegraph-clone/internal/domain"
	"telegraph-clone/internal/models"
	numbergeneration "telegraph-clone/pkg/numberGeneration"
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




type CreateDataOutput struct {
	URL  string
}


func (uc *DataUsecase) CreateData(ctx context.Context, input models.CreateDataInput) (*CreateDataOutput, error) {

	data := &models.Data{
		Title:     input.Title,
		YourName:  input.YourName,
		YourStory: input.YourStory,
	}

	
	createdData, err := uc.dataRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}


	randomTitle := numbergeneration.AddRandomNumberPrefix(input.Title)


	url := &models.URL{
		URL:    randomTitle,
		DataID: createdData.ID,
	}


	if err := uc.urlRepo.Create(ctx, url); err != nil {
		return nil, err
	}


	return &CreateDataOutput{
		URL:  url.URL,
	}, nil
}