package main

import (
	"fmt"
	"os"
	"errors"
	"log"
//	"io"
	"time"
//	"crypto/md5"
)

type directory_tree struct {

	root string
	directories []*directory
	files []*File
	file_counter, directory_counter, file_inc, dir_inc int

}

type directory struct {

	root string
	directories []*directory
	files []*File
	file_counter, directory_counter, file_inc, dir_inc int

}

type File struct {

	Name, Hash string
	size int64
	modification_time time.Time	

}

const (

	START int = 50

)

func main() {

	treeA := directory_tree{file_inc: 1, dir_inc: 1}
	treeA.directories = make([]*directory, START)
	treeA.files = make([]*File, START)

	treeB := directory_tree{file_inc: 1, dir_inc: 1}
	treeB.directories = make([]*directory, START)
	treeB.files = make([]*File, START)

	if len(os.Args) == 1 {

		fmt.Println("Enter path A:")
		fmt.Scan(&treeA.root)
		fmt.Println("Enter path B:")
		fmt.Scan(&treeB.root)
	} else {
		treeA.root = os.Args[1]
		treeB.root = os.Args[2]
	}


	if err := explore_tree(&treeA); err != nil {
		fmt.Println(err)
	}
	if err := explore_tree(&treeB); err != nil {
		fmt.Println(err)
	}

	fmt.Println(treeA.directories, treeA.files)
	fmt.Println(treeB.directories, treeB.files)

}

func explore_tree(tree *directory_tree) error {

	if info, fail := os.Stat(tree.root); fail == nil {

		if !info.IsDir() {

			err := errors.New("Path is not a directory.")
			return err
		}
	} else {
		return fail
	}

	if file, err := os.Open(tree.root); err == nil {

		defer file.Close()

		if out, err := file.Readdirnames(0); err == nil {

			for _, i := range out {

				if fInfo, fail := os.Stat(tree.root + "/" + i); fail == nil {

					if fInfo.IsDir() {

						tree.directories[tree.directory_counter] = newDir(tree.root + "/" + i)
						tree.directory_counter++
						// Correct auto lengthening slices
						//if tree.directory_counter == len(tree.directories) {
						//	tree.dir_inc++
							//b := make([]string, START * tree.dir_inc)
							//tree.directories = copy(tree.directories, b)
						//}
					} else {

						tree.files[tree.file_counter], _ = newFile(tree.root + "/" + i)
						tree.file_counter++

					} // End if IsDir
				} else {
					return fail
				} // End if fInfo
			} // End for
		}  // End if Readdirnames
	} else {
		log.Print(err)
	}  // End if Open
	/*
	outA, err = fileA.Readdirnames(0)
	for i := range outA {
		if aInfo, fail := os.Stat(i); fail == nil {

		}
		fmt.Println(pathA + "/" + outA[i])
	}
*/

	return nil;
}

func (file *File) String() string {

	return fmt.Sprint(file.Name)
}

func (dir *directory) String() string {

	return fmt.Sprint(dir.root)
}

func newFile(x string) (*File, error) {
	
	temp := File{}

	fInfo, fail := os.Stat(x)

		temp.size = fInfo.Size()
		temp.Name = fInfo.Name()

	return &temp, fail
}

func newDir(x string) ( *directory ) {

	temp := directory { root: x, file_inc: 1, dir_inc: 1 }

	temp.directories = make( []*directory, START )
	temp.files = make( []*File, START )

	return &temp
}

/*
	parent string
	directories []directory
	files []*File
	file_counter, directory_counter, file_inc, dir_inc int


		fmt.Println("Name:", aInfo.Name())
		fmt.Println("Size:", aInfo.Size())
		fmt.Println("Mode:", aInfo.Mode())
		fmt.Println(aInfo.Sys())


type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}
*/