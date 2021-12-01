package internel

import "context"

type CACHELEVEL uint32

const (
	Level_Direct  = iota // local cache
	Level_Cache          // distribute and fast db
	Level_Storage        // data storage
)

type DataCentorIFace interface {
	GetProcessor(ctx context.Context, level CACHELEVEL) *ProcessorIFace

	GetData(ctx context.Context)
}

type DataCentor struct {
}

func (s *DataCentor) GetProcessor(ctx context.Context, level CACHELEVEL) *ProcessorIFace {
	return nil
}

type ProcessorIFace struct {
	//
}

type Margzine struct {
}

func (d *DataCentor) GetMarzine(ctx context.Context, sid uint64) *Margzine {
	// storage 查找
	// d.GetProcessor(ctx, Level_Storage).Get(sid)
	//
	// 分配库存
	// get from db
	//
	// 扣减
	return nil
}

// [key-value]
// [key-set]
