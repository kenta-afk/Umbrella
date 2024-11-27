package controllers

import (
    "net/http"
    "diary/models"
    "diary/config"
    "github.com/gin-gonic/gin"
)

func GetEntries(c *gin.Context) {
    var entries []models.DiaryEntry
    rows, err := config.DB.Query("SELECT id, title, content, date FROM diary_entries")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    for rows.Next() {
        var entry models.DiaryEntry
        if err := rows.Scan(&entry.ID, &entry.Title, &entry.Content, &entry.Date); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        entries = append(entries, entry)
    }

    c.JSON(http.StatusOK, entries)
}

func CreateEntry(c *gin.Context) {
    var entry models.DiaryEntry
    if err := c.ShouldBindJSON(&entry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := config.DB.Exec("INSERT INTO diary_entries (title, content, date) VALUES (?, ?, ?)", entry.Title, entry.Content, entry.Date)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Entry created successfully"})
}