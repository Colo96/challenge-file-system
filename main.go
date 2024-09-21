package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type File struct {
	name string
}

type Directory struct {
	name    string
	files   []File
	subdirs []*Directory
	parent  *Directory
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		name:    name,
		files:   []File{},
		subdirs: []*Directory{},
		parent:  parent,
	}
}

func (d *Directory) touch(filename string) {
	d.files = append(d.files, File{name: filename})
}

func (d *Directory) mkdir(dirname string) {
	newDir := newDirectory(dirname, d)
	d.subdirs = append(d.subdirs, newDir)
}

func (d *Directory) cd(dirname string) *Directory {
	if dirname == ".." {
		if d.parent != nil {
			return d.parent
		}
		return d
	}
	for _, subdir := range d.subdirs {
		if subdir.name == dirname {
			return subdir
		}
	}
	fmt.Println("Directorio no encontrado: ", dirname)
	return d
}

func (d *Directory) ls() {
	for _, dir := range d.subdirs {
		fmt.Println("[Dir]", dir.name)
	}
	for _, file := range d.files {
		fmt.Println("[File]", file.name)
	}
}

func (d *Directory) pwd() {
	var path []string
	current := d
	for current != nil {
		path = append([]string{current.name}, path...)
		current = current.parent
	}
	fmt.Println("/", strings.Join(path, "/"))
}

func main() {
	root := newDirectory("", nil)
	currentDir := root

	reader := bufio.NewReader(os.Stdin)

	for {
		currentDir.pwd()
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]
		var argument string
		if len(parts) > 1 {
			argument = parts[1]
		}

		switch command {
		case "cd":
			if argument == "" {
				fmt.Println("Falta el nombre del directorio")
			} else {
				currentDir = currentDir.cd(argument)
			}
		case "touch":
			if argument == "" {
				fmt.Println("Falta el nombre del archivo")
			} else {
				currentDir.touch(argument)
			}
		case "mkdir":
			if argument == "" {
				fmt.Println("Falta el nombre de la carpeta")
			} else {
				currentDir.mkdir(argument)
			}
		case "ls":
			currentDir.ls()
		case "pwd":
			currentDir.pwd()
		case "exit":
			fmt.Println("Saliendo del sistema de archivos")
			return
		default:
			fmt.Println("Comando no reconocido", command)
		}
	}

}
