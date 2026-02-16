package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	path := os.Args[1]

	err := filepath.WalkDir(path, func(paths string, d fs.DirEntry, err error) error {

		if d.IsDir() {
			fmt.Printf("dir in: %s\n", d.Name())
			return nil
		}

		if strings.HasSuffix(paths, ".mkv") {

			info, err := d.Info()
			if err != nil {
				return err

			}

			filename := path + "/" + info.Name()
			newname := filename + ".mov"

			cmd := exec.Command("ffmpeg", "-i", filename, "-c", "copy", newname)

			output, err := cmd.CombinedOutput()

			if err != nil {

				fmt.Printf("ffmpeg failed with err: %s\n", err)
				fmt.Printf("stdout is : %s\n", output)

				return nil
			}

			fmt.Println(string(output))

		}
		return nil

	})
	if err != nil {
		fmt.Printf("err is: %s", err)
	}

}
