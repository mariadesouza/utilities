package fileutil

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"time"
)

//DeleteOldFiles : Delete all  files  in a directoy older than passed days
/* Input :
  	dirpath - absolute path of directory
 		days - number of days to remove the files 
*/
func DeleteOldFiles(dirpath string, days int) error {

	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return (err)
	}

	t0 := time.Now().AddDate(0, 0, 0-days)

	for _, file := range files {
		modtime := file.ModTime()
		if modtime.Before(t0) {
				os.Remove(filepath.Join(dirpath, file.Name()))
				fmt.Println("Removed:" + filepath.Join(dirpath, file.Name()) + ":" + modtime.Format("2006-01-02"))
		}
	}
	return nil
}
