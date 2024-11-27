package models

type DiaryEntry struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    Date    string `json:"date"`
}