package types

import (
	"errors"
	"fmt"
)

type CreateFishMappingPayload struct {
	DiveID int32
	FishID int32
}

func (fm CreateFishMappingPayload) Validate() error {
	if fm.DiveID == 0 {
		fmt.Println(fm.DiveID)
		return errors.New("diveId is required")
	} else if fm.FishID == 0 {
		return errors.New("fishId is required")
	}

	return nil
}
