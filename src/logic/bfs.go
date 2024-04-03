package logic

import (
	"encoding/json"
	"fmt"
)

func BFS(linkMulai string, linkTujuan string) []string {
	result := Result{
		Method:     "BFS",
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
