package repository

import (
	"database/sql"
	"math"
	"time"
	"go-blog/model"
)

// func ArticleList() ([]*model.Article, error) {
// 	// 実行する SQL
// 	query := `SELECT * FROM articles;`

// 	// データベースから取得した値を格納する変数を宣言
// 	var articles []*model.Article

// 	// Query を実行して、取得した値を変数に格納
// 	if err := db.Select(&articles, query); err != nil {
// 		return nil, err
// 	}

// 	return articles, nil
// }

func ArticleListByCursor(cursor int)([]*model.Article,error){
	if cursor <= 0{
		cursor = math.MaxInt32
	}
	// ID の降順に記事データを 10 件取得するクエリ文字列を生成します。
	query := `SELECT *
	FROM articles
	WHERE id < ?
	ORDER BY id desc
	LIMIT 10`

	articles := make([]*model.Article, 0, 10)

	if err := db.Select(&articles,query,cursor);err!=nil{
		return nil,err
	}
	return articles,nil
}


func ArticleCreate(article *model.Article) (sql.Result, error) {
  // 現在日時を取得します
  now := time.Now()

  // 構造体に現在日時を設定します。
  article.Created = now
  article.Updated = now

  // クエリ文字列を生成します。
  query := `INSERT INTO articles (title, body, created, updated)
  VALUES (:title, :body, :created, :updated);`

  // トランザクションを開始します。
  tx := db.MustBegin()

  // クエリ文字列と構造体を引数に渡して SQL を実行します。
  // クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
  // 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
  res, err := tx.NamedExec(query, article)
  if err != nil {
    // エラーが発生した場合はロールバックします。
    tx.Rollback()

    // エラー内容を返却します。
    return nil, err
  }

  // SQL の実行に成功した場合はコミットします。
  tx.Commit()

  // SQL の実行結果を返却します。
  return res, nil
}

func ArticleDelete(id int)error{
	query := `delete from articles where id =?;`

	tx := db.MustBegin()
	
	if _,err := tx.Exec(query,id);err!=nil{
		tx.Rollback()

		return err
	}
	return tx.Commit()
}