package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"strings"
	"time"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}

func UserSendMessage(uids []string, content string) error {
	//事务
	err := orm.NewOrm().DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		//保存发送信息
		insertMessageSql := "INSERT INTO message (content,add_time) VALUES (?,?)"
		result, err := txOrm.Raw(insertMessageSql, content, time.Now().Unix()).Exec()
		if err != nil {
			return fmt.Errorf("保存发送信息异常:%w", err)
		}
		//获取到发送信息的ID
		messageId, _ := result.LastInsertId()
		//保存发送信息所属用户
		insertMessageUserSql := "insert into message_user (user_id,message_id,status,add_time) values "
		var valueParams []string
		var values []interface{}
		currentUnixTime := time.Now().Unix()
		for _, uid := range uids {
			fmt.Println(uid)
			fmt.Println("--------")
			valueParams = append(valueParams, "(?,?,?,?)")
			values = append(values, uid, messageId, messageUserStatusOn, currentUnixTime)
		}
		//拼接SQL
		insertMessageUserSql += strings.Join(valueParams, ",")
		// 执行批量插入
		result, err = txOrm.Raw(insertMessageUserSql, values...).Exec()
		if err != nil {
			return fmt.Errorf("保存信息所属用户异常：%w", err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("无法获取受影响行数：%w", err)
		}
		if int(rowsAffected) != len(uids) {
			return errors.New("保存信息所属用户失败，部分用户数据未插入")
		}
		return nil
	})
	if err != nil {
		logs.Error(err)
		return errors.New("内部异常")
	}
	return nil
}
