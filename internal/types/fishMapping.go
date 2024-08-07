package types

import (
	"errors"
	"fmt"
)

type CreateFishPayload struct {
	DiveID uint
	FishID uint
}

func (fm CreateFishPayload) Validate() error {
	if fm.DiveID == 0 {
		fmt.Println(fm.DiveID)
		return errors.New("diveId is required")
	} else if fm.FishID == 0 {
		return errors.New("fishId is required")
	}

	return nil
}
