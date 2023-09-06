package usecase

import (
	"ai-typing/model"
	"ai-typing/repository"
	"fmt"
)

type IGameUsecase interface {
	CreateGame(game model.Game) (model.Game, error)
	GetGameRanking(border int) ([]model.Game, error)
	GetGameHistory(userId string, limit int) ([]model.Game, error)
	GetAllGame() ([]model.Game, error)
	GetLatestGames(offset int) ([]model.Game, error)
	GetTotalGameCount() (int64, error)
	UpdateGameScore(game *model.Game) (model.UpdateGameResponse, error)
	GetAllByUserId(userId string) ([]model.Game, error)
}

type gameUsecase struct {
	gameRepository        repository.IGameRepository
	createdTextRepository repository.ICreatedTextRepository
	batchRepository       repository.IBatchRepository
}

func NewGameUsecase(gameRepository repository.IGameRepository,
	createdTextRepository repository.ICreatedTextRepository,
	batchRepository repository.IBatchRepository,
) IGameUsecase {
	return &gameUsecase{gameRepository, createdTextRepository, batchRepository}
}

func (gameUsecase *gameUsecase) CreateGame(game model.Game) (model.Game, error) {
	if err := gameUsecase.gameRepository.CreateGame(&game); err != nil {
		return model.Game{}, err
	}
	return game, nil
}

func (gameUsecase *gameUsecase) GetGameRanking(border int) ([]model.Game, error) {
	games := []model.Game{}
	if err := gameUsecase.gameRepository.GetGameRanking(&games, border); err != nil {
		return nil, err
	}
	return games, nil
}

func (gameUsecase *gameUsecase) GetGameHistory(userId string, limit int) ([]model.Game, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetGameHistory(&games, userId, limit)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (gameUsecase *gameUsecase) GetAllGame() ([]model.Game, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetAllGame(&games)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (gameUsecase *gameUsecase) GetLatestGames(offset int) ([]model.Game, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetLatestGames(&games, offset)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (gameUsecase *gameUsecase) GetTotalGameCount() (int64, error) {
	count, err := gameUsecase.gameRepository.GetTotalGameCount()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func selectBatch(game *model.Game, newBatches *[]model.Batch, currentBatches *[]model.Batch) error {
	//スコアによるバッジの作成
	mapList := map[string]int{
		"earth":    1,
		"moon":     500,
		"mars":     1000,
		"mercury":  1500,
		"jupiter":  2000,
		"venus":    2500,
		"saturn":   3000,
		"sun":      3500,
		"universe": 4500,
	}
	for key, value := range mapList {
		if game.Score >= value {
			batch := model.Batch{
				UserId: game.UserId,
				ModeId: game.ModeId,
				Name:   key,
			}
			isExist := false
			for _, currentBatch := range *currentBatches {
				//すべて一致しない(現在取得していない)ときだけnewに追加
				if currentBatch.Name == key && currentBatch.UserId == game.UserId && currentBatch.ModeId == game.ModeId {
					isExist = true
					break
				}
			}
			if !isExist {
				*newBatches = append(*newBatches, batch)
			}
		}
	}
	return nil
}

func (gameUsecase *gameUsecase) UpdateGameScore(game *model.Game) (model.UpdateGameResponse, error) {
	err := gameUsecase.gameRepository.UpdateGameScore(game) //gameには変更後の全てのデータを含む
	if err != nil {
		return model.UpdateGameResponse{}, err
	}
	border := 100
	count, err := gameUsecase.gameRepository.GetRankingCount(border)
	if err != nil {
		return model.UpdateGameResponse{}, err
	}
	gameId := game.ID
	rank, err := gameUsecase.gameRepository.GetRankByGameId(100, gameId)
	if err != nil {
		return model.UpdateGameResponse{}, err
	}

	//ランキングに反映されるゲームのみバッチを作成
	newBatches := []model.Batch{}
	if !game.DisableRanking {
		if game.UserId != "" {
			//現在のバッチを取得
			currentBatches := []model.Batch{}
			if err := gameUsecase.batchRepository.GetAllByUserId(&currentBatches, game.UserId); err != nil {
				fmt.Println(err.Error())
				return model.UpdateGameResponse{}, err
			}

			//スコアに応じてバッチを作成
			err := selectBatch(game, &newBatches, &currentBatches)
			if err != nil {
				fmt.Println(err.Error())
				return model.UpdateGameResponse{}, err
			}

			//新規バッチがあるときだけ新規バッチの更新
			if len(newBatches) > 0 {
				if err := gameUsecase.batchRepository.Create(&newBatches); err != nil {
					fmt.Println(err.Error())
					return model.UpdateGameResponse{}, err
				}
			}
		}
	}

	response := model.UpdateGameResponse{
		Count:   int(count),
		Rank:    int(rank),
		Batches: newBatches,
	}
	return response, nil
}

func (gameUsecase *gameUsecase) GetAllByUserId(userId string) ([]model.Game, error) {
	games := []model.Game{}
	err := gameUsecase.gameRepository.GetAllByUserId(&games, userId)
	if err != nil {
		return nil, err
	}
	return games, nil
}
