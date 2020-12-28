package models

import "time"



type News struct {
	NewsId			int32			`orm:"auto"`
	UserId			int32			`orm:"null"`
	Title			string			`orm:"size(200);null"`
	Content			string			`orm:"type(text)"`
	Time			time.Time		`orm:"auto_now_add;type(date);default(2019-05-01 19:00:00)"`
	Importance		int8			`orm:"default(0)"`
	Defunct			string			`orm:"type(char);size(1);default(N)"`
}




func AddArt(uid int32, title string, content string) (int64, error) {
	var Art News
	Art.UserId = uid
	Art.Content = content
	Art.Title = title
	Art.Time = time.Now()
	artId, err := DB.Insert(&Art)
	if err != nil {
		return artId, err
	}
	return artId, nil
}


func QueryAllArt() ([]*News, error) {
	var art []*News
	news := new(News)
	qs := DB.QueryTable(news)
	_, err := qs.OrderBy("-news_id").All(&art)
	if err != nil {
		return nil, err
	}
	return art, nil
}

func QueryLimitArt() ([]*News, error) {
	var art []*News
	news := new(News)
	qs := DB.QueryTable(news)
	_, err := qs.OrderBy("-news_id").Limit(2).All(&art)
	if err != nil {
		return nil, err
	}
	return art, nil
}



func QueryArtByArtId(id int32)(News, error) {
	art := News{NewsId:id}
	err := DB.Read(&art,"NewsId")
	if err != nil {
		return News{}, err
	}
	return art, nil

}
