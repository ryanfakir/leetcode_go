var dict map[string]bool 

func findWords(board [][]byte, words []string) []string {
    dict = make(map[string]bool)
    t := &Trie{}
    
    for _ ,v := range words {
        t.Insert(v)
    }
    m, n := len(board), len(board[0])
    visited := make([][]bool, m)
    for i:= range visited {
        visited[i] = make([]bool, n)
    }
    
    for i := 0; i< m; i++ {
        for j:= 0; j < n; j++ {
            dfs(visited, board, i, j, "", t)
        }
    }
    var res []string
    for v := range dict {
        res = append(res, v)
    }
    return res
}


func dfs(visited [][]bool, board [][]byte, x, y int, s string, t *Trie) {
    if x  < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || visited[x][y] {
        return
    }
    
    s += string(board[x][y])
    if !t.StartsWith(s) { return }
    if t.Search(s) {
        dict[s] = true
    }
    visited[x][y] = true
    dfs(visited, board, x-1, y, s, t)
    dfs(visited, board, x+1, y, s, t)
    dfs(visited, board, x, y-1, s, t)
    dfs(visited, board, x, y+ 1, s, t)
    visited[x][y] = false
}


type Trie struct {
    path []*Trie
    end bool
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    if len(word) == 0 {return}
    if this.path == nil {
        this.path = make([]*Trie, 26)
    }
    if this.path[word[0] - 'a'] == nil {
        this.path[word[0]-'a'] = &Trie{}
        this.path[word[0]- 'a'].end = len(word) == 1
    } else {
        if !this.path[word[0]- 'a'].end {
            this.path[word[0]- 'a'].end = len(word) == 1
        }
    }
    this = this.path[word[0]-'a']
    this.Insert(word[1:])
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    if len(word) == 0 {return true}
    if this == nil || this.path == nil || this.path[word[0] - 'a'] == nil {return false}
    if len(word) == 1 { return this.path[word[0]-'a'].end}
    this = this.path[word[0]-'a']
    return this.Search(word[1:])
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    if len(prefix) == 0 {return true}
    if this == nil || this.path == nil || this.path[prefix[0] - 'a'] == nil {return false}
    if len(prefix) == 1 {return this.path[prefix[0] - 'a'] != nil}
    this = this.path[prefix[0] - 'a']
    return this.StartsWith(prefix[1:])
}

                                                      


