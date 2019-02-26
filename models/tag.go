package models

type Tag struct {
    Name       string `json:"name" bson:"name"`
    CreatedBy  string `json:"created_by" bson:"createdBy"`
    ModifiedBy string `json:"modified_by" bson:"modifiedBy"`
    State      int    `json:"state" bson:"state"`
}
