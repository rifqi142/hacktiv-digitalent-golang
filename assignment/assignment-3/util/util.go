package util

import (
	"encoding/json"
	"fmt"
	"hacktiv-digitalent-golang/assignment/assignment-3/model"
	"os"
)


func WriteStatusFile(status model.Status, filePath string) error {
	data, err := json.Marshal(status)
	if err != nil {
		return fmt.Errorf("failed to marshal status data: %v", err)
	}
	
	if _, err := os.Stat(filePath); err == nil {
		
		err = os.WriteFile(filePath, data, 0644)
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	} else if os.IsNotExist(err) {
		
		err = os.WriteFile(filePath, data, 0644)
		if err != nil {
			return fmt.Errorf("failed to create and write to file: %v", err)
		}
	} else {
		
		return fmt.Errorf("failed to check file existence: %v", err)
	}

	return nil
}