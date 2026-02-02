package resume

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Save(resume Resume) error
}

type mongoRepo struct {
	db         *mongo.Client
	collection *mongo.Collection
}

// 2. ปรับฟังก์ชัน New เพื่อรับ Connection ของ Mongo เข้ามา
func NewRepository(db *mongo.Client) Repository {
	col := db.Database("AI-Resume-project").Collection("resumes")
	return &mongoRepo{
		db:         db,
		collection: col,
	}
}

// 3. เขียน Logic การบันทึกลง MongoDB จริงๆ
func (r *mongoRepo) Save(resume Resume) error {
	_, err := r.collection.InsertOne(context.TODO(), resume)
	return err
}
