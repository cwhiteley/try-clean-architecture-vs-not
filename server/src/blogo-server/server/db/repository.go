package db

import (
	"blogo-server/server/entity"
	"fmt"
)

//InsertPost - insert post
func InsertPost(p domain.Post) (id int, err error) {

}

// SelectAllPosts - select all
func SelectAllPosts() (posts domain.Posts, err error) {

}

// SelectPost - select post
func SelectPost(table string, identifier int) (post domain.Post, err error) {
	stmtOut, err := Conn.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE message_id = ? LIMIT 1", table))
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var id int64
	var title string
	var content string
	if err := stmtOut.QueryRow(id).Scan(&id, &title, &content); err != nil {
		return
	}
	var postData Post
	postData.RoomName = room_name
	postData.Message = message
	postData.MessageId = message_id
	postData.IsSend = is_send
	return
}

// DeletePost - remove post
func DeletePost(identifier int) (err error) {

}
