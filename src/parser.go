package src

import (
	"encoding/json"
	"fmt"
	"io"
)

type (
	VMAFData struct {
		Version string  `json:"version"`
		FPS     float64 `json:"fps"`
		Frames  []Frame `json:"frames"`
	}
	Frame struct {
		FrameNum int          `json:"frameNum"`
		Metrics  FrameMetrics `json:"metrics"`
	}
	FrameMetrics struct {
		IntegerADM2      float64 `json:"integer_adm2"`
		IntegerADMScale0 float64 `json:"integer_adm_scale0"`
		IntegerADMScale1 float64 `json:"integer_adm_scale1"`
		IntegerADMScale2 float64 `json:"integer_adm_scale2"`
		IntegerADMScale3 float64 `json:"integer_adm_scale3"`
		IntegerMotion2   float64 `json:"integer_motion2"`
		IntegerMotion    float64 `json:"integer_motion"`
		IntegerVIFScale0 float64 `json:"integer_vif_scale0"`
		IntegerVIFScale1 float64 `json:"integer_vif_scale1"`
		IntegerVIFScale2 float64 `json:"integer_vif_scale2"`
		IntegerVIFScale3 float64 `json:"integer_vif_scale3"`
		VMAF             float64 `json:"vmaf"`
	}
)

// ParseJSON parses the given JSON data into a JSONData struct
func ParseJSON(reader io.Reader) (*VMAFData, error) {
	var data VMAFData
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	return &data, nil
}
