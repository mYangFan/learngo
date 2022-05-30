package syncx

type Processor struct {
	syncFunc func(campaignTypeJx string)
	//pullFunc func(ctx context.Context, *model.SyncParam)
}