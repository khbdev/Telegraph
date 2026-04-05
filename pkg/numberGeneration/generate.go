func getMaxCounter(db *gorm.DB, base string) int {
    var lastSlug string
    db.Model(&models.Post{}).
        Select("slug").
        Where("slug LIKE ?", base+"%").
        Order("id DESC").
        Limit(1).
        Scan(&lastSlug)

    // salom-02 → 2 ni ajratish
    parts := strings.Split(lastSlug, "-")
    if len(parts) > 1 {
        n, err := strconv.Atoi(parts[len(parts)-1])
        if err == nil {
            return n
        }
    }
    return 0
}