package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"part1/internal/constant"
	"part1/internal/model"
	"strings"
)

func (d *datalogic) GetDataFromJSON(ctx context.Context) error {
	var (
		err     error
		dirPath = "/subsetdata"
	)

	// 1. get all files in subsetdata dir
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// 2. read all files using go routine
	for _, file := range files {
		go d.readFiles(ctx, file, dirPath)
	}

	return err
}

func (d *datalogic) readFiles(ctx context.Context, file fs.FileInfo, dirPath string) error {
	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", dirPath, file.Name()))
	if err != nil {
		log.Printf("failed open file, err : %v \n", err)
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("failed read file, err : %v , file name : %v\n", err, file.Name())
		return err
	}

	got := json.NewDecoder(strings.NewReader(string(byteValue)))
	for {
		var tx model.Transaction
		err = got.Decode(&tx)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("failed decode json, err : %v\n", err)
			return err
		}

		txByte, err := json.Marshal(tx)
		if err != nil {
			log.Printf("failed marshal transaction : %v , err : %v \n", tx, err)
			return err
		}
		d.producer.SendMessage(ctx, constant.TransactionTopic, string(txByte))

		fmt.Printf("tx === %v \n", tx)
	}

	return err
}
