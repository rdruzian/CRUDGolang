package news

import (
	"awesomeProject/news/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetValues() ([]types.News, error) {
	var news []types.News

	file, err := os.Open("validation.json")
	if err != nil {
		fmt.Printf("Erro ao ler arquivo json")
		return news, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Erro ao ler bytes %v", err)
		return news, err
	}

	if json.Unmarshal(byteValue, &news); err != nil {
		fmt.Printf("Erro ao trazer dados do json %v", err)
		return news, err
	}

	return news, nil
}