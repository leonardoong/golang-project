package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"part1/internal/constant"
	"part1/internal/model"
	"sort"
	"strings"
)

func (d *datalogic) GetDataFromJSON(ctx context.Context) error {
	var (
		err     error
		dirPath = "/subsetdata"
	)

	// 1. get all files in subsetdata dir
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// 2. read all files using go routine
	tasks := make(chan model.AsyncTaskRes)
	for idx, file := range files {
		go readFiles(idx, file, dirPath, tasks)
	}

	tasksRes := []model.AsyncTaskRes{}
	for range files {
		got := <-tasks
		if got.Err != nil {
			err = got.Err
			return err
		}

		tasksRes = append(tasksRes, got)
	}

	// 3. sort based on file index
	sort.Slice(tasksRes, func(i, j int) bool {
		return tasksRes[i].Idx < tasksRes[j].Idx
	})

	for _, task := range tasksRes {
		if val, ok := task.Res.([]model.Transaction); ok {
			for _, tx := range val {
				// skip type A with quantity
				// quantity 0 needed for open price
				if tx.Type == "A" && tx.Quantity != "" {
					continue
				}
				txByte, errMarshal := json.Marshal(tx)
				if errMarshal != nil {
					log.Printf("failed marshal transaction : %v , err : %v \n", tx, err)
					return errMarshal
				}
				// 4. send producer
				_, _, err = d.producer.SendMessage(ctx, constant.TransactionTopic, string(txByte))
				if err != nil {
					log.Printf("failed send message transaction : %v , err : %v \n", tx, err)
					return err
				}
			}
		}
	}

	return err
}

func readFiles(idx int, file fs.DirEntry, dirPath string, out chan<- model.AsyncTaskRes) {
	var (
		err error
		res []model.Transaction
	)

	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", dirPath, file.Name()))
	if err != nil {
		log.Printf("failed open file, err : %v \n", err)
		out <- model.AsyncTaskRes{
			Err: err,
			Idx: idx,
			Res: res,
		}
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Printf("failed read file, err : %v , file name : %v\n", err, file.Name())
		out <- model.AsyncTaskRes{
			Err: err,
			Idx: idx,
			Res: res,
		}
	}

	got := json.NewDecoder(strings.NewReader(string(byteValue)))
	for {
		var tx model.Transaction
		err = got.Decode(&tx)
		if err == io.EOF {
			err = nil
			break
		}

		if err != nil {
			log.Printf("failed decode json, err : %v\n", err)
			out <- model.AsyncTaskRes{
				Err: err,
				Idx: idx,
				Res: res,
			}
		}
		res = append(res, tx)
	}
	out <- model.AsyncTaskRes{
		Err: err,
		Idx: idx,
		Res: res,
	}
}
