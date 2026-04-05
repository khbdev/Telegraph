package numbergeneration

import "gorm.io/gorm"



func GenerateSlug(db *gorm.DB, title string) (string, error) {
    base := normalize(title) // Salom → salom
    slug := base

    // Hash index orqali tekshiruv
    exists := false
    db.Model(& qo‘shamiz
        counter := getMaxCounter(db, base) + 1
        slug = fmt.Sprintf("%s-%02d", base, counter)
    }

    return slug, nil
}