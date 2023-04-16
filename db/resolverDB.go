package db

import (
	"context"
	"newCurriculum/graph/model"
)

// 3つのリゾルバはここにつながる。
func CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	db := GetDB()
	var taskT taskTable
	var labelT labelTable
	var taskLabelT taskLabelTable
	err := taskT.insert(ctx, db, convertToTask(input))
	if err != nil {
		return nil, err
	}
	labelID, err := labelT.selectID(ctx, db, input.LabelValue)
	if err != nil {
		return nil, err
	}
	err = taskLabelT.insert(ctx, db, createTaskLabelRealation(input.ID, labelID))
	if err != nil {
		return nil, err
	}
	return shapeInput(input), nil
}

func UpdateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	err := checkInput(input)
	if err != nil {
		return nil, err
	}
	db := GetDB()
	var taskT taskTable
	var labelT labelTable
	var taskLabelT taskLabelTable
	err = taskT.update(ctx, db, convertToTask(input))
	if err != nil {
		return nil, err
	}
	labelID, err := labelT.selectID(ctx, db, input.LabelValue)
	if err != nil {
		return nil, err
	}
	err = taskLabelT.update(ctx, db, createTaskLabelRealation(input.ID, labelID))
	if err != nil {
		return nil, err
	}
	return shapeInput(input), nil
}

func DeleteTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	err := checkInput(input)
	if err != nil {
		return nil, err
	}
	db := GetDB()
	var taskT taskTable
	var taskLabelT taskLabelTable
	if err != nil {
		return nil, err
	}
	err = taskLabelT.delete(ctx, db, input.ID)
	if err != nil {
		return nil, err
	}
	err = taskT.delete(ctx, db, input.ID)
	if err != nil {
		return nil, err
	}

	return shapeInput(input), nil
}
