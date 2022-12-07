package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const sizeThresh = 100000

var (
	ans  = 0
	req  = math.MaxInt
	ans2 = math.MaxInt
)

type File struct {
	name string
	size int
}

type Directory struct {
	name   string
	files  map[string]*File
	dirs   map[string]*Directory
	parent *Directory
}

func (d *Directory) Size() int {
	size := 0
	for _, v := range d.files {
		size += v.size
	}
	for _, v := range d.dirs {
		size += v.Size()
	}
	if size <= sizeThresh {
		fmt.Println(size, d.name)
		ans += size
	}
	if size >= req && size < ans2 {
		ans2 = size
	}
	return size
}

func main() {
	f, err := os.Open("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	root := &Directory{
		name:  "/",
		files: make(map[string]*File),
		dirs:  make(map[string]*Directory),
	}
	root.parent = root

	curDir := root
	var s string
	for scanner.Scan() {
		s = scanner.Text()
		if s[0] == '$' {
			if s[2] == 'c' {
				q := strings.Split(s, " ")
				cd := q[len(q)-1]
				switch cd {
				case "/":
					curDir = root
				case "..":
					curDir = curDir.parent
				default:
					curDir = curDir.dirs[cd]
				}
			}

		} else {
			if s[0] == 'd' {
				dirName := s[4:]
				curDir.dirs[dirName] = &Directory{
					name:   dirName,
					files:  make(map[string]*File),
					dirs:   make(map[string]*Directory),
					parent: curDir,
				}
			} else {
				snum, nam, _ := strings.Cut(s, " ")
				num, _ := strconv.Atoi(snum)
				curDir.files[nam] = &File{
					name: nam,
					size: num,
				}
			}
		}
	}
	req = 30000000 - 70000000 + root.Size()
	fmt.Println(ans)
	root.Size()
	fmt.Println(ans2)

}
