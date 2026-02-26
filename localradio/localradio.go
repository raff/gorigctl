package localradio

import (
	"log"

	hl "github.com/dh1tw/goHamlib"
)

type LocalRadio struct {
	rig hl.Rig
	log *log.Logger
	vfo hl.VFOType
}

func NewLocalRadio(rigModel hl.RigModelID, debugLevel hl.DebugLevel, port hl.Port, log *log.Logger) (*LocalRadio, error) {
	lr := LocalRadio{}
	lr.rig = hl.Rig{}
	lr.log = log
	lr.vfo = hl.VFOCurrent

	hl.SetDebugLevel(debugLevel)

	if err := lr.rig.Init(rigModel); err != nil {
		return nil, err
	}

	// only set the port if it's not the dummy model
	if rigModel != 1 {
		if err := lr.rig.SetPort(port); err != nil {
			return nil, err
		}
	}

	if err := lr.rig.Open(); err != nil {
		return nil, err
	}

	vfo, err := lr.rig.GetVfo()
	if err != nil {
		lr.log.Println(err)
	} else {
		lr.vfo = vfo
	}

	return &lr, nil
}
