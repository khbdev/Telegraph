package numbergeneration

import "gorm.io/gorm"



func GenerateSlug(db *gorm.DB, title string) (string, error) {
    base := norm(title) // Salom → salom
    slug := base

    // Hash index orqali tekshiruv
    exists := false
    db.Model(&Post{}).Select("count(*) > 0").Where("slug = ?", slug).Find(&exists)

    if exists {
        // Bazada slug bor → number qo‘shamiz
        counter := getMaxCounter(db, base) + 1
        slug = fmt.Sprintf("%s-%02d", base, counter)
    }

    return slug, nil
}