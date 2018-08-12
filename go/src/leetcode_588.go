type FileSystem struct {
    file *File
}

type File struct {
    isFile bool
    content string
    files map[string]*File
}


func Constructor() FileSystem {
    return FileSystem{file: new(File)}
}


func (this *FileSystem) Ls(path string) []string {
    f := this.file
    if path != "/" {
       paths := strings.Split(path, "/")
        for i, p := range paths {
            if i == 0 {continue}
            f = f.files[p]
        }
        if f.isFile {
            return []string{paths[len(paths)-1]}
        }
    }
    fmt.Println(f)
    var res []string
    for k := range f.files {
        res = append(res, k)
    }
    sort.Strings(res)
    return res
}


func (this *FileSystem) Mkdir(path string)  {
    f := this.file
    paths := strings.Split(path, "/")
    for i, p := range paths {
        if i == 0 {continue}
        if f.files == nil {f.files = make(map[string]*File)}
        if _, ok := f.files[p]; !ok  {
            f.files[p] = new(File)
        }
        f = f.files[p]
    }
}


func (this *FileSystem) AddContentToFile(filePath string, content string)  {
    f := this.file
    paths := strings.Split(filePath, "/")
    for i, p := range paths {
        if i == 0 || i == len(paths)-1 {continue}
        f = f.files[p]
    }
    if f.files == nil {
        f.files =  make(map[string]*File)
        f.files[paths[len(paths)-1]] = new(File)
    } else {
        if _, ok := f.files[paths[len(paths)-1]]; !ok  {
            f.files[paths[len(paths)-1]] = new(File)
        }
    }
    t := f.files[paths[len(paths)-1]]
    t.isFile = true
    f.files[paths[len(paths)-1]].content += content
}


func (this *FileSystem) ReadContentFromFile(filePath string) string {
    f := this.file
    paths := strings.Split(filePath, "/")
    for i, p := range paths {
        if i == 0 || i == len(paths)-1 {continue}
        f = f.files[p]
    }
    return f.files[paths[len(paths)-1]].content 
    
}

