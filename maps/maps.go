package maps

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrorNotFound        = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("Cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
func (d Dictionary) Search(value string) (string, error) {
	definition, ok := d[value]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Update(word, defintion string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = defintion
	default:
		return err
	}
	return nil
}
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
