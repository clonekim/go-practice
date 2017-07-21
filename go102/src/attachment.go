package main

type Attachment struct {
	Id          int     "json:`id`"
	RefId       int     "json:`ref_id`"
	TableName   string  "json:`table_name`"
	FileName    string  "json:`filename`"
	DiskName    string  "json:`diskname`"
	FileSize    int64   "json:`filesize`"
	ContentType string  "json:`content_type`"
	Deleted     string  "json:`deleted`"
	CreatedAt   string  "json:`created_at`"
	DeletedAt   *string "json:`deleted_at,omitempty`"
}

func SelectAttachmentsByRefId(refid int) (attachments []Attachment, err error) {

	rows, err := DB.Query("SELECT id, ref_id, entity_name, file_name, disk_filename, file_size, mime_type, deleted, file_create, file_delete FROM attachments WHERE ref_id =?", refid)

	if err == nil {
		defer rows.Close()

		for rows.Next() {
			atta := Attachment{}
			rows.Scan(&atta.Id, &atta.RefId, &atta.TableName, &atta.FileName, &atta.DiskName, &atta.FileSize, &atta.ContentType, &atta.Deleted, &atta.CreatedAt, &atta.DeletedAt)
			attachments = append(attachments, atta)
		}

	}
	return

}

func SelectAttachmentsById(id int) (atta Attachment, err error) {
	err = DB.QueryRow("SELECT id, ref_id, entity_name, file_name, disk_name, file_size, mime_type, deleted, file_create, file_delete FROM attachments WHERE id =?", id).Scan(&atta.Id, &atta.RefId, &atta.TableName, &atta.FileName, &atta.DiskName, &atta.FileSize, &atta.ContentType, &atta.Deleted, &atta.CreatedAt, &atta.DeletedAt)
	return
}
