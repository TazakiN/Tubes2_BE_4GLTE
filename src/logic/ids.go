package logic

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Method     string `json:"method"`
	LinkAwal   string `json:"linkAwal"`
	LinkTujuan string `json:"linkTujuan"`
}

func IDS(linkMulai string, linkTujuan string, bahasa string) []string {
	result := Result{
		Method:     "IDS",
		LinkAwal:   linkMulai,
		LinkTujuan: linkTujuan,
	}

	// proses pencarian jalur di sini

	jsonResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return []string{string(jsonResult)}
}
