package upai

import (
	"github.com/7sDream/rikka/api"
	"github.com/7sDream/rikka/plugins"
)

const (
	stateUploading     = "uploading"
	stateUploadingCode = 2
	stateUploadingDesc = "Rikka is uploading your image to UPai cloud"
)

func buildUploadingState(taskID string) *api.State {
	return &api.State{
		TaskID:      taskID,
		State:       stateUploading,
		StateCode:   stateUploadingCode,
		Description: stateUploadingDesc,
	}
}

func (qnp upaiPlugin) StateRequestHandle(taskID string) (pState *api.State, err error) {
	l.Debug("Receive a state request of taskID", taskID)

	pState, err = plugins.GetTaskState(taskID)
	if err == nil {
		if pState.StateCode == api.StateErrorCode {
			l.Warn("Get a error state of task", taskID, *pState)
		} else {
			l.Debug("Get a normal state of task", taskID, *pState)
		}
		return pState, nil
	}

	l.Debug("State of task", taskID, "not found, just return a finish state")
	return api.BuildFinishState(taskID), nil
}
