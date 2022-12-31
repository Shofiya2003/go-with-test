package dictionary

const (
	WordNotFoundError      = DictionaryError("word not found")
	WordAlreadyExistsError = DictionaryError("word already exists")
)

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}
func (d Dictionary) Search(word string) (string, error) {
	definiton, ok := d[word]
	if !ok {
		return "", WordNotFoundError
	}
	return definiton, nil
}

func (d Dictionary) AddWord(word string, def string) error {
	_, ok := d[word]
	if ok {
		return WordAlreadyExistsError
	}
	d[word] = def
	return nil
}

func (d Dictionary) Update(word string, newDef string) error {
	_, ok := d[word]
	if !ok {
		return WordNotFoundError
	}
	d[word] = newDef
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
