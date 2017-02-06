package main

type WordDictionary struct {
	sub []*WordDictionary
	end bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for i := 0; i < len(word); i++ {

		if this.sub == nil {
			this.sub = make([]*WordDictionary, 26)
		}
		if this.sub[word[i]-'a'] == nil {
			this.sub[word[i]-'a'] = new(WordDictionary)
		}
		this = this.sub[word[i]-'a']
	}

	this.end = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return Sear(word, 0, this)
}

func Sear(word string, level int, dict *WordDictionary) bool {
	if len(word) == level {
		return dict.end
	}
	if dict.sub == nil {
		return false
	}
	if word[level] == '.' {
		for _, v := range dict.sub {
			if v != nil {
				if Sear(word, level+1, v) {
					return true
				}
			}
		}
	} else {
		if dict.sub[word[level]-'a'] != nil && Sear(word, level+1, dict.sub[word[level]-'a']) {
			return true
		}
	}
	return false
}
