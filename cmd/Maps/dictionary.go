package main

func Search(dictionary map[string]string, word string) (string, error) {
	return dictionary[word]
}
