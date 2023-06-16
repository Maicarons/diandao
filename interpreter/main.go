//    gameText, err := UnmarshalGameText(bytes)
//    bytes, err = gameText.Marshal()

package interpreter

import "encoding/json"

func UnmarshalGameText(data []byte) (GameText, error) {
	var r GameText
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GameText) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GameText struct {
	Scriptname    string `json:"scriptname"`
	Description   string `json:"description"`
	Author        string `json:"author"`
	Formatversion string `json:"formatversion"`
	Text          []Text `json:"text"`
}

type Text struct {
	Speaker string `json:"speaker"`
	Remark  string `json:"remark"`
	Content string `json:"content"`
}
